// Code generated by entc, DO NOT EDIT.

package realmoauthclient

const (
	// Label holds the string label denoting the realmoauthclient type in the database.
	Label = "realm_oauth_client"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRealmID holds the string denoting the realm_id field in the database.
	FieldRealmID = "realm_id"
	// FieldClientSecret holds the string denoting the client_secret field in the database.
	FieldClientSecret = "client_secret"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldRedirectUrls holds the string denoting the redirect_urls field in the database.
	FieldRedirectUrls = "redirect_urls"
	// Table holds the table name of the realmoauthclient in the database.
	Table = "clients"
)

// Columns holds all SQL columns for realmoauthclient fields.
var Columns = []string{
	FieldID,
	FieldRealmID,
	FieldClientSecret,
	FieldClientID,
	FieldRedirectUrls,
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