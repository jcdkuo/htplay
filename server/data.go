package main

import "gopkg.in/mgo.v2/bson"

const (
	COL_USER = "user"
	COL_CARD = "card"
)

type Loc struct {
	Lng float64 `bson:"lng"`
	Lat float64 `bson:"lat"`
}

type DBUser struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Loc      Loc           `bson:"loc"`
	Img      bson.Binary   `bson:"img"`
	Skill    string        `json:"skill"`
}

func (this *Server) FindNearPeople(lng float64, lat float64, maxDistance float64, users *[]DBUser) error {

	var err error
	if session, col, err := this.sessionCopy(COL_USER); err == nil {
		defer session.Close()

		return col.Find(bson.M{"loc": bson.M{"$near": []float64{lng, lat}, "$maxDistance": maxDistance}}).All(users)
	}
	return err
}

func (this *Server) GetUser(username string, user *DBUser) error {
	var err error
	if session, col, err := this.sessionCopy(COL_USER); err == nil {
		defer session.Close()

		return col.Find(bson.M{"username": username}).One(user)
	}
	return err
}

func (this *Server) UpdateUser(user *DBUser) error {
	var err error
	if session, col, err := this.sessionCopy(COL_USER); err == nil {
		defer session.Close()

		return col.UpdateId(user.Id, user)
	}
	return err
}

func (this *Server) AddUser(user *DBUser) error {
	var err error
	if session, col, err := this.sessionCopy(COL_USER); err == nil {
		defer session.Close()

		return col.Insert(user)
	}
	return err
}
