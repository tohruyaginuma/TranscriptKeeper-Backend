package domain

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByGoogleID(googleID GoogleID) (*User, error)
	ExistsByGoogleID(googleID GoogleID) (bool, error)
}
