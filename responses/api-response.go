package responses

import (
	"gin-mongo-api/exceptions"
	"gin-mongo-api/logger"
	"net/http"
)

const (
	SUCCESS = "SUCCESS"
	FAILURE = "FAILURE"
)

type APIResponse struct {
	StatusCode int                    `json:"statusCode"` // these are struct tags. Used to reformat the JSON response returned by the API
	Status     string                 `json:"status"`
	ErrorCode  string                 `json:"errorCode,omitempty"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data,omitempty"`
}

func failureAPIResponse(err string) APIResponse {
	response := APIResponse{}

	response.Status = FAILURE
	response.Message = exceptions.ErrorMessage(err)

	if response.Message != "" {
		response.StatusCode = http.StatusBadRequest
		response.ErrorCode = err
	} else {
		response.StatusCode = http.StatusInternalServerError
		response.Message = err
	}
	logger.FailureLogger.Println(response.ErrorCode, response.Message)

	return response
}

func BuildAPIResponse(status string, message string, data ...map[string]interface{}) APIResponse {
	if (status == FAILURE) {
		return failureAPIResponse(message)
	}
	response := APIResponse{}
	response.StatusCode = http.StatusOK
	response.Status = status
	response.Message = message
	if len(data) == 1 {
		response.Data = data[0]
	}
	
	return response
}
