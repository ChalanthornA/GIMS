package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type UserUseCase interface{
	RegisterAdmin(u *models.User, secret string) error
	SignIn(u *models.User) (string, error)
}

type UserRepository interface{
	InsertUser(u *models.User) error
	CheckUsername(username string) error
	FindUser(username string) (*models.User, error)
}