package dao

import (
	"context"
	"github.com/zhaoxlchn/gin-basic-structure/internal/data"
	"github.com/zhaoxlchn/gin-basic-structure/internal/entity"
	"github.com/zhaoxlchn/gin-basic-structure/internal/service"
)

type userDao struct {
	data *data.Data
}

func NewUserDao(data *data.Data) service.UserDao {
	return &userDao{data: data}
}

func (u userDao) GetById(ctx context.Context, id int64) (info *entity.User, err error) {
	info = &entity.User{}
	err = u.data.Db.WithContext(ctx).Where("user_id = ?", id).First(info).Error
	return
}
