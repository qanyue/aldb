package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type River struct {
	field.DefaultField `bson:",inline"`
	Name               string               `json:"name" bson:"name"`
	Address            string               `json:"address" bson:"address"`
	Algae              []primitive.ObjectID `json:"algae" bson:"algae"`
}

func (m *Mgo) QueryRivers() ([]*River, error) {
	var batch []*River
	err := river.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

func (m *Mgo) QueryRiverById(obj primitive.ObjectID) River {
	var one River
	if err := river.Find(ctx, bson.M{"_id": obj}).One(&one); err != nil {
		return River{}
	}
	return one
}

func (m *Mgo) QueryRiverByName(obj string) River {
	var one River
	if err := river.Find(ctx, bson.M{"name": obj}).One(&one); err != nil {
		return River{}
	}
	return one
}

func (m *Mgo) ExistsRiver(name string) bool {
	var one River
	if err := river.Find(ctx, bson.M{"name": name}).One(&one); err != nil {
		return false
	}
	return true
}

func (m *Mgo) UpdateRiver(id primitive.ObjectID, algae []primitive.ObjectID) error {
	return river.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"algae": algae}})
}

func (m *Mgo) InsertRiver(r *River) error {
	_, err := river.InsertOne(ctx, r)
	return err
}

func (m *Mgo) DropRiver(name string) error {
	err := river.Remove(ctx, bson.M{"name": name})
	return err
}
