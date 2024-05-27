package interfaces

import (
	"context"

	"github.com/GunnNguyen/BasicAPI/repository/model"
)

type Auth interface {
	SignUp(context.Context, model.User) error
	CheckGmail(context.Context, string) (model.User, error)
}

type ImlpAuth interface {
	Auth() Auth
}
