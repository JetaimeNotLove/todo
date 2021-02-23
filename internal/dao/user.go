/**
 * @Author: Huyantian
 * @Date: 2021/2/23 下午7:15
 */

package dao

import (
	"context"
	"todo/internal/model/user"
)

var userDao UserDao

func SetUser(dao UserDao) {
	userDao = dao
}

func User() UserDao {
	return userDao
}

type UserDao interface {
	Create(ctx context.Context, user user.User) error
	GetByName(ctx context.Context, userName string) (user.User, error)
}
