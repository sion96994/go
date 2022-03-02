// Code generated by 'micro gen' command.
// DO NOT EDIT!

package api

import (
	tp "github.com/henrylee2cn/teleport"

	"github.com/sion96994/go/tp-template/account/args"
	"github.com/sion96994/go/tp-template/account/logic"
)

// V1_User controller
type V1_User struct {
	tp.PullCtx
}

// 增加用户
func (v *V1_User) Set(arg *args.SetUserArgsV1) (*args.SetUserResultV1, *tp.Rerror) {
	return logic.V1_User_Set(v.PullCtx, arg)
}

// 根据ID获取user
func (v *V1_User) GetById(arg *args.GetUserByIdArgsV1) (*args.GetUserByIdResultV1, *tp.Rerror) {
	return logic.V1_User_GetById(v.PullCtx, arg)
}