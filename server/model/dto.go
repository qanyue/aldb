package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// River 在本系统为数据集概念,algae 为藻类图片集
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

// Alga 藻类图片元数据
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

// Tag 标签
type Tag struct {
	Name         string `json:"name" binding:"required"`
	ResourceName string `json:"resourceName" binding:"required"`
}

// Annotation 标注数据
type Annotation struct {
	Description string `json:"description"`
	CreateAt    string `json:"createAt" binding:"-"`
	UpdateAt    string `json:"updateAt" binding:"-"`
	// 坐标点
	Segmentation []float64 `json:"segmentation" binding:"required"`
	Tag          *Tag      `json:"tag" binding:"required"`
}

type Operator struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Access   int    `json:"access" binding:"-"`
	CreateAt string `json:"createAt" binding:"-"`
	// 用户拥有的数据集，即river
	DataSet []string `json:"dataSet" binding:"-"`
}

// Anno 将标注数据与藻类图片关联时使用
type Anno struct {
	AlgaId       string    `json:"algaId" binding:"required"`
	Description  string    `json:"description"`
	Tag          Tag       `json:"tag" binding:"required"`
	Segmentation []float64 `json:"segmentation" binding:"required"`
}

// AlgaUpload 上传藻类图片时使用,将图片与数据集关联
type AlgaUpload struct {
	RiverId string `json:"riverId" binding:"required"`
	Alga    Alga   `json:"alga" binding:"required"`
}

type RiverAdd struct {
	River     River  `json:"river" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
}

// 将ObjectId转换为string
func objectIdToString(id []primitive.ObjectID) []string {
	res := make([]string, 0)
	for _, v := range id {
		res = append(res, v.Hex())
	}
	return res
}
