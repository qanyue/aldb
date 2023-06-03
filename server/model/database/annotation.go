package database

import (
	"github.com/qiniu/qmgo/field"
	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type ModelTag struct {
	Name         string `json:"name" binding:"required"`
	ResourceName string `json:"resourceName" binding:"required"`
}

// Annotation  标注使用Tag
// segmentation 为标注的多边形
// Tag 为标注的类型
// Description 为标注的描述
type Annotation struct {
	field.DefaultField `bson:",inline"`
	Tag                ModelTag  `json:"tag" bson:"tag"`
	Description        string    `json:"description" bson:"description"`
	Segmentation       []float64 `json:"segmentation" bson:"segmentation"`
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

// InsertAnnotation 插入标注
func (m *Mgo) InsertAnnotation(id primitive.ObjectID, annotation *Annotation) error {
	opts := opts.UpdateOptions{
		nil,
		&options.UpdateOptions{
			ArrayFilters:             nil,
			BypassDocumentValidation: nil,
			Collation:                nil,
			Hint:                     nil,
			Upsert:                   &[]bool{true}[0],
		},
	}

	return algae.UpdateOne(ctx, bson.D{{"_id", id}},
		bson.M{"$push": bson.M{"annotations": annotation}}, opts)
}

// DropAnnotation TODO 可能没有删除标注的功能
func (m *Mgo) DropAnnotation(id primitive.ObjectID, annotation *Annotation) error {
	return algae.UpdateOne(ctx, bson.D{
		{"_id", id}}, bson.D{
		{"annotations", bson.D{
			{"$elemMatch", bson.D{
				{"tag", annotation.Tag},
				{"description", annotation.Description},
			}},
		}},
	})
}

// ExportAnnotation Export alga[] that each alga with the latest annotation
// an alga hold an array of annotation
//
//	a river hold an array of alga
//
// an annotation in hold create time
/*func (m *Mgo) ExportAnnotation(riverId primitive.ObjectID) []Alga {
	var r River
	river.Find(ctx, bson.M{"_id": riverId}).One(&r)
	var algas []Alga
	algae.Find(ctx, bson.M{"_id": bson.M{"$in": r.Algae}}).All(&algas)

}
*/
