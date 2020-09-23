// +build !windows

package cfgo

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	logchan   = make(chan *Msg, 100)
	getOnce   sync.Once
	hasLogger bool
)

func GetLogchan() (c <-chan *Msg, ok bool) {
	getOnce.Do(func() {
		c = logchan
		ok = true
		hasLogger = true
	})
	return
}

type Msg struct {
	Ok  bool
	Txt string
}

func init() {
	go func() {
		chSignal := make(chan os.Signal)
		defer signal.Stop(chSignal)
		signal.Notify(chSignal, syscall.SIGUSR1)
		for {
			<-chSignal
			err := ReloadAll()
			var msg = new(Msg)
			if err != nil {
				msg.Ok = false
				msg.Txt = "reload config: " + err.Error()
			} else {
				msg.Ok = true
				msg.Txt = "reload config ok"
			}
			if hasLogger {
				logchan <- msg
			} else {
				log.Println(msg.Txt)
			}
		}
	}()
}
