package response

/*Driver Package structure */
type Driver struct{}

/*GetDriver Returns the instance of the package */
func GetDriver() *Driver {
	return new(Driver)
}
