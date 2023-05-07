package model

import (
	"github.com/qanyue/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func GetAlgaData(id primitive.ObjectID) *Alga {
	var a *Alga
	alga := mgo.QueryAlgaById(id)
	if alga == nil {
		return &Alga{}
	}
	a.Src = alga.Src
	a.Name = alga.Name
	annos := make([]Annotation, 0)
	for _, obj := range alga.Annotations {
		annos = append(annos, Annotation{
			Description: obj.Description,
			CreateAt:    obj.CreateAt.Format("2006-01-02 15:04"),
			UpdateAt:    obj.UpdateAt.Format("2006-01-02 15:04"),
			Tag:         &obj.Tag,
		})
	}
	a.Annotations = annos
	return a
}

func AddAlga(rid primitive.ObjectID, obj Alga) error {
	aId, err := mgo.InsertAlga(&database.Alga{
		Name:        obj.Name,
		Src:         obj.Src,
		Annotations: []database.Annotation{},
	})
	if err != nil {
		zap.L().Error("添加藻类图片失败", zap.String("info", err.Error()))
		return err
	}
	return BindToRiver(rid, aId.(primitive.ObjectID))
}

func BindToRiver(rId primitive.ObjectID, aId primitive.ObjectID) error {
	river := mgo.QueryRiverById(rId)
	var algae []primitive.ObjectID
	if algae == nil {
		algae = []primitive.ObjectID{aId}
	} else {
		algae = append(algae, aId)
	}
	return mgo.UpdateRiverAlgae(river.Id, algae)
}

func SearchAlga(id primitive.ObjectID, key string) []Alga {
	res := make([]Alga, 0)
	mp := make(map[string]struct{})
	// 精确查找
	if algae := mgo.QueryAlgaByName(id, key); len(algae) > 0 {
		// 只有找到后才会执行
		for _, v := range algae {
			mp[v.Name] = struct{}{}
			res = append(res, Alga{
				Name:        v.Name,
				Src:         v.Src,
				Annotations: DataBaseAnnoToModelAnno(v.Annotations),
			})
		}
	}
	// 模糊查找
	if algae := mgo.QueryAlgaByKey(id, key); algae != nil {
		for _, v := range algae {
			// 去重复
			if _, ok := mp[v.Name]; ok {
				continue
			}
			mp[v.Name] = struct{}{}
			// 只有找到后才会执行
			res = append(res, Alga{
				Name:        v.Name,
				Src:         v.Src,
				Annotations: DataBaseAnnoToModelAnno(v.Annotations),
			})
		}
	}
	return res
}
