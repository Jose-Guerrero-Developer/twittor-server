package configs

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

/*_Context stores the status of the configuration driver */
var _Context *Driver

/*Driver structure manages configurations */
type Driver struct {
	storages map[string]string
}

/*LoadDriver sets the configuration driver instance */
func (Controller *Driver) LoadDriver() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	_Context = new(Driver)
	_Context.storages = make(map[string]string)
	_ConfigsApp()
	_ConfigsDatabase()
	if err := recover(); err != nil {
		return errors.New("")
	}
	return nil
}

/*subscribe settings */
func subscribe(key string, value string, defaultValue string) {
	if len(key) <= 0 {
		log.Fatal("Length of the configuration key is required")
	}
	if len(value) >= 1 {
		_Context.storages[key] = value
		return
	}
	if len(defaultValue) <= 0 {
		log.Fatal("Default value length of the configuration is required")
	}
	_Context.storages[key] = defaultValue
}

/*Get returns the requested configuration */
func (Controller *Driver) Get(key string) string {
	value := _Context.storages[key]
	if len(value) <= 0 {
		log.Printf("Configuration " + key + " has no registered value or does not exist in configuration storage")
	}
	return value
}
