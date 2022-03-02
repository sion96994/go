package mob_push_sdk

type PushOperator struct {
	// 运营保障 消息 修改类型： 1 取消 2 替换 3 撤回
	DropType int `json:"dropType"`
	// 推送任务的唯一id
	DropId string `json:"dropId"`
}
