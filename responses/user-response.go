package responses

type UserResponse struct {
	StatusCode int                    `json:"statusCode"` // these are struct tags. Used to reformat the JSON response returned by the API
	Status     string                 `json:"status"`
	ErrorCode  string                 `json:"errorCode,omitempty"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data,omitempty"`
}