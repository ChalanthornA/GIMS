package usecases

import (
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct{
	userRepo domains.UserRepository
}

func NewUserUseCase(ur domains.UserRepository) domains.UserUseCase{
	return &userUseCase{
		userRepo: ur,
	}
}

func (uu *userUseCase) GenerateHashPassword(u *models.User) error{
	p := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (uu *userUseCase) RegisterAdmin(u *models.User, secret string) error{
	if secret != "BaBaBuBu"{
		return fmt.Errorf("invalid secret")
	}
	if err := uu.userRepo.CheckUsername(u.Username); err != nil{
		return err
	}
	u.Role = "admin"
	err := uu.GenerateHashPassword(u)
	if err != nil {
		return err
	}
	if err := uu.userRepo.InsertUser(u); err != nil{
		return err
	}
	return nil
}

func (uu *userUseCase) Register(u *models.User) error{
	if err := uu.userRepo.CheckUsername(u.Username); err != nil{
		return err
	}
	err := uu.GenerateHashPassword(u)
	if err != nil {
		return err
	}
	if err := uu.userRepo.InsertUser(u); err != nil{
		return err
	}
	return nil
}