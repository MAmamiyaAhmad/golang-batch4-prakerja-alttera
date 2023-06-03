package main

import (
	"prakerja4/configs"
	"prakerja4/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	e := routes.Init()
	e.Start(":8000")
}

