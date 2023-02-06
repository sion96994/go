package material

// 接口文档 https://open.taobao.com/api.htm?docId=33947&docType=2
// MaterialParams 精选商品参数
type MaterialParams struct {
	// 每页数量
	PageSize int
	// 页数
	PageNo int
	// 推广id
	AdzoneId string
	// 物料ID
	MaterialId int32
}

// 物料信息
type MaterInfo struct {
	CategoryID           int          `json:"category_id"`
	ClickURL             string       `json:"click_url"`
	CommissionRate       string       `json:"commission_rate"`
	CouponAmount         int          `json:"coupon_amount"`
	CouponClickURL       string       `json:"coupon_click_url"`
	CouponEndTime        string       `json:"coupon_end_time"`
	CouponRemainCount    int          `json:"coupon_remain_count"`
	CouponShareURL       string       `json:"coupon_share_url"`
	CouponStartFee       string       `json:"coupon_start_fee"`
	CouponStartTime      string       `json:"coupon_start_time"`
	CouponTotalCount     int          `json:"coupon_total_count"`
	ItemDescription      string       `json:"item_description"`
	ItemID               int64        `json:"item_id"`
	JhsPriceUspList      string       `json:"jhs_price_usp_list"`
	LevelOneCategoryID   int          `json:"level_one_category_id"`
	LevelOneCategoryName string       `json:"level_one_category_name"`
	Nick                 string       `json:"nick"`
	PictURL              string       `json:"pict_url"`
	ReservePrice         string       `json:"reserve_price"`
	SellerID             int          `json:"seller_id"`
	ShopTitle            string       `json:"shop_title"`
	ShortTitle           string       `json:"short_title"`
	SmallImages          *SmallImages `json:"small_images"`
	SubTitle             string       `json:"sub_title"`
	Title                string       `json:"title"`
	UserType             int          `json:"user_type"`
	Volume               int          `json:"volume"`
	WhiteImage           string       `json:"white_image"`
	ZkFinalPrice         string       `json:"zk_final_price"`
}

// SmallImages 小图
type SmallImages struct {
	String []string `json:"string"`
}
