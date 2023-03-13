package model

type River struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type Alga struct {
	Name  string `json:"name" binding:"required"`
	Src   string `json:"src" binding:"required"`
	River string `json:"river" binding:"required"`
}

type Annotation struct {
	Description string `json:"description" binding:"required"`
	Format      string `json:"format" binding:"required"`
	Url         string `json:"url" binding:"required"`
	CreateAt    string `json:"createAt" binding:"-"`
	UpdateAt    string `json:"updateAt" binding:"-"`
	Id          string `json:"id" binding:"-"`
}

type Operator struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Access   int    `json:"access" binding:"-"`
	CreateAt string `json:"createAt" binding:"-"`
}

type Anno struct {
	User        string `json:"user" binding:"required"`
	Alga        string `json:"alga" binding:"required"`
	Description string `json:"description" binding:"required"`
	Format      string `json:"format" binding:"required"`
	Url         string `json:"url" binding:"required"`
}
