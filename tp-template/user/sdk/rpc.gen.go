// Code generated by 'micro gen' command.
// DO NOT EDIT!

package sdk

import (
	"fmt"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"
	"github.com/xiaoenai/tp-micro/model/etcd"

	"github.com/sion96994/go/tp-template/user/args"
)

var _ = fmt.Sprintf
var client *micro.Client

// Init initializes client with configs.
func Init(cliConfig micro.CliConfig, etcdConfing etcd.EasyConfig) {
	client = micro.NewClient(
		cliConfig,
		discovery.NewLinker(etcdConfing),
	)
}

// InitWithClient initializes client with specified object.
func InitWithClient(cli *micro.Client) {
	client = cli
}

// Add handler
func V1_User_Add(arg *args.AddUserArgsV1, setting ...socket.PacketSetting) (*args.AddUserResultV1, *tp.Rerror) {
	result := new(args.AddUserResultV1)
	rerr := client.Pull("/user/v1/user/add", arg, result, setting...).Rerror()
	return result, rerr
}

// 获取用户
func V1_User_GetById(arg *args.GetUserByIdArgsV1, setting ...socket.PacketSetting) (*args.GetUserByIdResultV1, *tp.Rerror) {
	result := new(args.GetUserByIdResultV1)
	setting = append(setting, tp.WithQuery("id", fmt.Sprintf("%v", arg.Id)))
	rerr := client.Pull("/user/v1/user/get_by_id", arg, result, setting...).Rerror()
	return result, rerr
}

// Get handler
func V1_User_Get(arg *args.GetUserArgsV1, setting ...socket.PacketSetting) (*args.GetUserResultV1, *tp.Rerror) {
	result := new(args.GetUserResultV1)
	rerr := client.Pull("/user/v1/user/get", arg, result, setting...).Rerror()
	return result, rerr
}