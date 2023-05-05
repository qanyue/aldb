package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
)

type Tag struct {
	field.DefaultField `bson:",inline"`
	Name               string `json:"name" bson:"name"`
	ResourceName       string `json:"resourceName" bson:"resourceName"`
}

func (m *Mgo) QueryTags() ([]*Tag, error) {
	var batch []*Tag
	err := tag.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

func (m *Mgo) QueryTagByName(obj string) Tag {
	var one Tag
	if err := tag.Find(ctx, bson.M{"name": obj}).One(&one); err != nil {
		return Tag{}
	}
	return one
}

func (m *Mgo) ExistsTag(name string) bool {
	var one Tag
	if err := tag.Find(ctx, bson.M{"name": name}).One(&one); err != nil {
		return false
	}
	return true
}

func (m *Mgo) InsertTag(t *Tag) error {
	_, err := tag.InsertOne(ctx, t)
	return err
}

func (m *Mgo) DropTag(name string) error {
	err := tag.Remove(ctx, bson.M{"name": name})
	return err
}
