package karyawan

import (
	"api/model"
	"api/repository/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func TambahKaryawan(c *fiber.Ctx) error {
	var requestData model.DataKaryawan
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	requestData.ID = uuid.New().String()

	if err := db.CountDataKaryawan(requestData.Username); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count data in the database",
		})
	} else if err != nil && err.Error() == "data already exists" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data already exists in the database",
		})
	}

	// If data doesn't exist, proceed with insertion
	if err := db.InsertDataKaryawan(requestData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into the database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
		"data":    requestData,
	})
}

func GetAllDataKaryawan(c *fiber.Ctx) error {
	filter := bson.M{}

	requestData, err := db.GetDataKaryawanFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}

func GetDataKaryawanByUname(c *fiber.Ctx) error {
	username := c.Params("username")
	filter := bson.M{"username": username}
	requestData, err := db.GetOneDataKaryawanFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}
