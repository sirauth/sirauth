package siridentity

type ManageIdentity interface {
	CreateIdentity(identity Identity) error
	GetIdentity(email string) (Identity, error)
	UpdateIdentity(identity Identity) error
	DeleteIdentity(email string) error
}

type Identity struct {
	ID       int64
	Email    string
	Password string
}
