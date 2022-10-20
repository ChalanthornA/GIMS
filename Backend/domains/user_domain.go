package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type UserUseCase interface{
	RegisterAdmin(u *models.User, secret string) error
	Register(u *models.User) error
	SignIn(u *models.User) (*models.User, string, error)
	GenerateHashPassword(password string) (string, error)
	GenerateHashPasswordAndReplaceInUserModel(u *models.User) error
	RenameUsername(oldUsername, newUsername string) error
	UpdatePassword(username, password string) error
}

type UserRepository interface{
	InsertUser(u *models.User) error
	CheckUsername(username string) error
	FindUser(username string) (*models.User, error)
	QueryAllUser() ([]models.User, error)
	UpdateUsername(oldUsername, newUsername string) error
	UpdatePassword(username, newHashPassword string) error
}