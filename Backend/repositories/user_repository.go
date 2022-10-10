package repositories

import (
	"context"
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct{
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domains.UserRepository{
	return &userRepository{db}
}

func (ur *userRepository) InsertUser(u *models.User) error{
	ctx := context.Background()
	insertUserSql := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`
	if _, err := ur.db.Exec(ctx, insertUserSql, u.Username, string(u.Password), u.Role); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CheckUsername(username string) error{
	ctx := context.Background()
	checkUsernameSql := `SELECT * FROM users WHERE username = $1;`
	rows, err := ur.db.Query(ctx, checkUsernameSql, username);
	if err != nil {
		return err
	}
	u := new(models.User)
	for rows.Next(){
		if err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role); err != nil{
			return err
		}
	}
	if u.Username != "" {
		return fmt.Errorf("username already used")
	}
	return nil
}

func (ur *userRepository) FindUser(username string) (*models.User, error) {
	ctx := context.Background()
	u := new(models.User)
	checkUsernameSql := `SELECT * FROM users WHERE username = $1;`
	rows, err := ur.db.Query(ctx, checkUsernameSql, username);
	if err != nil {
		return u, err
	}
	for rows.Next(){
		if err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role); err != nil{
			return u, err
		}
	}
	return u, nil
}