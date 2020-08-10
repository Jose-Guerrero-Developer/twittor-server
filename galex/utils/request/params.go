package request

import (
	"strconv"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/url"
)

/*AddHeader Add header */
func (Controller *Driver) AddHeader(name string, value string) {
	var GalexURL url.Driver
	GalexURL.AddHeader(name, value)
}

/*GetPage Returns the page number to be displayed */
func (Controller *Driver) GetPage() int64 {
	var GalexURL url.Driver
	page, _ := strconv.ParseInt(GalexURL.GetRequest().Header.Get("Current-Page"), 10, 64)
	if page < 1 {
		page = 1
	}
	return page
}

/*GetCount Returns the number of records to be displayed */
func (Controller *Driver) GetCount() int64 {
	var GalexURL url.Driver
	count, _ := strconv.ParseInt(GalexURL.GetRequest().Header.Get("Total-Count"), 10, 64)
	if count < 1 {
		count = 10
	}
	return count
}
