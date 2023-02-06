package item

// ItemInfo taobao item info
type ItemInfo struct {
	ItemID      int64  `json:"num_iid"`
	Title       string `json:"title"`
	PictURL     string `json:"pict_url"`
	SmallImages struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ReservePrice string `json:"reserve_price"`
	ZkFinalPrice string `json:"zk_final_price"`
	UserType     int    `json:"user_type"`
	Provcity     string `json:"provcity"`
	ItemURL      string `json:"item_url"`
	SellerName   string `json:"nick"`
	SellerID     int64  `json:"seller_id"`
	Volume       int32  `json:"volume"`
	CatName      string `json:"cat_name"`
	CatLeafName  string `json:"cat_leaf_name"`
}
