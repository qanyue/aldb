package model

import (
	"errors"
	"github.com/qanyue/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RiverBindToUser 根据用户邮箱绑定数据集对应用户
func RiverBindToUser(userEmail string, id primitive.ObjectID) error {
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return err
	}
	var sets []primitive.ObjectID
	if user.DataSetID == nil {
		sets = []primitive.ObjectID{id}
	} else {
		sets = append(user.DataSetID, id)
	}
	err = mgo.UpdateOperator(user.Id, sets)
	return err
}

func AddRiver(userEmail string, obj River) (interface{}, error) {
	if mgo.ExistsRiver(obj.Name) {
		return nil, errors.New("river exists")
	}
	id, err := mgo.InsertRiver(&database.River{
		Name:    obj.Name,
		Address: obj.Address,
		Algae:   []primitive.ObjectID{},
	})
	v, ok := id.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("数据集id到")
	}
	err = RiverBindToUser(userEmail, v)
	return id, err
}

func GetRivers(userEmail string) []River {
	rId, err := mgo.QueryRiverListByEmail(userEmail)
	if err != nil {
		return nil
	}
	res := make([]River, 0)
	for _, obj := range rId {
		r := mgo.QueryRiverById(obj)

		res = append(res, River{
			Name:    r.Name,
			Address: r.Address,
			Algae:   objectIdToString(r.Algae),
		})
	}
	return res
}
