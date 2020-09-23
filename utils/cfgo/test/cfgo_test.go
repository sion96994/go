package main

import (
	"fmt"
	"testing"

	"github.com/sion96994/go/utils/cfgo"
	"gitlab.xiaoenai.net/xserver/utils/cfgo/test/m1"
	_ "gitlab.xiaoenai.net/xserver/utils/cfgo/test/m2"
)

type T struct {
	C string
	m1.T1
}

func (t *T) Reload(bind cfgo.BindFunc) error {
	fmt.Println("main T reload do some thing...")
	return bind()
}

func Test1(t *testing.T) {
	structPtr := &T{
		C: "c",
		T1: m1.T1{
			B: 2, //default value
		},
	}

	// output: config/config.yaml
	c := cfgo.MustGet("config/config.yaml")
	c.MustReg("section", structPtr)
	// or
	// cfgo.MustReg("section", structPtr)

	fmt.Printf("structPtr(config/config.yaml): %+v\n\n", structPtr)

	// output: config/config3.yaml
	c3 := cfgo.MustGet("config/config3.yaml", true)
	c3.MustReg("section", structPtr)
	fmt.Printf("structPtr(config/config3.yaml): %+v\n\n", structPtr)

	fmt.Printf(" ----------------------------------------------------------- \n\n")

	fmt.Printf("config/config.yaml content:\n%s\n\n", c.Content())
	// or
	// fmt.Printf("config.yaml content:\n%s\n\n", cfgo.Content())

	fmt.Printf(" ----------------------------------------------------------- \n\n")

	fmt.Printf("config/config3.yaml content:\n%s\n\n", c3.Content())
}

type M struct {
	Auto bool
}

func (m *M) Reload(bind cfgo.BindFunc) error {
	return bind()
}

func Test2(t *testing.T) {
	m := new(M)
	mixed := cfgo.MustGet("config/mixed_config.yaml")
	mixed.MustReg("register", m)

	fmt.Printf("config/mixed_config.yaml content:\n%s\n\n", mixed.Content())
	fmt.Printf("config/mixed_config.yaml config m:\n%#v\n\n", m)
	{
		custom, _ := mixed.GetSection("custom")
		fmt.Printf("config/mixed_config.yaml GetSection 'custom':\n%#v\n\n", custom)
	}
	{
		var custom bool
		_ = mixed.BindSection("custom", &custom)
		fmt.Printf("config/mixed_config.yaml BindSection 'custom':\n%#v\n\n", custom)
	}
}