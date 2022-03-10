package controllers

import (
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"gin-mongo-api/services"
	"gin-mongo-api/validations"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if res, err := validations.Validate(c, &user); err != nil {
			c.JSON(res.StatusCode, res)
			return
		}

		createdUser, err := services.CreateUser(user)
		if err != nil {
			res := responses.BuildAPIResponse(responses.FAILURE, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}
		res := responses.BuildAPIResponse(responses.SUCCESS, "User created successfully", map[string]interface{}{"data": createdUser})
		c.JSON(res.StatusCode, res)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		userId, _ := primitive.ObjectIDFromHex(paramId)

		fetchedUser, err := services.GetUser(userId)
		if err != nil {
			res := responses.BuildAPIResponse(responses.FAILURE, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}
		res := responses.BuildAPIResponse(responses.SUCCESS, "User fetched successfully", map[string]interface{}{"data": fetchedUser})
		c.JSON(res.StatusCode, res)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if res, err := validations.Validate(c, &user); err != nil {
			c.JSON(res.StatusCode, res)
			return
		}

		err := services.UpdateUser(user)
		if err != nil {
			res := responses.BuildAPIResponse(responses.FAILURE, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}

		res := responses.BuildAPIResponse(responses.SUCCESS, "User updated successfully")
		c.JSON(res.StatusCode, res)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		userId, _ := primitive.ObjectIDFromHex(paramId)

		err := services.DeleteUser(userId)
		if err != nil {
			res := responses.BuildAPIResponse(responses.FAILURE, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}

		res := responses.BuildAPIResponse(responses.SUCCESS, "User deleted successfully")
		c.JSON(res.StatusCode, res)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, totalRecords, err := services.GetAllUsers()
		if err != nil {
			res := responses.BuildAPIResponse(responses.FAILURE, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}

		data := map[string]interface{}{
			"data":         users,
			"totalRecords": totalRecords,
			"drawCount":    len(users),
		}
		res := responses.BuildAPIResponse(responses.SUCCESS, "Users fetched successfully", data)
		c.JSON(res.StatusCode, res)
	}
}
