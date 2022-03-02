package main

import (
	"github.com/sion96994/go/tp-template/account/api"
	"github.com/sion96994/go/tp-template/plugin"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"
)

func main() {
	srv := micro.NewServer(
		cfg.Srv,
		discovery.ServicePlugin(cfg.Srv.InnerIpPort(), cfg.Etcd),
		plugin.NewInnerAuth(true),
	)
	api.Route("/account", srv.Router())
	srv.ListenAndServe()
}
