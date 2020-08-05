package configs

import (
	"log"

	"github.com/joho/godotenv"
)

/*_Context Stores instance context */
var _Context *Driver

/*Driver Package structure */
type Driver struct {
	storages map[string]string
}

/*GetDriver Returns the instance of the package */
func GetDriver() *Driver {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Impossible to load configuration file")
	}
	_Context = new(Driver)
	_Context.storages = make(map[string]string)
	_ConfigsApp()
	_ConfigsDatabase()
	return _Context
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
