package helpers

/*Driver Source of the package */
type Driver struct {
	collection string
}

/*EstablishDriver Set package driver */
func EstablishDriver(collection string) *Driver {
	Helper := new(Driver)
	Helper.collection = collection
	return Helper
}
