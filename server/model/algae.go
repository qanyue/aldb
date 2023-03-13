package model

import (
	"errors"
	"github.com/SukiEva/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetData() []Alga {
	res := make([]Alga, 0)
	algae, err := mgo.QueryAlgae()
	if err != nil {
		return res
	}
	for _, obj := range algae {
		river := mgo.QueryRiverById(obj.River)
		res = append(res, Alga{
			Name:  obj.Name,
			Src:   obj.Src,
			River: river.Name,
		})
	}
	return res
}

func AddAlga(obj Alga) (interface{}, error) {
	river := mgo.QueryRiverByName(obj.River)
	return mgo.InsertAlga(&database.Alga{
		Name:        obj.Name,
		Src:         obj.Src,
		River:       river.Id,
		Annotations: nil,
	})
}

func BindToRiver(riverName string, id primitive.ObjectID) error {
	river := mgo.QueryRiverByName(riverName)
	var algae []primitive.ObjectID
	if algae == nil {
		algae = []primitive.ObjectID{id}
	} else {
		algae = append(algae, id)
	}
	return mgo.UpdateRiver(river.Id, algae)
}

func AddRiver(obj River) error {
	if mgo.ExistsRiver(obj.Name) {
		return errors.New("river exists")
	}
	err := mgo.InsertRiver(&database.River{
		Name:    obj.Name,
		Address: obj.Address,
	})
	return err
}

func GetRivers() []River {
	rivers, err := mgo.QueryRivers()
	if err != nil {
		return nil
	}
	res := make([]River, 0)
	for _, obj := range rivers {
		res = append(res, River{
			Name:    obj.Name,
			Address: obj.Address,
		})
	}
	return res
}

func SearchAlga(key string) []Alga {
	res := make([]Alga, 0)
	mp := make(map[string]struct{})
	// 精确查找
	if alga, err := mgo.QueryAlgaByName(key); err == nil {
		// 只有找到后才会执行
		river := mgo.QueryRiverById(alga.River)
		mp[alga.Name] = struct{}{}
		res = append(res, Alga{
			Name:  alga.Name,
			Src:   alga.Src,
			River: river.Name,
		})
	}
	// 模糊查找
	if algae, err := mgo.QueryAlgaByKey(key); err == nil {
		for _, alga := range algae {
			// 去重复
			if _, ok := mp[alga.Name]; ok {
				continue
			}
			mp[alga.Name] = struct{}{}
			// 只有找到后才会执行
			river := mgo.QueryRiverById(alga.River)
			res = append(res, Alga{
				Name:  alga.Name,
				Src:   alga.Src,
				River: river.Name,
			})
		}
	}
	return res
}
