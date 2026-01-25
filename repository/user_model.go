package repository

import "github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"

type userModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (u userModel) toDomain() (domain.User, error) {
	userID := domain.UserID(u.ID)
	userName, err := domain.NewUserName(u.Name)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.NewUser(userID, userName)

	return user, nil
}
