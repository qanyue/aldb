package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO 图片直接拥有标注

type Alga struct {
	field.DefaultField `bson:",inline"`
	Name               string       `json:"name" bson:"name"`
	Src                string       `json:"src" bson:"src"`
	Annotations        []Annotation `json:"annotations" bson:"annotations"`
}

func (m *Mgo) QueryAlgae() ([]*Alga, error) {
	var batch []*Alga
	err := algae.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

// QueryAlgaByName 通过名称查找
func (m *Mgo) QueryAlgaByName(name string) (*Alga, error) {
	var one *Alga
	err := algae.Find(ctx, bson.M{"name": name}).One(&one)
	return one, err
}

// QueryAlgaByKey 正则查找
func (m *Mgo) QueryAlgaByKey(key string) ([]*Alga, error) {
	var batch []*Alga
	err := algae.Find(ctx, bson.M{"name": bson.M{"$regex": key}}).All(&batch)
	return batch, err
}

// InsertAlga InsertedID 作用待更新
func (m *Mgo) InsertAlga(obj *Alga) (interface{}, error) {
	res, err := algae.InsertOne(ctx, obj)
	return res.InsertedID, err
}

// UpdateAlga model层引用作用待更新
func (m *Mgo) UpdateAlga(id primitive.ObjectID, annotations []Annotation) error {
	return algae.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"annotations": annotations}})
}

func (m *Mgo) DropAlga(name string) error {
	err := algae.Remove(ctx, bson.M{"name": name})
	return err
}
