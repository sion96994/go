package item

import (
	"testing"

	tp "github.com/henrylee2cn/teleport"
	client2 "github.com/sion96994/go/openapi/client"
)

func TestItem(t *testing.T) {
	// client
	client := client2.GetClient("32527641", "994507ab1a8c42fd3a203fd4ce8a97ae", "")
	itemInfo, err := GetItemInfo(client, []string{"628493994111L"})
	if err != nil {
		tp.Errorf("err-> %v", err)
	}
	t.Logf("%d", len(itemInfo))
}
