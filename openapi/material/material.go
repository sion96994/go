package material

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sion96994/go/openapi/client"
)

// DoGetMaterialItems 获取精选商品
func DoGetMaterialItems(c *client.Client, arg *MaterialParams) ([]*MaterInfo, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Set("adzone_id", arg.AdzoneId)
	params.Set("page_size", fmt.Sprintf("%d", arg.PageSize))
	params.Set("page_no", fmt.Sprintf("%d", arg.PageNo))
	params.Set("material_id", fmt.Sprintf("%d", arg.MaterialId))
	params.Set("method", "taobao.tbk.dg.optimus.material")

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
	resp, err := data.Get("tbk_dg_optimus_material_response").GetPath("result_list", "map_data").MarshalJSON()
	if err != nil {
		return nil, err
	}

	// 5. 解析响应
	var (
		materialList []*MaterInfo
	)
	if err = json.Unmarshal(resp, &materialList); err != nil {
		return nil, err
	}
	return materialList, nil
}
