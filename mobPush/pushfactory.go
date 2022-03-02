package mob_push_sdk

// PushFactoryExtra
type PushFactoryExtra struct {
	// 小米厂商特殊配置
	XiaomiExtra *XiaomiExtra `json:"xiaomiExtra,omitempty"`
	// OPPO厂商特殊配置
	OppoExtra *OppoExtra `json:"oppoExtra,omitempty"`
	// VIVO厂商特殊配置
	VivoExtra *VivoExtra `json:"vivoExtra,omitempty"`
}

type (
	XiaomiExtra struct {
		ChannelId string `json:"channel_id"`
	}

	OppoExtra struct {
		ChannelId string `json:"channel_id"`
	}

	VivoExtra struct {
		// VIVO消息类型 0：运营类型消息 ；1：系统类型消息
		Classification string `json:"classification"`
	}
)
