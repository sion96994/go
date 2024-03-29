package logic

import (
	"fmt"

	micro "github.com/xiaoenai/tp-micro"

	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
	"github.com/sion96994/go/tp-template/accessToken"
	"github.com/sion96994/go/tp-template/gateway/sign"
	"github.com/xiaoenai/tp-micro/gateway/types"
)

const (
	// 授权通过key
	PASS_KEY = "_pass"
	// 授权网关名称key
	AUTH_NAME_KEY = "auth_name_"
	// 签名参数key
	SIGN_KEY = sign.SIGN_KEY
	// uid
	UID_KEY       = "_uid"
	DEVICE_ID_KEY = "_device_id"
)

var (
	// authFuncLib 授权查询库
	authFuncLib = make(map[string]accessToken.Builder)
)

// RegBuilder 注册构建AccessToken的函数
// 注：若注册AccessToken的构建函数为nil，或重复注册，均会发生panic
func RegBuilder(authName string, fn accessToken.Builder) {
	if SupportAuthName(authName) {
		tp.Fatalf("重复注册获取AccessToken的函数：%s", authName)
	}
	if fn == nil {
		tp.Fatalf("获取AccessToken的函数不能有空值：%s", authName)
	}
	authFuncLib[authName] = fn
}

// SupportAuthName 判断是否支持指定授权
func SupportAuthName(authName string) bool {
	_, ok := authFuncLib[authName]
	return ok
}

// LookupBuilder
func LookupBuilder(authName string) (accessToken.Builder, *tp.Rerror) {
	fn, ok := authFuncLib[authName]
	if !ok {
		return nil, tp.NewRerror(tp.CodeUnauthorized, "不支持的授权类型", "")
	}
	return fn, nil
}

// authFunc
func authFunc(authInfo string) (types.AccessToken, *tp.Rerror) {
	args := utils.AcquireArgs()
	defer utils.ReleaseArgs(args)
	args.Parse(authInfo)

	var (
		notallowParams []string
	)
	args.VisitAll(func(k, _ []byte) {
		if len(k) > 0 && k[0] == '_' {
			notallowParams = append(notallowParams, goutil.BytesToString(k))
		}
	})
	if len(notallowParams) > 0 {
		return nil, micro.RerrInvalidParameter.Copy().SetReason(
			fmt.Sprintf("Query parameter keys cannot contain an underscore prefix: %+v", notallowParams),
		)
	}

	// 网关授权名称
	authName := goutil.BytesToString(args.Peek(AUTH_NAME_KEY))
	fn, rerr := LookupBuilder(authName)
	if rerr != nil {
		return nil, rerr
	}
	t, rerr := fn(args)
	if rerr != nil {
		return nil, rerr
	}
	if t == nil {
		return nil, tp.NewRerror(tp.CodeUnauthorized, "不支持的授权类型", "")
	}

	// 标记已通过网关验证
	t.AddedQuery().Set(PASS_KEY, "gw")
	// 添加用户信息及设备标识
	t.AddedQuery().Set(UID_KEY, t.Uid())
	return t, rerr
}
