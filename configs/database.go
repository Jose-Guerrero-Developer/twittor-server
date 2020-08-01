package configs

import "os"

/*_ConfigsDatabase export configurations database */
func _ConfigsDatabase() {
	/*Database name */
	subscribe("DATABASE_NAME", os.Getenv("DATABASE_NAME"), "estudio")
	/*Database user name */
	subscribe("DATABASE_USERNAME", os.Getenv("DATABASE_USERNAME"), "root")
	/*Database password */
	subscribe("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"), "root")
}
