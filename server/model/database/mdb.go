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

	algae    *qmgo.Collection
	river    *qmgo.Collection
	operator *qmgo.Collection
	//annotation *qmgo.Collection
	tag *qmgo.Collection
)

func init() {
	var err error
	ctx = context.Background()
	cli, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: cfg.Uri})
	if err != nil {
		panic("Fail to connect mongodb")
	}
	mdb = cli.Database(cfg.DB)
	river = mdb.Collection("river")
	algae = mdb.Collection("algae")
	operator = mdb.Collection("operator")
	tag = mdb.Collection("tag")
}

type Mgo struct{}

func Close() {
	if err := cli.Close(ctx); err != nil {
		fmt.Println(err)
	}
}

func CreateIndex() {
	err := river.CreateIndexes(ctx, []options.IndexModel{
		{Key: []string{"_id"}}, //river _id索引
		{Key: []string{"algae"}},
	})
	if err != nil {
		panic("river 索引创建失败")
	}
	err = algae.CreateIndexes(ctx, []options.IndexModel{
		{Key: []string{"_id"}},
	})
	if err != nil {
		panic("algae 索引创建失败")
	}
	err = tag.CreateOneIndex(ctx, options.IndexModel{Key: []string{"name"}})
	if err != nil {
		panic("tag 索引创建失败")
	}

}
