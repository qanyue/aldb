package model

import (
	"errors"
	"github.com/qanyue/aldb/server/model/database"
)

func AddTag(obj Tag) error {
	if mgo.ExistsTag(obj.Name) {
		return errors.New("tag exists")
	}
	err := mgo.InsertTag(&database.Tag{
		Name:         obj.Name,
		ResourceName: obj.ResourceName,
	})
	return err
}

func GetTags() []Tag {
	tags, err := mgo.QueryTags()
	if err != nil {
		return nil
	}
	res := make([]Tag, 0)
	for _, obj := range tags {
		res = append(res, Tag{
			Name:         obj.Name,
			ResourceName: obj.ResourceName,
		})
	}
	return res
}
