package model

import (
	"gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session

const database = "course"

// Init mongodb
func Init(mongoURL string) error {
	var err error
	mongoSession, err = mgo.Dial(mongoURL)
	if err != nil {
		return err
	}
	return nil
}
