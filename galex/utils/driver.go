package utils

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/bcrypt"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/pagination"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*Driver structure manages utils */
type Driver struct {
	URL        *url.Driver
	Bcrypt     *bcrypt.Driver
	Pagination *pagination.Driver
}

/*LoadDriver Set initial settings for utils */
func LoadDriver() *Driver {
	return new(Driver)
}
