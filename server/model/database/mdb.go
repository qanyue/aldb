package database

import (
	"context"
	"fmt"
	"github.com/qanyue/aldb/server/config"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
)

var (
	cli *qmgo.Client
	mdb *qmgo.Database
	ctx context.Context
	cfg = config.Conf.Mongo

	algae      *qmgo.Collection
	river      *qmgo.Collection
	operator   *qmgo.Collection
	annotation *qmgo.Collection
	tag        *qmgo.Collection
)

func init() {
	var err error
	ctx = context.Background()
	cli, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: cfg.Uri})
	if err != nil {
		panic("Fail to connect mongodb")
	}
	mdb = cli.Database(cfg.DB)
	algae = mdb.Collection("algae")
	river = mdb.Collection("river")
	operator = mdb.Collection("operator")
	annotation = mdb.Collection("annotation")
	tag = mdb.Collection("tag")
}

type Mgo struct{}

func Close() {
	if err := cli.Close(ctx); err != nil {
		fmt.Println(err)
	}
}

func CreateIndex() {
	if err := algae.CreateOneIndex(ctx, options.IndexModel{
		Key: []string{"name"},
	}); err != nil {
		panic("Fail to create index")
	}
}
