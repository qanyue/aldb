package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operator struct {
	field.DefaultField `bson:",inline"`
	Name               string               `json:"name" bson:"name"`
	Password           string               `json:"password" bson:"password"`
	Email              string               `json:"email" bson:"email"`
	Access             int                  `json:"access" bson:"access"`
	Annotations        []primitive.ObjectID `json:"annotations" bson:"annotations"`
}

func (m *Mgo) QueryOperatorByEmail(email string) (*Operator, error) {
	var one *Operator
	err := operator.Find(ctx, bson.M{"email": email}).One(&one)
	return one, err
}

func (m *Mgo) QueryOperators() ([]*Operator, error) {
	var batch []*Operator
	err := operator.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

func (m *Mgo) CheckOperator(email, password string) (*Operator, error) {
	var one *Operator
	err := operator.Find(ctx, bson.M{"email": email, "password": password}).One(&one)
	return one, err
}

func (m *Mgo) InsertOperator(obj *Operator) error {
	_, err := operator.InsertOne(ctx, obj)
	return err
}

func (m *Mgo) UpsertOperator(email, name, password string, access int) error {
	return operator.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"name": name, "password": password, "access": access}})
}

func (m *Mgo) ExistsOperator(email string) bool {
	var one Operator
	if err := operator.Find(ctx, bson.M{"email": email}).One(&one); err != nil {
		return false
	}
	return true
}

func (m *Mgo) UpdateOperator(id primitive.ObjectID, annotations []primitive.ObjectID) error {
	return operator.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"annotations": annotations}})
}

func (m *Mgo) UpdatePassword(email string, newPassword string) error {
	return operator.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": newPassword}})
}

func (m *Mgo) DropOperator(email string) error {
	return operator.Remove(ctx, bson.M{"email": email})
}
