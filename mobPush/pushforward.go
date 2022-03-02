package mob_push_sdk

type PushForward struct {
	// 0 打开首页 1 link跳转 2 scheme 跳转
	NextType int `json:"nextType"`
	// 1 link跳转 moblink功能的的uri
	Url string `json:"url"`
	// 2 scheme moblink功能的的scheme
	Scheme string `json:"scheme"`
	// schema参数
	SchemeDataList []PushMap `json:"schemeDataList"`
}
