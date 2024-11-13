package main

import (
	"github.com/myron-meng/sqlboiler/v4/drivers"
	"github.com/myron-meng/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
