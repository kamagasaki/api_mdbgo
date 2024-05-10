package config

import (
	"os"
)

var ATSMDB = os.Getenv("ATSMONGODB")
var KaryawanColl = os.Getenv("KARYAWAN_COLL")
var AbsenColl = os.Getenv("ABSEN_COLL")
