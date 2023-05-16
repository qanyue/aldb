package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type River struct {
	Name    string   `json:"name" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Algae   []string `json:"algae" binding:"-"`
}
type RiverWithId struct {
	Id      string   `json:"id" binding:"required"`
	Name    string   `json:"name" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Algae   []string `json:"algae" binding:"-"`
}

type Alga struct {
	Name        string       `json:"name" binding:"required"`
	Src         string       `json:"src" binding:"required"`
	Annotations []Annotation `json:"annotations" binding:"-"`
}
type AlgaWithId struct {
	Id          string       `json:"id" binding:"-"`
	Name        string       `json:"name" binding:"required"`
	Src         string       `json:"src" binding:"required"`
	Annotations []Annotation `json:"annotations" binding:"-"`
}
type Tag struct {
	Name         string `json:"name" binding:"required"`
	ResourceName string `json:"resourceName" binding:"required"`
}

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

type Anno struct {
	AlgaId      string `json:"algaId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Tag         Tag    `json:"tag" binding:"required"`
}

type AlgaUpload struct {
	RiverId string `json:"riverId" binding:"required"`
	Alga    Alga   `json:"alga" binding:"required"`
}

type RiverAdd struct {
	River     River  `json:"river" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
}

func objectIdToString(id []primitive.ObjectID) []string {
	res := make([]string, 0)
	for _, v := range id {
		res = append(res, v.Hex())
	}
	return res
}
