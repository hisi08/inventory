// Code generated by ent, DO NOT EDIT.

package staff

const (
	// Label holds the string label denoting the staff type in the database.
	Label = "staff"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the staff in the database.
	Table = "staffs"
)

// Columns holds all SQL columns for staff fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)