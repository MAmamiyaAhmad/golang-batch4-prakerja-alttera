package controllers

import (
	"net/http"
	"prakerja4/configs"
	"prakerja4/models"

	"github.com/labstack/echo/v4"
)

func InsertUserController(c echo.Context) error {
	var userInput models.User
	c.Bind(&userInput)

	result := configs.DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "success",
		Data: userInput,
	})
}

func GetUserController(c echo.Context) error {
	var users []models.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data: users,
	})
}