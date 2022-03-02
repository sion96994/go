package mob_push_sdk

const (
	// 安卓推送
	PUSH_PLAT_ANDROID = 1
	// iOS推送
	PUSH_PLAT_IOS = 2
)

const (
	// 通知消息
	MSG_TYPE_NOTICE = 1
	// 自定义消息
	MSG_TYPE_CUSTOM = 2
)

const (
	// ios测试
	IOS_ENV_TEST = 0
	// ios正式
	IOS_ENV_PROD = 1
)

type PushNotify struct {
	// 1 android;2 ios
	Plats []int `json:"plats"`
	// plat = 2下，0测试环境，1生产环境，默认1
	IosProduct int `json:"iosProduction"`
	// 离线消息保存时间，如若不传此参数，默认86400s（1天）
	OfflineSeconds int `json:"offlineSeconds"`
	// 推送内容, 如果内容长度超过厂商的限制, 则内容会被截断. vivo不允许纯表情
	Content string `json:"content"`
	// 如果不设置，则默认的通知标题为应用的名称。如果标题长度超过厂商的限制, 则标题会被截断. vivo不允许纯表情
	Title string `json:"title"`
	// 推送类型：1通知；2自定义
	Type int `json:"type"`

	// android通知消息, type=1, android
	AndroidNotify *AndroidNotify `json:"androidNotify"`
	// ios通知消息, type=1, ios
	IosNotify *IosNotify `json:"iosNotify"`

	// 是否是定时任务：0否，1是，默认0
	TaskCron int `json:"taskCron"`
	// 定时消息 发送时间， taskCron=1时必填 传入时间戳单位（毫秒 例如 1594277916000 ）
	TaskTime int64 `json:"taskTime"`
	// 每秒推送数量 只是趋势 默认0:不开启
	Speed int `json:"speed"`
	// 参数为1生效, 跳过在线设备, 不对在线设备做推送 1: 开启状态， 默认: 不开启,传入其他：不开启
	SkipOnline int `json:"skipOnline"`
	// 推送时，不走厂商通道1: 开启 默认: 其它,关闭
	SkipFactory int `json:"skipFactory"`

	// 自定义内容, type=2
	CustomNotify *CustomNotify `json:"customNotify"`

	// JSON格式 保留字段可参考下面附加参数示例(额外参数)
	ExtrasMapList []PushMap `json:"extrasMapList"`
}

type CustomNotify struct {
	CustomType  string `json:"customType"`
	CustomTitle string `json:"customTitle"`
}

// AndroidNotify android通知消息, type=1, android
type AndroidNotify struct {
	// 推送内容,配合style参数使用，style=0时候默认不生效 style=1时候部分机型可以生效覆盖，style=2时候传入图片链接部分低版本手机不支持 style=3时候对应传入图片
	Content []string `json:"content"`
	// 显示样式标识 0,"普通通知"，1,"BigTextStyle通知，点击后显示大段文字内容"，2,"BigPictureStyle，大图模式"，3,"横幅通知"，默认：0
	Style int `json:"style"`
	// 提醒类型： 1提示音；2震动；3指示灯，如果多个组合则对应编号组合如：12 标识提示音+震动
	Warn string `json:"warn"`
	// 自定义声音
	Sound string `json:"sound"`
	// 附带小图标的推送(icon和image只能二选一,都填只会传icon中的数据) （目前客户端版本暂不支持）
	Icon string `json:"icon"`
	// 推送大图标的url地址(icon和image只能二选一,都填只会传icon中的数据)（透传消息不支持）
	Image string `json:"image"`

	// 厂商channelId
	AndroidChannelId string `json:"androidChannelId"`
	// 1.表示将当前app的角标设置成androidBadge中的值 2.表示将当前app的角标加上androidBadge中的值（透传消息不支持）
	AndroidBadgeType string `json:"androidBadgeType"`
	// 需要设置/累加的角标的值
	AndroidBadge string `json:"androidBadge"`
}

// IosNotify ios通知消息, type=1, ios
type IosNotify struct {
	// 角标
	Badge string `json:"badge"`
	// badge类型, 1:绝对值 不能为负数，2增减(正数增加，负数减少)，减到0以下会设置为0
	BadgeType int `json:"badgeType"`
	// apns的category字段，只有IOS8及以上系统才支持此参数推送
	CateGory string `json:"category"`
	// APNs通知，通过这个字段指定声音。默认为default，即系统默认声音。 如果设置为空值，则为静音。如果设置为特殊的名称，则需要你的App里配置了该声音才可以正常。
	Sound string `json:"sound"`
	// 副标题
	SubTitle string `json:"subtitle"`
	// 如果只携带content-available: 1,不携带任何badge，sound 和消息内容等参数， 则可以不打扰用户的情况下进行内容更新等操作即为“Silent Remote Notifications”
	SlientPush int `json:"slientPush"`
	// 将该键设为 1 则表示有新的可用内容。带上这个键值，意味着你的 App 在后台启动了或恢复运行了，application:didReceiveRemoteNotification:fetchCompletionHandler:被调用了
	ContentAvailable int `json:"contentAvailable"`
	// 需要在附加字段中配置相应参数
	MutableContent int `json:"mutableContent"`
	// ios富文本 0 无 ； 1 图片 ；2 视频 ；3 音频
	AttachmentType int `json:"attachmentType"`
	// ios富文本内容
	Attachment string `json:"attachment"`
}

type PushMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (push *Push) setTitle(title string) *Push {
	push.PushNotify.Title = title
	return push
}

func (push *Push) setContent(content string) *Push {
	push.PushNotify.Content = content
	return push
}

func (push *Push) setPlats(plats []int) *Push {
	push.PushNotify.Plats = plats
	return push
}

func (push *Push) setCustomNotify(customNotify CustomNotify) *Push {
	push.PushNotify.CustomNotify = &customNotify
	return push
}

func (push *Push) setAndroidNotify(androidNotify AndroidNotify) *Push {
	push.PushNotify.AndroidNotify = &androidNotify
	return push
}

func (push *Push) setIosNotify(iosNotify IosNotify) *Push {
	push.PushNotify.IosNotify = &iosNotify
	return push
}
