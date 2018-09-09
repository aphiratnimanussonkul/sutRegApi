package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Courseid type
type Courseid struct {
	CourseID map[string]string `bson:"courseid"`
}

//GetCourseid from courseCode
func GetCourseid(d string) (string, error) {
	id := "5b8bb93543aeadaf9f276467"
	if !bson.IsObjectIdHex(id) {
		return "", fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var c Courseid
	err := s.DB(database).C("courseid").FindId(objectID).One(&c)
	if err != nil {
		return "", fmt.Errorf("noon id")
	}

	v1 := c.CourseID[d]
	fmt.Println(v1)
	return v1, nil
}
