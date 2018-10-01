package model

import (
	"fmt"
)

// Courseid type
type Courseid struct {
	ID       string `bson:"_id"`
	CourseID string
}

//GetCourseid from courseCode
// func GetCourseid(d string) (string, error) {
// 	id := "5b8bb93543aeadaf9f276467"
// 	if !bson.IsObjectIdHex(id) {
// 		return "", fmt.Errorf("invalid id")
// 	}
// 	objectID := bson.ObjectIdHex(id)
// 	s := mongoSession.Copy()
// 	defer s.Close()
// 	var c Courseid
// 	err := s.DB(database).C("courseid").FindId(objectID).One(&c)
// 	if err != nil {
// 		return "", fmt.Errorf("noon id")
// 	}

// 	v1 := c.CourseID[d]
// 	fmt.Println(v1)
// 	return v1, nil
// }
func GetCid(id string) (string, error) {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var n Courseid
	err := s.DB(database).C("abc").FindId(id).One(&n)
	if err != nil {
		return "", err
	}
	fmt.Println("x", n.CourseID)
	return n.CourseID, nil

}
