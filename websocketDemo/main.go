package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sion96994/go/logger"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Client struct {
	Addr string          `json:"addr"`
	Conn *websocket.Conn `json:"conn"`
}

type Msg struct {
	Seq     string `json:"seq"`
	Content string `json:"content"`
}

type Server struct {
	Login chan *Client
}

var (
	server *Server
)

func Init() {
	server = &Server{
		Login: make(chan *Client, 1000),
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))
	//收到http请求(upgrade),完成websocket协议转换
	//在应答的header中放上upgrade:websoket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//报错了，直接返回底层的websocket链接就会终断掉
		log.Errorf("Upgrade err -> %v", err)
		return
	}
	client := &Client{
		Addr: ws.RemoteAddr().String(),
		Conn: ws,
	}
	server.Login <- client
	log.Debugf("client -> %v", client.Addr)

	//得到了websocket.Conn长连接的对象，实现数据的收发
	for {
		//Text(json), Binary
		_, data, err := ws.ReadMessage()
		if err != nil {
			//报错关闭websocket
			log.Errorf("err -> %v", err)
			goto ERR
		}
		log.Infof("data -> %v", string(data))
		//发送数据，判断返回值是否报错
		//if err = ws.WriteMessage(websocket.TextMessage, data); err != nil {
		//	//报错了
		//	goto ERR
		//}
		if err := ws.WriteJSON(Msg{
			Seq:     "1",
			Content: "hello",
		}); err != nil {
			log.Errorf("err -> %v", err)
			goto ERR
		}
	}
	//error的标签
ERR:
	ws.Close()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func worker() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		log.Warnf("chan for len -> %v", len(server.Login))
		select {
		case coon := <-server.Login:
			i := 2
			log.Infof("用户连接 addr -> %v", coon.Addr)
			go func() {
				for {
					time.Sleep(time.Duration(5) * time.Second)
					coon.Conn.WriteJSON(Msg{
						Seq:     fmt.Sprintf("%d", i),
						Content: "test",
					})
					i += 1
					log.Debugf("addr -> %v, worker -> %d", coon.Addr, i)
				}
			}()
		default:
			log.Warnf("chan select len -> %v", len(server.Login))
			break
		}
	}
}

func main() {
	//http://localhost:7777/ws
	Init()

	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/v1", helloHandler)
	go worker()
	//服务端启动
	log.Infof("服务端启动")
	http.ListenAndServe("0.0.0.0:7777", nil)
}
