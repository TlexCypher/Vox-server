package infrastructure

import (
	"Vox/db"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"Vox/internal/infrastructure/model"
	"log"

	"github.com/samber/lo"
)

type UserRepositoryImpl struct {
	db *db.DB
}

func NewUserRepositoryImpl(db *db.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func entityToModel(user entity.User) model.User {
	return model.User{
		ID:       user.IdVO().String(),
		Role:     user.RoleVO().Value(),
		Username: user.UsernameVO().Value(),
	}
}

/*Databseから戻ってきた値がValidationで落ちることはないのでlo.Mustでerrorを握りつぶす*/
func modelToEntity(user model.User) entity.User {
	return entity.User{
		Id:       lo.Must(valueobject.NewUserIdFromString(user.ID)),
		Role:     lo.Must(valueobject.NewRoleFromString(user.Role)),
		Username: lo.Must(valueobject.NewUsernameFromString(user.Username)),
	}
}

func (ur *UserRepositoryImpl) Save(e entity.User) error {
	m := entityToModel(e)
	ur.db.Conn.Create(&m)
	return nil
}

func (ur *UserRepositoryImpl) Update(e entity.User) error {
	m := entityToModel(e)
	ur.db.Conn.Save(&m)
	return nil
}

func (ur *UserRepositoryImpl) FindById(userId valueobject.UserId) (*entity.User, error) {
	m := model.User{
		ID: userId.String(),
	}
	result := ur.db.Conn.First(&m, userId.String())
	if result.Error != nil {
		log.Printf("Failed to find one user by userId. UserId:%v\n", userId.String())
		return nil, result.Error
	}
	return lo.ToPtr(modelToEntity(m)), nil
}

func (ur *UserRepositoryImpl) FindAll() (*[]entity.User, error) {
	return nil, nil
}

func (ur *UserRepositoryImpl) DeleteById(userId valueobject.UserId) error {
	u, err := ur.FindById(userId)
	if err != nil {
		log.Printf("Failed to delete user by userId. UserId:%v\n", userId.String())
		return err
	}
	ur.db.Conn.Delete(u)
	return nil
}
