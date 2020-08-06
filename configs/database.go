package configs

import (
	"os"
)

/*Database Configs Database */
func (Configs *Driver) Database() {
	/*Database HOST */
	subscribe("DATABASE_HOST", os.Getenv("DATABASE_HOST"), "estudios")
	/*Database NAME */
	subscribe("DATABASE_NAME", os.Getenv("DATABASE_NAME"), "twittor")
	/*Database USERNAME */
	subscribe("DATABASE_USERNAME", os.Getenv("DATABASE_USERNAME"), "root")
	/*Database PASSWORD */
	subscribe("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"), "root")
}
