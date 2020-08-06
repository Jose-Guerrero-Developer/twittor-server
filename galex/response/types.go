package response

/*responseBasic basic data output */
type responseBasic struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
