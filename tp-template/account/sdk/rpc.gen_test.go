// Code generated by 'micro gen' command.
// DO NOT EDIT!

package sdk_test

import (
	"encoding/json"
	"fmt"

	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/model/etcd"

	"github.com/sion96994/go/tp-template/account/args"
	"github.com/sion96994/go/tp-template/account/sdk"
)

func init() {
	sdk.Init(
		micro.CliConfig{
			Failover:        3,
			HeartbeatSecond: 4,
		},
		etcd.EasyConfig{
			Endpoints: []string{"http://127.0.0.1:2379"},
		},
	)
}

func toJsonBytes(i interface{}) []byte {
	b, _ := json.MarshalIndent(i, "", "  ")
	return b
}

func ExampleV1_User_Set() {
	result, rerr := sdk.V1_User_Set(&args.SetUserArgsV1{})
	if rerr != nil {
		tp.Errorf("V1_User_Set: rerr: %s", toJsonBytes(rerr))
	} else {
		tp.Infof("V1_User_Set: result: %s", toJsonBytes(result))
	}
	fmt.Printf("")
	// Output:
}

func ExampleV1_User_GetById() {
	result, rerr := sdk.V1_User_GetById(&args.GetUserByIdArgsV1{})
	if rerr != nil {
		tp.Errorf("V1_User_GetById: rerr: %s", toJsonBytes(rerr))
	} else {
		tp.Infof("V1_User_GetById: result: %s", toJsonBytes(result))
	}
	fmt.Printf("")
	// Output:
}
