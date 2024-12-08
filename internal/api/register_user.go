package api

import (
	di "Vox/cmd/wire-gen"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"Vox/internal/usecase"
	"Vox/openapi"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func RegisterUser(ctx echo.Context, username, role string) error {
	usernameVO, err := valueobject.NewUsernameFromString(username)
	if err != nil {
		slog.Error("RegisterUser", "Failed to register user, because of username validation.", err)
	}

	roleVO, err := valueobject.NewRoleFromString(role)
	if err != nil {
		slog.Error("RegisterUser", "Failed to register user, because of role validation.", err)
		return err
	}

	user, err := entity.NewUser(roleVO, usernameVO)
	if err != nil {
		slog.Error("RegisterUser", "Failed to register user, because of user validation.", err)
		return err
	}
	u := usecase.RegisterUser{
		UserRepository: lo.ToPtr(di.InitUserRepository()),
	}
	input := usecase.RegisterUserInput{
		User: user,
	}
	output, err := u.Do(ctx, input)
	if err != nil {
		slog.Error("RegisterUser", "Failed to execute usecase. Usecase: RegisterUser", err)
		return err
	}
	return ctx.JSON(
		http.StatusCreated,
		openapi.CreateUserResponse{
			UserId: output.Id,
		},
	)
}
