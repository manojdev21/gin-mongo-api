package validations

import (
	"gin-mongo-api/logger"
	"gin-mongo-api/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate = validator.New()

func Validate(c *gin.Context, obj interface{}) (responses.APIResponse, error) {
	//validate the request body
	err := c.BindJSON(obj)
	if err != nil {
		return responses.BuildAPIResponse(responses.FAILURE, err.Error()), err
	}

	//use the validator library to validate required fields
	err = validate.Struct(obj)
	if err != nil {
		return responses.BuildAPIResponse(responses.FAILURE, err.Error()), err
	}

	logger.InfoLogger.Println("Validated request body and required fields")

	return responses.APIResponse{}, nil
}
