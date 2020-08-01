package models

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/database/orm"
)

/*Email return session email */
var Email string

/*IDUser return session user id */
var IDUser string

/*ORM retorna installs orm database */
var ORM orm.Driver
