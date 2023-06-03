package model

import (
	"errors"
	"regexp"

	"github.com/qanyue/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRiverByID(id string) *River {
	rId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	r := mgo.QueryRiverById(rId)
	return &River{
		Name:    r.Name,
		Address: r.Address,
		Algae:   objectIdToString(r.Algae),
	}
}

// RiverBindToUser 根据用户邮箱绑定数据集对应用户
func RiverBindToUser(userEmail string, id primitive.ObjectID) error {
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return err
	}
	s := user.DataSetID
	if len(s) <= 0 {
		s = []primitive.ObjectID{id}
	} else {
		s = append(user.DataSetID, id)
	}
	err = mgo.UpdateOperator(user.Id, s)
	return err
}

func AddRiver(obj River) (primitive.ObjectID, error) {
	if mgo.ExistsRiver(obj.Name) {
		return primitive.NewObjectID(), errors.New("river exists")
	}
	id, err := mgo.InsertRiver(&database.River{
		Name:    obj.Name,
		Address: obj.Address,
		Algae:   []primitive.ObjectID{},
	})
	v, ok := id.(primitive.ObjectID)
	if !ok {
		return primitive.NewObjectID(), errors.New("id 转换失败")
	}
	return v, err
}

func GetRiversWithoutAlgae(userEmail string) []RiverWithId {
	rId, err := mgo.QueryRiverListByEmail(userEmail)
	if err != nil {
		return nil
	}
	res := make([]RiverWithId, 0)
	for _, obj := range rId {
		r := mgo.QueryRiverByIdWithoutAlgae(obj)
		res = append(res, RiverWithId{
			Id:      obj.Hex(),
			Name:    r.Name,
			Address: r.Address,
			Algae:   nil,
		})
	}
	return res
}

// SearchRiverByKeyWithoutAlgae  根据关键字搜索河流
func SearchRiverByKeyWithoutAlgae(userEmail string, keyword string) []RiverWithId {
	rId, err := mgo.QueryRiverListByEmail(userEmail)
	if err != nil {
		return nil
	}
	res := make([]RiverWithId, 0)
	for _, obj := range rId {
		r := mgo.QueryRiverByIdWithoutAlgae(obj)
		if ok, _ := regexp.MatchString(keyword, r.Name); ok {
			res = append(res, RiverWithId{
				Id:      obj.Hex(),
				Name:    r.Name,
				Address: r.Address,
				Algae:   nil,
			})
		}
	}
	return res
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
