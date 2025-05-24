package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hackathon_be/internal/config"
	"hackathon_be/internal/handler"
	"hackathon_be/internal/model"

	"hackathon_be/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

func InitDB(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	//初始化所有表
	err = InitTables(db)
	if err != nil {
		panic(err)
	}

}
func InitTables(db *gorm.DB) error {
	return db.Table("score").AutoMigrate(&model.Score{})
}

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	InitDB(c.DSN())
	DB := sqlx.NewMysql(c.DSN())
	ctx := svc.NewServiceContext(c, DB)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
