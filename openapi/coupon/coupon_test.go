package coupon

import (
	"testing"

	tp "github.com/henrylee2cn/teleport"
	client2 "github.com/sion96994/go/openapi/client"
)

func TestCouponGet(t *testing.T){
	client := client2.GetClient("32527641", "994507ab1a8c42fd3a203fd4ce8a97ae", "")
	coupon, err := GetCouponInfo(client, &CouponGetParams{
		ItemId: "609597271306",
		ActivityId: "fe1ef6223854433fac3b17d86d04e96e",
	})
	if err != nil {
		tp.Errorf("err-> %v",err)
	}
	tp.Infof("coupon-> %v", coupon)
}