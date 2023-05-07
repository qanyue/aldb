package database

import (
	"github.com/qanyue/aldb/server/model"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Annotation  标注使用Tag
type Annotation struct {
	field.DefaultField `bson:",inline"`
	Tag                model.Tag `json:"tag" bson:"tag"`
	Description        string    `json:"description" bson:"description"`
}

//  QueryAnnotationById  弃用
/*func (m *Mgo) QueryAnnotationById(obj primitive.ObjectID) *Annotation {
	var one *Annotation
	if err := annotation.Find(ctx, bson.M{"_id": obj}).One(&one); err != nil {
		return nil
	}
	return one
}
*/

func (m *Mgo) QueryAnnotation(id primitive.ObjectID) []Annotation {
	alga := m.QueryAlgaById(id)
	if alga == nil {
		zap.L().Error("algae获取标注失败", zap.String("algaeID", id.String()))
		return []Annotation{}
	}
	return alga.Annotations
}

// InsertAnnotation TODO 可能需要在前端判断是否需要标注一致
func (m *Mgo) InsertAnnotation(id primitive.ObjectID, annotation *Annotation) error {
	return algae.UpdateOne(ctx, bson.D{{"_id", id}},
		bson.M{"$addToSet": bson.M{"annotation": annotation}})
}

// DropAnnotation TODO 可能没有删除标注的功能
func (m *Mgo) DropAnnotation(id primitive.ObjectID, annotation *Annotation) error {
	return algae.Remove(ctx, bson.D{
		{"_id", id},
		{"annotations", bson.D{
			{"$elemMatch", bson.D{
				{"tag", annotation.Tag},
				{"description", annotation.Description},
			}},
		}},
	})
}
