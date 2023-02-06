package search

// SearchReq
type SearchReqParams struct {
	// 必须参数
	// 推广位ID adzone_id
	AdzoneId string `json:"adzone_id"`
	// 页面大小
	PageSize int32 `json:"page_size"`
	// 第几页
	PageNo int64 `json:"page_no"`
	// 查询词
	Q string `json:"q"`

	// 可选参数
	// 链接形式[1:PC 2:无线 默认为1]
	Platform int32 `json:"platform"`
	// 折扣价范围上限
	EndPrice int32 `json:"end_price"`
	// 折扣价范围下限
	StartPrice int32 `json:"start_price"`
	// 是否为天猫商品
	IsTmall bool `json:"is_tmall"`
	// 排序(_des（降序），排序_asc（升序），销量（total_sales），价格（price）)
	Sort string `json:"sort"`
	// 物料id（默认：2836，算法优化：17004）
	MaterialId int32 `json:"material_id"`
	// 是否有优惠券
	HasCoupon bool `json:"has_coupon"`
}

// 搜索结果
type SearchResult struct {
	// 搜索结果数量
	TotalResults int64 `json:"total_results"`
	// 商品列表
	ResultList *ResultList `json:"result_list"`
	// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
	PageResultKey string `json:"page_result_key"`
}

// 搜索错误结构体
type SearchError struct {
	Code    int32  `json:"code"`
	SubMsg  string `json:"sub_msg"`
	SubCode string `json:"sub_code"`
	Msg     string `json:"msg"`
}

type ResultList struct {
	MapData []*Products `json:"map_data"`
}

// Products
type Products struct {
	// 优惠券开始时间 日期
	CouponStartTime string `json:"coupon_start_time"`
	// 优惠券结束时间 日期
	CouponEndTime string `json:"coupon_end_time"`
	// 优惠券ID
	CouponId string `json:"coupon_id"`
	// 商品ID
	ItemId int64 `json:"item_id"`
	// 商品标题
	Title string `json:"title"`
	// 商品短标题
	ShortTitle string `json:"short_title"`
	// 商品主图
	PictUrl string `json:"pict_url"`
	// 商品小图列表
	SmallImages *SmallImages `json:"small_images"`
	// 商品一口价
	ReservePrice string `json:"reserve_price"`
	// 折扣价
	ZkFinalPrice string `json:"zk_final_price"`
	// 卖家类型（0：集市场 1：天猫）
	UserType int32 `json:"user_type"`
	// 商品所在地
	Provcity string `json:"provcity"`
	// 商品链接
	ItemUrl string `json:"item_url"`
	// 30天销量
	Volume int64 `json:"volume"`
	// 优惠券满减信息
	CouponInfo string `json:"coupon_info"`
	// 优惠券总量
	CouponTotalCount int64 `json:"coupon_total_count"`
	// 优惠券剩余量
	CouponRemainCount int64 `json:"coupon_remain_count"`
	// 优惠券起用门槛
	CouponStartFee string `json:"coupon_start_fee"`
	// 优惠券价格
	CouponAmount string `json:"coupon_amount"`
	// 商品+券二合一页面链接
	CouponShareUrl string `json:"coupon_share_url"`
}

type SmallImages struct {
	String []string `json:"string"`
}

// SuggestReq 联想词请求参数
type SuggestReq struct {
	Q    string `json:"q"`
	Area string `json:"area"`
	Code string `json:"code"`
}

// SuggestResp 联想词响应参数
type SuggestResp struct {
	Result [][]string `json:"result"`
}
