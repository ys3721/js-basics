package reflectbasic

import (
	"fmt"
	"reflect"
)

type IRoleToken interface {
	GetUserId() int64
	SetUserId(int64)
}

type BasicRoleToken struct {
	UserId int64
}

func (b *BasicRoleToken) GetUserId() int64 {
	return b.UserId
}
func (b *BasicRoleToken) SetUserId(userId int64) {
	b.UserId = userId
}

type SecureRoleToken struct {
	UserId  int64
	Session string
}

func (t *SecureRoleToken) GetUserId() int64 {
	return t.UserId
}

func (t *SecureRoleToken) SetUserId(userId int64) {
	t.UserId = userId
}

type Config struct {
	TypeOfRoleToken reflect.Type
}

func doReflectTypeNewMain() {
	cfg := Config{
		TypeOfRoleToken: reflect.TypeOf((*BasicRoleToken)(nil)).Elem(),
	}
	newInstance := reflect.New(cfg.TypeOfRoleToken).Interface()

	token := newInstance.(IRoleToken)
	/*	token, ok := newInstance.(IRoleToken)
		if !ok {
			fmt.Println("Type assertion successful")
		}*/
	token.SetUserId(47)
	fmt.Println(token)
	fmt.Println("用户 ID:", token.GetUserId())
}
