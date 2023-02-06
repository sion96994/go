package item

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/sion96994/go/openapi/client"
)

// GetItemInfo 获取商品信息
func GetItemInfo(c *client.Client, items []string) ([]*ItemInfo, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Set("fields", "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,volume,seller_id,nick,cat_name,cat_leaf_name")
	params.Set("num_iids", strings.Join(items, ","))
	params.Set("method", "taobao.tbk.item.info.get")

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
	resp, err := data.Get("tbk_item_info_get_response").GetPath("results", "n_tbk_item").MarshalJSON()
	if err != nil {
		return nil, err
	}

	// 5. 解析响应
	var (
		itemList []*ItemInfo
	)
	if err = json.Unmarshal(resp, &itemList); err != nil {
		return nil, err
	}
	return itemList, nil
}
