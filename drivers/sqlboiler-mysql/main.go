package main

import (
	"github.com/myron-meng/sqlboiler/v4/drivers"
	"github.com/myron-meng/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
