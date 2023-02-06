package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	tp "github.com/henrylee2cn/teleport"
)

// ParseCouponUrlResp
type ParseCouponUrlResp struct {
	// 商品名称
	Itemname string `json:"itemName"`
	// 价格
	Price string `json:"price"`
	// 券后价
	Priceaftercoupon string `json:"priceAfterCoupon"`
	Volume           string `json:"volume"`
	Istm             int    `json:"isTM"`
	// 店铺名称
	Shoptitle string `json:"shopTitle"`
	Shopid    string `json:"shopId"`
	Shoplogo  string `json:"shopLogo"`
	// 优惠券id
	QuanActivityID string `json:"quan_activity_id"`
	// 优惠券金额
	QuanAmount      int    `json:"quan_amount"`
	QuanEndTime     string `json:"quan_end_time"`
	QuanTotalCount  int    `json:"quan_total_count"`
	QuanRemainCount int    `json:"quan_remain_count"`
	QuanStartTime   string `json:"quan_start_time"`
	QuanStartFee    int    `json:"quan_start_fee"`
	QuanInfo        string `json:"quan_info"`
	Pid             string `json:"pid"`
	// 商品id
	Itemid string `json:"itemId"`
	Jb     int    `json:"jb"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

// ParseCouponUrl 解析二合一链接，得到优惠券ID以及商品ID
func ParseCouponUrl(apiKey, couponUrl string) (*ParseCouponUrlResp, error) {
	params := url.Values{}
	params.Add("apikey", apiKey)
	params.Add("url", couponUrl)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest(
		"POST",
		"https://api.taokouling.com/tkl/TbkEhyjx",
		body,
	)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 请求client
	client := &http.Client{}

	// 发起请求
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	tp.Infof("resp-> %s", string(resp))

	// 解析数据
	var (
		parseResp *ParseCouponUrlResp
	)

	err = json.Unmarshal(resp, &parseResp)
	if err != nil {
		return nil, err
	}
	if parseResp.Code != 1 {
		return nil, fmt.Errorf("解析优惠券二合一链接出错, url-> %s", couponUrl)
	}
	return parseResp, nil
}
