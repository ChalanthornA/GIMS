package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type UserUseCase interface{
	RegisterAdmin(u *models.User, secret string) error
	Register(u *models.User) error
	SignIn(u *models.User) (*models.User, string, error)
	GenerateHashPassword(u *models.User) error
}

type UserRepository interface{
	InsertUser(u *models.User) error
	CheckUsername(username string) error
	FindUser(username string) (*models.User, error)
}