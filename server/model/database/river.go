package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO 数据集功能

// River 数据集 Address 作为备注
type River struct {
	field.DefaultField `bson:",inline"`
	Name               string               `json:"name" bson:"name"`
	Address            string               `json:"address" bson:"address"`
	Algae              []primitive.ObjectID `json:"algae" bson:"algae"`
}

// QueryRiverListByEmail 获取用户拥有数据集
func (m *Mgo) QueryRiverListByEmail(email string) ([]primitive.ObjectID, error) {
	var one Operator
	err := operator.Find(ctx, bson.M{"email": email}).Select(bson.D{{"dataSetID", 1}}).One(&one)
	return one.DataSetID, err
}

// QueryRiverByIdWithoutAlgae 待测试数据没有algae是否可以被绑定
func (m *Mgo) QueryRiverByIdWithoutAlgae(obj primitive.ObjectID) *River {
	var one River
	if err := river.Find(ctx, bson.M{"_id": obj}).Select(bson.M{"algae": 0}).One(&one); err != nil {
		return nil
	}
	return &one
}
func (m *Mgo) QueryRiverById(obj primitive.ObjectID) *River {
	var one River
	if err := river.Find(ctx, bson.M{"_id": obj}).One(&one); err != nil {
		return nil
	}
	return &one
}

func (m *Mgo) QueryRiverByName(obj string) *River {
	var one River
	if err := river.Find(ctx, bson.M{"name": obj}).One(&one); err != nil {
		return &River{}
	}
	return &one
}

func (m *Mgo) ExistsRiver(name string) bool {
	var one River
	if err := river.Find(ctx, bson.M{"name": name}).One(&one); err != nil {
		return false
	}
	return true
}

func (m *Mgo) UpdateRiver(id primitive.ObjectID, r River) error {
	return river.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": r.Name, "address": r.Address}})
}
func (m *Mgo) UpdateRiverAlgae(id primitive.ObjectID, algae []primitive.ObjectID) error {
	return river.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"algae": algae}})
}
func (m *Mgo) InsertRiver(r *River) (interface{}, error) {
	res, err := river.InsertOne(ctx, r)
	return res.InsertedID, err
}

func (m *Mgo) DropRiver(name string) error {
	err := river.Remove(ctx, bson.M{"name": name})
	return err
}
