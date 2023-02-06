package parse

import (
	"testing"

	tp "github.com/henrylee2cn/teleport"
)

func TestParseCouponUrl(t *testing.T) {
	result, err := ParseCouponUrl("YpcLQAdjFr", "https://uland.taobao.com/coupon/edetail?e=69upSlgW3U8NfLV8niU3R5TgU2jJNKOfNNtsjZw%2F%2FoIynMBOc6G5dw0KetbhezUkQs4S8hLo%2FzClvjmGEHklkLhtWWjRAlP0awvDoeqYhH3RR%2BU2m1liADqKKwc6jcWIu0H0jzxnhmm6Zhb%2BGhHK%2BL%2Bsbwi7vobRwSx%2FAhAJj9NYliOgPhsgBkRtnm4cMr8nNJY9ThRf5vklM1ZJHcLCJg%3D%3D&&app_pvid=59590_33.5.221.244_744_1616118521777&ptl=floorId:27453;app_pvid:59590_33.5.221.244_744_1616118521777;tpp_pvid:b9ae7088-d531-488e-a654-8028380e7016&union_lens=lensId%3AMAPI%401616118521%40b9ae7088-d531-488e-a654-8028380e7016_609597271306%401")
	if err != nil {
		tp.Errorf("err-> %v", err)
	}
	tp.Infof("ParseCouponUrl: r-> %v", result)
}
