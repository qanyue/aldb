package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Annotation struct {
	field.DefaultField `bson:",inline"`
	Description        string `json:"description" bson:"description"`
	Format             string `json:"format" bson:"format"`
	Url                string `json:"url" bson:"url"`
}

func (m *Mgo) QueryAnnotationById(obj primitive.ObjectID) *Annotation {
	var one *Annotation
	if err := annotation.Find(ctx, bson.M{"_id": obj}).One(&one); err != nil {
		return nil
	}
	return one
}

func (m *Mgo) InsertAnnotation(obj *Annotation) (interface{}, error) {
	res, err := annotation.InsertOne(ctx, obj)
	return res.InsertedID, err
}

func (m *Mgo) DropAnnotation(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return annotation.Remove(ctx, bson.M{"_id": oid})
}

func (m *Mgo) UpsertAnnotation(description, format, url string, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return annotation.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": bson.M{"description": description, "format": format, "url": url}})
}
