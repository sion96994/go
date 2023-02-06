package coupon

// CouponGetParams
type CouponGetParams struct {
	// 商品id
	ItemId string
	// 优惠券id
	ActivityId string
	// me参数
	E string
}

// CouponGetResp
type CouponGetResp struct {
	CouponStartFee    string `json:"coupon_start_fee"`
	CouponRemainCount int    `json:"coupon_remain_count"`
	CouponTotalCount  int    `json:"coupon_total_count"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponStartTime   string `json:"coupon_start_time"`
	// 优惠券金额
	CouponAmount      string `json:"coupon_amount"`
	CouponSrcScene    int    `json:"coupon_src_scene"`
	CouponType        int    `json:"coupon_type"`
	CouponActivityID  string `json:"coupon_activity_id"`
}