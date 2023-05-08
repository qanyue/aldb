package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Alga 图片直接拥有标注
type Alga struct {
	field.DefaultField `bson:",inline"`
	Name               string       `json:"name" bson:"name"`
	Src                string       `json:"src" bson:"src"`
	Annotations        []Annotation `json:"annotations" bson:"annotations"`
}

// QueryAlgaById 数据集直接内嵌algae,所以通过数据集id获取图片
func (m *Mgo) QueryAlgaById(id primitive.ObjectID) *Alga {
	var one Alga
	err := algae.Find(ctx, bson.M{"_id": id}).One(&one)
	if err != nil {
		zap.L().Error("获取图片发生错误", zap.String("QueryAlgae", err.Error()))
		return nil
	}
	return &one
}

// QueryAlgaByName 在数据集中通过名称查找，返回alga切片(名字会重复)
func (m *Mgo) QueryAlgaByName(id primitive.ObjectID, name string) []Alga {
	var batch []Alga
	a := m.QueryRiverById(id)
	if a == nil {
		return []Alga{}
	}
	err := algae.Find(ctx, bson.D{{"name", name}, {"_id", bson.M{"$in": a.Algae}}}).All(&batch)
	if err != nil {
		return []Alga{}
	}
	return batch
}

// QueryAlgaByKey 正则查找
func (m *Mgo) QueryAlgaByKey(id primitive.ObjectID, key string) []Alga {
	var batch []Alga
	a := m.QueryRiverById(id)
	if a == nil {
		return []Alga{}
	}
	err := algae.Find(ctx, bson.D{
		{"name", bson.D{
			{"$regex", key},
		}},
		{"_id", bson.M{"$in": a.Algae}},
	}).All(&batch)
	if err != nil {
		return []Alga{}
	}
	return batch
}

// InsertAlga 在数据集中插入一个alga
func (m *Mgo) InsertAlga(obj *Alga) (interface{}, error) {
	res, err := algae.InsertOne(ctx, obj)
	return res.InsertedID, err
}

// UpdateAlgaAnno TODO model层引用作用待更新
//
//	func (m *Mgo) UpdateAlgaAnno(aId primitive.ObjectID, annotation Annotation) error {
//		return algae.UpdateOne(ctx, bson.D{{"_id", aId}},
//			bson.M{"$addToSet": bson.M{"annotation": annotation}})
//	}

func (m *Mgo) DropAlga(id primitive.ObjectID) error {
	err := algae.Remove(ctx, bson.D{{"_id", id}})
	return err
}
