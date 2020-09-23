package main

import (
	"os"
	"testing"

	log "github.com/sion96994/go/logger"
	"github.com/sion96994/go/utils/cfgo"
)

func Test(t *testing.T) {
	t.Logf("%#v\n", log.DefaultLogger())
	v, ok := cfgo.GetSection("log")
	t.Logf("GetSection log: %v, ok: %v\n", v, ok)
	v, ok = cfgo.GetSection("0")
	t.Logf("GetSection 0: %v, ok: %v\n", v, ok)

	l := new(log.LogConfig)
	err := cfgo.BindSection("log", l)
	t.Logf("BindSection log: %v, err: %v\n", l, err)

	var a bool
	err = cfgo.BindSection("0", &a)
	t.Logf("BindSection 0: %v, err: %v\n", a, err)
}

func ExampleLog() {
	//log.Printf("print")
	//log.Tracef("trace")
	//log.Debugf("debug")
	//log.Infof("info")
	//log.Errorf("error")
	//log.Fatalf("fatal")

	log.DefaultLogger().Level = log.INFO

	log.Tracef("trace")
	log.Debugf("debug")
	log.Infof("info")
	log.Errorf("error")
	log.Fatalf("fatal")

	log.DefaultLogger().Colorful = false

	log.Tracef("trace")
	log.Debugf("debug")
	log.Infof("info")
	log.Errorf("error")
	log.Fatalf("fatal")

	log.NewNamedLogger("named", os.Stdout)
	nl := log.GetLogger("named")
	nl.Tracef("trace, %d", 1)
	nl.Debugf("debug, %d", 2)
	nl.Infof("info, %d", 3)
	nl.Errorf("error, %d", 4)
	nl.Fatalf("fatal, %d", 5)

	sl := nl.SetPrefix("STAT")
	sl.Tracef("trace, %d", 1)
	sl.Debugf("debug, %d", 2)
	sl.Infof("info, %d", 3)
	sl.Errorf("error, %d", 4)
	sl.Fatalf("fatal, %d", 5)

	// Output:
	//
}
