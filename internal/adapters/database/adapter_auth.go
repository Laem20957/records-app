package database

import (
	"context"
	"fmt"
	"records-app/internal/adapters/database/schemas"

	"github.com/jinzhu/gorm"
)

type DatabaseAuthORM struct {
	DB *gorm.DB
}

func AdapterGetAuth(db *gorm.DB) *DatabaseAuthORM {
	return &DatabaseAuthORM{DB: db}
}

func (d *DatabaseAuthORM) GetUserDB(ctx context.Context, userId int) (schemas.Users, error) {
	var user schemas.Users

	if err := db.Table(fmt.Sprintf("records_app.%s", usersTable)).Where("id = ?", userId).First(&user).Error; err != nil {
		logs.Error(err)
	}
	return user, nil
}

func (d *DatabaseAuthORM) GetTokenDB(ctx context.Context, tokenId int) (schemas.Tokens, error) {
	var token schemas.Tokens

	if err := db.Table(fmt.Sprintf("records_app.%s", refreshTokensTable)).Where("id = ?", tokenId).First(&token).Error; err != nil {
		logs.Error(err)
	}
	return token, nil
}

func (d *DatabaseAuthORM) CreateUserDB(ctx context.Context) (int, error) {
	var user schemas.Users

	if err := db.Table(fmt.Sprintf("records_app.%s", usersTable)).Create(&user).Error; err != nil {
		logs.Error(err)
	}
	return user.ID, nil
}

func (d *DatabaseAuthORM) CreateTokenDB(ctx context.Context) (int, error) {
	var token schemas.Tokens

	if err := db.Table(fmt.Sprintf("records_app.%s", refreshTokensTable)).Create(&token).Error; err != nil {
		logs.Error(err)
	}
	return token.ID, nil
}
