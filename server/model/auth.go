package model

import "github.com/SukiEva/aldb/server/model/database"

var mgo = database.Mgo{}

func CheckAuth(email, password string) bool {
	_, err := mgo.CheckOperator(email, password)
	if err != nil {
		return false
	}
	return true
}

func CheckAdmin(email string) bool {
	user, err := mgo.QueryOperatorByEmail(email)
	if err != nil || user.Access == 0 {
		return false
	}
	return true
}
