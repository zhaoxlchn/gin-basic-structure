package service

import (
	"context"
	"github.com/zhaoxlchn/gin-basic-structure/internal/entity"
	"github.com/zhaoxlchn/gin-basic-structure/internal/schema"
)

type UserDao interface {
	GetById(ctx context.Context, id int64) (info *entity.User, err error)
}

type UserService struct {
	userDao UserDao
}

func NewUserService(userDao UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s UserService) GetUser(ctx context.Context, id int64) (info *schema.UserInfoRes, err error) {
	res, err := s.userDao.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	info = &schema.UserInfoRes{}
	info.UserID = res.UserID
	info.UserNickname = res.UserNickname
	info.UserHeader = res.UserHeader
	return
}
