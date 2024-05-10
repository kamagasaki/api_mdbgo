package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

var ATSMDB = Config("ATSMONGODB")
var KaryawanColl = Config("KARYAWAN_COLL")
var AbsenColl = Config("ABSEN_COLL")
