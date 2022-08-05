package routes

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"web-api/database"
	"web-api/models"

	"github.com/gofiber/fiber/v2"
)

type Car struct {
	Plate      string `json:"plate"`
	ProducedAt time.Time
	Model      string `json:"model"`
	Color      string `json:"color"`
	Owner      User   `json:"owner"`
}

func CreateResponseCar(car models.Car, owner User) Car {
	return Car{Plate: car.Plate, ProducedAt: car.ProducedAt, Model: car.Model, Color: car.Color, Owner: owner}
}

func CreateCar(c *fiber.Ctx) error {
	var car models.Car

	if err := c.BodyParser(&car); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	match, _ := regexp.MatchString("[ABCDEFGHJKLMNPRSTVWXYZ]{2}[0-9]{3}[ABCDEFGHJKLMNPRSTVWXYZ]{2}", car.Plate)

	if match == false {
		return c.Status(400).JSON("Plate is not valid")
	}

	var user models.User

	if err := findUser(car.OwnerRefer, &user ); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&car)

	responseUser := CreateResponseUser(user)

	responseCar := CreateResponseCar(car, responseUser)
	return c.Status(200).JSON(responseCar)
}

func GetCars(c *fiber.Ctx) error {
	cars := []models.Car{}
	database.Database.Db.Find(&cars)
	responseCars := []Car{}
	for _, car := range cars {
		var user models.User

		if err := findUser(car.OwnerRefer, &user ); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responseCar := CreateResponseCar(car, CreateResponseUser(user))
		responseCars = append(responseCars, responseCar)
	}
	return c.Status(200).JSON(responseCars)
}

func GetCar(c *fiber.Ctx) error {
	plate := c.Params("plate")
	var car models.Car
	if err := findCar(plate, &car); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User

	if err := findUser(car.OwnerRefer, &user ); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseCar := CreateResponseCar(car, CreateResponseUser(user))
	return c.Status(200).JSON(responseCar)
}

func UpdateCar(c *fiber.Ctx) error {
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Save(&car)

	var owner models.User
	database.Database.Db.Find(&owner, "id = ?", car.OwnerRefer)

	responseCar := CreateResponseCar(car, CreateResponseUser(owner))
	return c.Status(200).JSON(responseCar)
}

func DeleteCar(c *fiber.Ctx) error {
	plate := c.Params("plate")

	var car models.Car

	err := findCar(plate, &car)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&car).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Car")
}

func findCar(plate string, car *models.Car) error {
	database.Database.Db.Find(&car, "plate = ?", plate)
	if car.Plate  == "" {
		return errors.New("Plate does not exist")
	}
	return nil
}