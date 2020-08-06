package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

/*Driver Source of the package */
type Driver struct {
	storages map[string]string
}

/*packageStorage Store package status */
var packageStorage *Driver

/*EstablishDriver Set package driver */
func (Configuration *Driver) EstablishDriver() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	packageStorage = new(Driver)
	packageStorage.storages = make(map[string]string)
	return nil
}

/*Subscribe settings */
func Subscribe(key string, value string, defaultValue string) {
	log.Println("key: ", key, " | value: ", value, " | default: ", defaultValue)
	if len(key) <= 0 {
		log.Println("Length of the configuration key is required")
	}
	if len(value) >= 1 {
		packageStorage.storages[key] = value
		return
	}
	if len(defaultValue) <= 0 {
		log.Println("Default value length of the configuration is required")
	}
	packageStorage.storages[key] = defaultValue
}

/*Get Returns the requested configuration */
func (Configuration *Driver) Get(key string) string {
	value := packageStorage.storages[key]
	if len(value) <= 0 {
		log.Printf("Configuration " + key + " has no registered value or does not exist in configuration storage")
	}
	return value
}
