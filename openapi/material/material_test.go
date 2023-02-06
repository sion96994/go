package material

import (
	"testing"

	tp "github.com/henrylee2cn/teleport"
	client2 "github.com/sion96994/go/openapi/client"
)

func TestMaterial(t *testing.T) {
	// client
	client := client2.GetClient("32527641", "994507ab1a8c42fd3a203fd4ce8a97ae", "")
	materialList, err := DoGetMaterialItems(client, &MaterialParams{
		PageSize:   1,
		PageNo:     1,
		AdzoneId:   "111039700183",
		MaterialId: 27446,
	})
	if err != nil {
		tp.Errorf("TestMaterial: err-> %v", err)
	}
	if len(materialList) > 0 {
		tp.Infof("TestMaterial: title-> %s", materialList[0].Title)
	}
}
