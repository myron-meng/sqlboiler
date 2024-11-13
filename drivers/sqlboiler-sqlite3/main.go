package main

import (
	"github.com/myron-meng/sqlboiler/v4/drivers"
	"github.com/myron-meng/sqlboiler/v4/drivers/sqlboiler-sqlite3/driver"
)

func main() {
	drivers.DriverMain(&driver.SQLiteDriver{})
}
