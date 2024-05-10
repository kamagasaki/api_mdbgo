package router

import (
	"api/handler/karyawan"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/karyawan", karyawan.TambahKaryawan)
	api.Get("/karyawan", karyawan.GetAllDataKaryawan)
	api.Get("/karyawan/:username", karyawan.GetDataKaryawanByUname)

}
