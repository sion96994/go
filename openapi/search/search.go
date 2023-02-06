package search

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sion96994/go/openapi/client"
)

// DoSearch 搜索
func DoSearch(c *client.Client, arg *SearchReqParams) (*SearchResult, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Set("adzone_id", arg.AdzoneId)
	params.Set("page_size", fmt.Sprintf("%d", arg.PageSize))
	params.Set("page_no", fmt.Sprintf("%d", arg.PageNo))
	params.Set("q", arg.Q)
	params.Set("has_coupon", fmt.Sprintf("%v", arg.HasCoupon))
	params.Set("method", "taobao.tbk.dg.material.optional")
	if len(arg.Sort) > 0 {
		params.Set("sort", arg.Sort)
	}

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
	resp, err := data.Get("tbk_dg_material_optional_response").MarshalJSON()
	if err != nil {
		return nil, err
	}

	// 5. 解析响应
	var (
		result *SearchResult
	)
	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSearchSuggestWords 获取搜索联想词
func GetSearchSuggestWords(keyWords string) ([]string, error) {
	req := &SuggestReq{
		Q:    keyWords,
		Area: "etao",
		Code: "utf-8",
	}
	respBs, err := DoHttpRequest("https://suggest.taobao.com/sug", req)
	if err != nil {
		return nil, err
	}

	var (
		words []string
	)
	result := &SuggestResp{}
	err = json.Unmarshal(respBs, result)
	if err != nil {
		return nil, err
	}
	for _, r := range result.Result {
		if len(r) > 0 {
			words = append(words, r[0])
		}
	}
	return words, nil
}
