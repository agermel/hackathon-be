package main

import (
	"flag"
	"fmt"

	"hackathon_be/internal/config"
	"hackathon_be/internal/handler"
	"hackathon_be/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ha-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	DB := sqlx.NewMysql(c.DSN())
	ctx := svc.NewServiceContext(c, DB)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
