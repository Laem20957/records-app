package repository

import (
	"context"
	"fmt"

	"records-app/internal/domain"

	gorm "github.com/jinzhu/gorm"
)

type AuthPostgreSQL struct {
	DB *gorm.DB
}

func RepositoryGetAuth(db *gorm.DB) *AuthPostgreSQL {
	return &AuthPostgreSQL{DB: db}
}

func (repo *AuthPostgreSQL) GetUserDB(ctx context.Context, userId int) (domain.Users, error) {
	var user domain.Users

	if err := db.Table(fmt.Sprintf("records_app.%s", usersTable)).Where("id = ?", userId).First(&user).Error; err != nil {
		logs.Error(err)
	}
	return user, nil
}

func (repo *AuthPostgreSQL) GetTokenDB(ctx context.Context, tokenId int) (domain.Tokens, error) {
	var token domain.Tokens

	if err := db.Table(fmt.Sprintf("records_app.%s", refreshTokensTable)).Where("id = ?", tokenId).First(&token).Error; err != nil {
		logs.Error(err)
	}
	return token, nil
}

func (repo *AuthPostgreSQL) CreateUserDB(ctx context.Context) (int, error) {
	var user domain.Users

	if err := db.Table(fmt.Sprintf("records_app.%s", usersTable)).Create(&user).Error; err != nil {
		logs.Error(err)
	}
	return user.ID, nil
}

func (repo *AuthPostgreSQL) CreateTokenDB(ctx context.Context) (int, error) {
	var token domain.Tokens

	if err := db.Table(fmt.Sprintf("records_app.%s", refreshTokensTable)).Create(&token).Error; err != nil {
		logs.Error(err)
	}
	return token.ID, nil
}
