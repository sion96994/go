package search

import (
	"testing"

	tp "github.com/henrylee2cn/teleport"
	client2 "github.com/sion96994/go/openapi/client"
)

func TestItem(t *testing.T) {
	// client
	client := client2.GetClient("32527641", "994507ab1a8c42fd3a203fd4ce8a97ae", "")
	res, err := DoSearch(client, &SearchReqParams{
		AdzoneId:   "111039700183",
		PageSize:   1,
		PageNo:     0,
		Q:          "鞋子",
		Platform:   2,
		MaterialId: 17004,
		HasCoupon:  true,
	})
	if err != nil {
		tp.Errorf("err-> %v", err)
	}
	t.Logf("%v", res.PageResultKey)

	//// client
	//res, err := GetSearchSuggestWords("birth")
	//if err != nil {
	//	tp.Errorf("err-> %v", err)
	//}
	//t.Logf("%v", res)
}
