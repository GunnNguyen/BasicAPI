package interfaces

import (
	"context"

	"github.com/GunnNguyen/BasicAPI/usercase/model"
)

type Auth interface {
	SignUp(context.Context, model.User) error
	SignIn(context.Context, model.Login) (*model.Reply, error)
}
