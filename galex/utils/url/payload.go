package url

/*payload Load all basic settings */
func (Controller *Driver) payload() {
	Controller.setHeaders()
}

/*setHeaders Sets the response headers*/
func (Controller *Driver) setHeaders() {
	_Context.Writer.Header().Set("Content-Type", "application/json")
}
