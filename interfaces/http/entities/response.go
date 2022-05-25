package entities

type Response struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ResponseFormater(Status uint, data map[string]interface{}) Response {
	return Response{
		Status,
		GetResponseCodeMessage(Status),
		data["meta"],
		data["data"],
		data["error"],
	}
}
