package response

/*_ResponseBasic basic data output */
type _ResponseBasic struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
