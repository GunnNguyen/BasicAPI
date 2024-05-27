package impl

import (
	"context"
	"net/http"

	"github.com/GunnNguyen/BasicAPI/models"
	"github.com/GunnNguyen/BasicAPI/usercase/interfaces"
	"github.com/GunnNguyen/BasicAPI/usercase/model"
	"github.com/labstack/echo/v4"
)

type Auth struct{
	usercase interfaces.Auth
}

func RegisterAuthRouter(g *echo.Group,useCase interfaces.Auth){
	a := Auth{
		usercase: useCase,
	}
	g.POST("/signup", a.CreateUser)
	g.POST("/signin", a.LoginUser)
}

func (a Auth) CreateUser(e echo.Context) error{
	var (req models.User)
	if err := e.Bind(&req); err != nil{
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	if err := req.Validate(); err != nil{
		return e.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	if err := a.usercase.SignUp(context.Background(),model.User{
		Gmail: req.Gmail,
		Name: req.Name,
		Age: req.Age,
		Address: req.Address,
		Password: req.Password,
	}); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "ok",
	})
}

func (a Auth) LoginUser(e echo.Context) error{
	var (req models.User)
	if err := e.Bind(&req); err != nil{
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	if err := req.Validate(); err != nil{
		return e.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	reply, err :=  {
		Gmail:    req.Gmail,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
return e.JSON(http.StatusOK, models.LoginReply{
	AccessToken: r,
})
}