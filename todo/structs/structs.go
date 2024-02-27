package structs

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}
