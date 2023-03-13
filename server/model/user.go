package model

import (
	"errors"
	"github.com/SukiEva/aldb/server/model/database"
)

func AddUser(obj Operator) error {
	if mgo.ExistsOperator(obj.Email) {
		return errors.New("email exists")
	}
	err := mgo.InsertOperator(&database.Operator{
		Name:        obj.Name,
		Password:    obj.Password,
		Email:       obj.Email,
		Access:      obj.Access,
		Annotations: nil,
	})
	return err
}

func UpdateUser(obj Operator) error {
	return mgo.UpsertOperator(obj.Email, obj.Name, obj.Password, obj.Access)
}

func GetUser(email string) Operator {
	user, err := mgo.QueryOperatorByEmail(email)
	if err != nil {
		return Operator{}
	}
	return Operator{
		Name:     user.Name,
		Password: "",
		Email:    user.Email,
		Access:   user.Access,
		CreateAt: user.CreateAt.Format("2006-01-02"),
	}
}

func GetUsers() []Operator {
	res := make([]Operator, 0)
	users, err := mgo.QueryOperators()
	if err != nil {
		return res
	}
	for _, obj := range users {
		res = append(res, Operator{
			Name:     obj.Name,
			Password: obj.Password,
			Email:    obj.Email,
			Access:   obj.Access,
			CreateAt: obj.CreateAt.Format("2006-01-02"),
		})
	}
	return res
}
func DeleteUser(email string) error {
	return mgo.DropOperator(email)
}

func ChangePassword(email string, newPassword string) error {
	return mgo.UpdatePassword(email, newPassword)
}
