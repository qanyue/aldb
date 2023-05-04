package model

import (
	"github.com/qanyue/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAnnotationByAlga(algaName string) []Annotation {
	res := make([]Annotation, 0)
	// 根据藻类名称查找对应藻类
	alga, err := mgo.QueryAlgaByName(algaName)
	if err != nil {
		return res
	}

	for _, obj := range alga.Annotations {
		res = append(res, Annotation{
			Tag:         obj.Tag,
			Description: obj.Description,
			CreateAt:    obj.CreateAt.Format("2006-01-02 15:04"),
			UpdateAt:    obj.UpdateAt.Format("2006-01-02 15:04"),
		})
	}
	return res
}

func GetAnnotationByUser(userEmail string) []Annotation {
	res := make([]Annotation, 0)
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return res
	}
	for index, obj := range user.Annotations {
		anno := mgo.QueryAnnotationById(obj)
		if anno == nil { // 标注不存在
			// 标注列表删除该标注
			user.Annotations = append(user.Annotations[:index], user.Annotations[index:]...)
			_ = mgo.UpdateOperator(user.Id, user.Annotations)
			continue
		}
		res = append(res, Annotation{
			Description: anno.Description,
			Format:      anno.Format,
			Url:         anno.Url,
			CreateAt:    anno.CreateAt.Format("2006-01-02 15:04"),
			UpdateAt:    anno.UpdateAt.Format("2006-01-02 15:04"),
			Id:          anno.Id.Hex(),
		})
	}
	return res
}

func AddAnnotation(obj Anno) (interface{}, error) {
	return mgo.InsertAnnotation(&database.Annotation{
		Description: obj.Description,
		Format:      obj.Format,
		Url:         obj.Url,
	})
}

func DeleteAnnotation(id string) error {
	return mgo.DropAnnotation(id)
}

func UpdateAnnotation(obj Annotation) error {
	return mgo.UpsertAnnotation(obj.Description, obj.Format, obj.Url, obj.Id)
}

// BindToAlga 根据藻类名称添加到对应藻类图像 TODO BindToAlga测试
func BindToAlga(algaName string, anno Annotation) error {
	// 通过藻类名称查找对应的藻类
	alga, err := mgo.QueryAlgaByName(algaName)
	if err != nil {
		return err
	}
	// 添加到藻类的标注列表
	var annos []database.Annotation
	if alga.Annotations == nil {
		annos = []database.Annotation{
			{Description: anno.Description, Tag: anno.Tag},
		}
	} else {
		annos = append(alga.Annotations, database.Annotation{
			Description: anno.Description,
			Tag:         anno.Tag,
		})
	}
	// 更新藻类数据
	err = mgo.UpdateAlga(alga.Id, annos)
	return err
}

// BindToUser 根据用户邮箱绑定对应用户
func BindToUser(userEmail string, id primitive.ObjectID) error {
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return err
	}
	var anno []primitive.ObjectID
	if user.Annotations == nil {
		anno = []primitive.ObjectID{id}
	} else {
		anno = append(user.Annotations, id)
	}
	err = mgo.UpdateOperator(user.Id, anno)
	return err
}
