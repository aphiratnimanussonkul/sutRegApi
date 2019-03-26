package model

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	CourseCode string `json:"coursecode" bson:"coursecode"`
	StudentId  string `json:"studentid" bson:"studentid"`
	Timestamp  time.Time
}

type DataRe struct {
	B55   int `json:"B55" `
	B56   int `json:"B56" `
	B57   int `json:"B57" `
	B58   int `json:"B58" `
	B59   int `json:"B59" `
	B60   int `json:"B60" `
	B61   int `json:"B61" `
	B62   int `json:"B62" `
	Other int `json:"other" `
}

type DataReCode struct {
	Sumcode []Sum
}

type Sum struct {
	CourseCode string `json:"coursecode" `
	Count      int    `json:"count" `
}

func Insert(data *Data) error {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()

	data.Timestamp = time.Now()

	err := s.DB(database).C("data").Insert(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "error"}
	}

	return nil

}

func Get() (error, DataRe) {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()

	var data []Data
	var dataRe DataRe

	err := s.DB(database).C("data").Find(nil).Sort("-timestamp").All(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}, dataRe
	}

	for _, d := range data {
		inputFmt := d.StudentId[:len(d.StudentId)-5]

		if inputFmt == "b58" || inputFmt == "B58" {
			dataRe.B58++
		} else if inputFmt == "b59" || inputFmt == "B59" {
			dataRe.B59++
		} else if inputFmt == "b60" || inputFmt == "B60" {
			dataRe.B60++
		} else if inputFmt == "b61" || inputFmt == "B61" {
			dataRe.B61++
		} else if inputFmt == "b62" || inputFmt == "B62" {
			dataRe.B62++
		} else if inputFmt == "b55" || inputFmt == "B55" {
			dataRe.B55++
		} else if inputFmt == "b56" || inputFmt == "B56" {
			dataRe.B56++
		} else if inputFmt == "b57" || inputFmt == "B57" {
			dataRe.B57++
		} else {
			dataRe.Other++
		}
	}

	return nil, dataRe

}

func GetCC() (error, DataReCode) {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()

	var data []Data
	var dataReCode DataReCode
	var sum []Sum

	m := make(map[string]int, 4000)

	err := s.DB(database).C("data").Find(nil).Sort("-timestamp").All(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}, dataReCode
	}

	for _, d := range data {
		// if dataRe.Planet[d.CourseCode] == nil {
		// 	dataRe.Planet[d.CourseCode] = 0
		// }

		m[d.CourseCode]++

	}

	for k, v := range m {
		// fmt.Printf("key[%s] value[%d]\n", k, v)
		sum = append(sum, Sum{k, v})
	}

	// sort.SliceStable(sum, func(i, j int) bool {
	// 	return sum[i].CourseCode < sum[j].CourseCode
	// })

	sort.SliceStable(sum, func(i, j int) bool {
		return sum[i].Count > sum[j].Count
	})

	dataReCode.Sumcode = sum[0:11]

	return nil, dataReCode

}

func GetCC2() (error, int) {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()

	var data []Data
	// var dataReCode DataReCode
	// var sum []Sum

	// m := make(map[string]int, 4000)

	err := s.DB(database).C("data").Find(nil).Sort("-timestamp").All(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}, len(data)
	}

	return nil, len(data)

}

func GetCourseData() (error, DataReCode) {
	s := mongoSession.Copy()
	defer s.Close()

	var data []Data
	var dataReCode DataReCode
	var sum []Sum

	m := make(map[string]int, 4000)

	err := s.DB(database).C("data").Find(nil).Sort("-timestamp").All(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}, dataReCode
	}

	for _, d := range data {
		if len(d.CourseCode) > 4 {
			inputFmt := d.CourseCode[:len(d.CourseCode)-3]
			if inputFmt == "522" {
				m["วิศวกรรมขนส่งและโลจิสติกส์"]++
			} else if inputFmt == "521" {
				m["วิศวกรรมเกษตรและอาหาร"]++
			} else if inputFmt == "523" {
				m["วิศวคอมพิวเตอร์"]++
			} else if inputFmt == "524" {
				m["วิศวกรรมเคมี"]++
			} else if inputFmt == "525" {
				m["วิศวกรรมเครื่องกล"]++
			} else if inputFmt == "526" {
				m["วิศวกรรมเซรามิก"]++
			} else if inputFmt == "527" {
				m["วิศวกรรมโทรคมนาคม"]++
			} else if inputFmt == "528" {
				m["วิศวกรรมพอลิเมอร์"]++
			} else if inputFmt == "529" {
				m["วิศวกรรมไฟฟ้า"]++
			} else if inputFmt == "530" {
				m["วิศวกรรมโยธา"]++
			} else if inputFmt == "531" {
				m["วิศวกรรมโลหการ"]++
			} else if inputFmt == "532" {
				m["วิศวกรรมสิ่งแวดล้อม"]++
			} else if inputFmt == "533" {
				m["วิศวกรรมอุตสาหการ"]++
			} else if inputFmt == "534" {
				m["เทคโนโลยีธรณี"]++
			} else if inputFmt == "535" {
				m["วิศวกรรมการผลิต"]++
			} else if inputFmt == "536" {
				m["วิศวกรรมยานยนต์"]++
			} else if inputFmt == "537" {
				m["วิศวกรรมอากาศยาน"]++
			} else if inputFmt == "538" {
				m["วิศวกรรมธรณี"]++
			} else if inputFmt == "539" {
				m["วิศวกรรมอิเล็กทรอนิกส์"]++
			} else if inputFmt == "540" {
				m["วิศวกรรมออกแบบผลิตภัณฑ์"]++
			} else if inputFmt == "541" {
				m["วิศวกรรมเครื่องมือ"]++
			} else if inputFmt == "551" {
				m["วิศวกรรมเมคคาทรอนิกส์"]++
			} else if inputFmt == "559" {
				m["วิศวกรรมโยธาและโครงสร้างพื้นฐาน"]++
			} else if inputFmt == "102" {
				m["เคมี"]++
			} else if inputFmt == "103" {
				m["คณิตศาสตร์"]++
			} else if inputFmt == "104" {
				m["ชีววิทยา"]++
			} else if inputFmt == "105" {
				m["ฟิสิกส์"]++
			} else if inputFmt == "114" {
				m["วิทยาศาสตร์การกีฬา"]++
			} else if inputFmt == "204" {
				m["เทคโนโลยีสารสนเทศ"]++
			} else if inputFmt == "235" {
				m["เทคโนโลยีการจัดการ"]++
			} else if inputFmt == "322" {
				m["เทคโนโลยีการผลิตพืช"]++
			} else if inputFmt == "323" {
				m["เทคโนโลยีการผลิตสัตว์"]++
			} else if inputFmt == "325" {
				m["เทคโนโลยีอาหาร"]++
			} else if inputFmt == "601" {
				m["แพทยศาสตร์"]++
			} else if inputFmt == "617" {
				m["อนามัยสิ่งแวดล้อม"]++
			} else if inputFmt == "618" {
				m["อาชีวอนามัยและความปลอดภัย"]++
			} else if inputFmt == "904" {
				m["ทันตแพทยศาสตร์"]++
			} else {
				m["other"]++
			}
		}
	}

	for k, v := range m {
		sum = append(sum, Sum{k, v})
	}

	sort.SliceStable(sum, func(i, j int) bool {
		return sum[i].Count < sum[j].Count
	})

	dataReCode.Sumcode = sum

	return nil, dataReCode
}

func GetA() (error, int) {

	//objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()

	var data []Data
	// var dataReCode DataReCode
	// var sum []Sum

	// m := make(map[string]int, 4000)

	var start = time.Now().AddDate(0, 0, -1)
	var end = time.Now()

	err := s.DB(database).C("data").Find(bson.M{
		"timestamp": bson.M{
			"$gt": start,
			"$lt": end,
		},
	}).Sort("-timestamp").All(&data)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}, len(data)
	}

	// fmt.Println(data)
	return nil, len(data)

}
