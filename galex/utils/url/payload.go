package url

/*payload Load all basic settings */
func (Url *Driver) payload() {
	storage.Writer.Header().Set("Content-Type", "application/json")
}
