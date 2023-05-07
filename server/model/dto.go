package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type River struct {
	Name    string   `json:"name" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Algae   []string `json:"algae" binding:"-"`
}

type Alga struct {
	Name        string       `json:"name" binding:"required"`
	Src         string       `json:"src" binding:"required"`
	Annotations []Annotation `json:"annotations" binding:"-"`
}

type Tag struct {
	Name         string `json:"name" binding:"required"`
	ResourceName string `json:"resourceName" binding:"required"`
}

// TODO tag标注使用tag来进行标注
type Annotation struct {
	Description string `json:"description"`
	CreateAt    string `json:"createAt" binding:"-"`
	UpdateAt    string `json:"updateAt" binding:"-"`
	Tag         *Tag   `json:"tag" binding:"required"`
}

type Operator struct {
	Name     string   `json:"name" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Access   int      `json:"access" binding:"-"`
	CreateAt string   `json:"createAt" binding:"-"`
	DataSet  []string `json:"dataSet" binding:"-"`
}

// TODO 标注逻辑修改
type Anno struct {
	AlgaId      string `json:"algaId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Tag         Tag    `json:"tag" binding:"required"`
}

func objectIdToString(id []primitive.ObjectID) []string {
	res := make([]string, 0)
	for _, v := range id {
		res = append(res, v.Hex())
	}
	return res
}
