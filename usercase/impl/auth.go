package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	localErr "github.com/GunnNguyen/BasicAPI/errors"
	rInterfaces "github.com/GunnNguyen/BasicAPI/repository/interfaces"
	rModel "github.com/GunnNguyen/BasicAPI/repository/model"
	uInterfaces "github.com/GunnNguyen/BasicAPI/usercase/interfaces"
	"github.com/GunnNguyen/BasicAPI/usercase/model"
	hashpass "github.com/GunnNguyen/BasicAPI/utils/hash"
	jwt "github.com/GunnNguyen/BasicAPI/utils/jwt"
	"gorm.io/gorm"
)

type Auth struct {
	Repo          rInterfaces.Auth
	SecretKey     string
	TokenTimeLife time.Duration
}

func NewAuthUsercase(repo rInterfaces.Auth, secretKey string, tokenTimeLife time.Duration) uInterfaces.Auth {
	return &Auth{
		Repo:          repo,
		SecretKey:     secretKey,
		TokenTimeLife: tokenTimeLife,
	}
}

func (a Auth) SignUp(ctx context.Context, req model.User) error {
	//check account exists
	if _, err := a.Repo.CheckGmail(ctx, req.Gmail); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return localErr.Error{
			Code:   localErr.Code_ERR_DATA_NOT_FOUND,
			Detail: "gmail not found",
		}
	}
	//hash password
	hashPassword, err := hashpass.HashPassword(req.Password)
	if err != nil {
		return err
	}
	//create account
	if err := a.Repo.SignUp(ctx, rModel.User{
		Name:     req.Name,
		Age:      req.Age,
		Gmail:    req.Gmail,
		Password: hashPassword,
		Address:  req.Address,
	}); err != nil {
		return err
	}
	return nil
}

func (a Auth) SignIn(ctx context.Context, req model.Login) (*model.Reply, error) {
	//check gmail exists
	user, err := a.Repo.CheckGmail(ctx, req.Gmail)
	if err != nil {
		return nil, err
	}
	//check gamil

	//check pass
	if !hashpass.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("wrong password")
	}
	//generate accestoken
	accestoken, err := jwt.GenerateJwt(ctx, a.SecretKey, a.TokenTimeLife, user.Name, user.Gmail)
	if err != nil {
		return nil, err
	}
	//generate refreshtoken
	refreshtoken, err := jwt.GenerateJwt(ctx, a.SecretKey, a.TokenTimeLife+time.Minute, user.Name, user.Gmail)
	if err != nil {
		return nil, err
	}
	return &model.Reply{
		AccessToken:  accestoken,
		RefreshToken: refreshtoken,
	}, nil
}
