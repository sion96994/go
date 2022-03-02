package main

import (
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"

	"github.com/sion96994/go/tp-template/user/api"
)

func main() {
	srv := micro.NewServer(
		cfg.Srv,
		discovery.ServicePlugin(cfg.Srv.InnerIpPort(), cfg.Etcd),
	)
	api.Route("/user", srv.Router())
	srv.ListenAndServe()
}
