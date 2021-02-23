/**
 * @Author: Huyantian
 * @Date: 2021/2/23 下午7:20
 */

package filedao

import (
	"context"
	"todo/internal/model/user"
)

type UserDao struct {
}

func (u UserDao) Create(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (u UserDao) GetByName(ctx context.Context, userName string) (user.User, error) {
	panic("implement me")
}
