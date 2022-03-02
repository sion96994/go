package mob_push_sdk

type Push struct {
	// 自定义任务id（用户自自定义生成且唯一、不能重复）
	WorkNo string `json:"workno"`
	// 枚举值 webapi
	Source string `json:"source"`
	// appkey
	Appkey string `json:"appkey"`
	// 推送目标
	PushTarget PushTarget `json:"pushTarget"`
	// 推送展示细节配置
	PushNotify PushNotify `json:"pushNotify"`
	// 运营保障相关配置, 默认有个空对象
	PushOperator *PushOperator `json:"pushOperator,omitempty"`
	// link 相关打开配置
	PushForward *PushForward `json:"pushForward,omitempty"`

	// 厂商特殊配置
	PushFactoryExtra *PushFactoryExtra `json:"pushFactoryExtra"`
}

func NewPushModel(appKey string) *Push {
	push := &Push{}
	push.Appkey = appKey
	push.getDefaultPushModel()
	return push
}

func (push *Push) getDefaultPushModel() *Push {
	push.Source = "webapi"
	push.PushTarget.TagsType = "1"

	push.PushNotify.Plats = []int{1, 2}
	push.PushNotify.Type = 1
	push.PushNotify.IosProduct = 1
	push.PushNotify.OfflineSeconds = 3600

	return push
}

func (push *Push) setWorkno(workno string) *Push {
	push.WorkNo = workno
	return push
}
