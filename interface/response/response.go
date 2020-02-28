package response

type HttpResponse struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Response interface{} `json:"response,omitempty"`
	Errors   []string    `json:"errors,omitempty"`
}
