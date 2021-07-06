package main

import (
	api "github.com/aambayec/tut-grpc-go-web/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	db, err := xorm.NewEngine("mysql", "root:Ulyanin123@/tut_grpc")
	if err != nil {
		panic(err)
	}

	_, err = db.DBMetas()
	if err != nil {
		panic(err)
	}

	api.Run(50051, db)
}