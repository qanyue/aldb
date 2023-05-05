package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// TODO 图片直接拥有标注
type Alga struct {
	field.DefaultField `bson:",inline"`
	Name               string       `json:"name" bson:"name"`
	Src                string       `json:"src" bson:"src"`
	Annotations        []Annotation `json:"annotations" bson:"annotations"`
}

// QueryAlgae 数据集直接内嵌algae,所以通过数据集id获取图片
func (m *Mgo) QueryAlgaeById(id primitive.ObjectID) *Alga {
	var one Alga
	err := algae.Find(ctx, bson.M{"_id": id}).One(&one)
	if err != nil {
		zap.L().Error("获取图片发生错误", zap.String("QueryAlgae", err.Error()))
		return nil
	}
	return &one
}

// QueryAlgaByName 在数据集中通过名称查找，
func (m *Mgo) QueryAlgaByName(name string) (*Alga, error) {
	var one Alga
	err := algae.Find(ctx, bson.D{{"name", name}}).One(&one)
	return &one, err
}

// QueryAlgaByKey 正则查找
// TODO 可能可以使用前端实现
func (m *Mgo) QueryAlgaByKey(key string) ([]Alga, error) {
	var batch []Alga
	err := algae.Find(ctx, bson.D{{"name", bson.D{{"$regex", key}}}}).All(&batch)
	return batch, err
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
