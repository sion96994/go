package coupon

import (
	"encoding/json"
	"net/url"

	tp "github.com/henrylee2cn/teleport"
	"github.com/sion96994/go/openapi/client"
)

// GetCouponInfo 获取优惠券信息
func GetCouponInfo(c *client.Client, arg *CouponGetParams) (*CouponGetResp, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Set("item_id", arg.ItemId)
	params.Set("activity_id", arg.ActivityId)
	params.Set("method", "taobao.tbk.coupon.get")

	// 2. 执行请求
	data, err := c.Execute(params)
	if err != nil {
		return nil, err
	}

	// 3. 判断是否有错误
	if err := client.IsErrorResponseExists(data); err != nil {
		return nil, err
	}

	// 4. 读取数据
	resp, err := data.Get("tbk_coupon_get_response").GetPath("data").MarshalJSON()
	if err != nil {
		return nil, err
	}
	tp.Infof("resp-> %v", string(resp))

	// 5. 解析响应
	var (
		couponInfo *CouponGetResp
	)
	if err = json.Unmarshal(resp, &couponInfo); err != nil {
		return nil, err
	}
	return couponInfo, nil
}
