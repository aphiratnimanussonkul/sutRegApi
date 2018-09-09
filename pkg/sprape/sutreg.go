package sprape

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/labstack/echo"
	"golang.org/x/net/html/charset"
)

// dt  ไลด์ว่างของ dayTime
var dt []model.DayTime

var d model.DReg
var data model.Datajsonn

// stetus 0 > เริ่มใหม่, 1 > ต่อ
var stetuss int
var checkk int

func deleteDatax() {
	var new1 model.Datajsonn
	var new2 model.DReg

	data = new1
	d = new2
}

func exampleScrape(url string) {
	// Request the HTML page.

	defer timer()()

	stetuss = 0
	checkk = 00
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	contentType := res.Header.Get("Content-Type") // Optional, better guessing
	utf8reader, err := charset.NewReader(res.Body, contentType)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(utf8reader)
	if err != nil {
		log.Fatal(err)
	}

	checkk := 99

	//type aa []string
	// Find the review items
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title

		band := s.Find("td").Text()

		//fmt.Println(i, band)

		if i == 6 {
			c := chunkString(band, 6)
			data.Coursecode = c[0]
			data.NameEnglish = c[1]
		}
		if i == 7 {
			data.NameThai = band

		}
		if i == 9 {
			data.Credit = band
		}

		if checkk == i {
			//fmt.Println(i)

			stetuss = 1
		}
		if strings.Contains(band, "กลุ่มวันเวลาห้องอาคารเรียนที่นั่ง(เปิด-ลง-เหลือ)หมวด") && i != 0 {
			checkk = i + 1

		}
		if stetuss == 2 {
			if strings.Contains(band, "หมายเหตุ") {
				s := strings.Split(band, "หมายเหตุ:")
				t := strings.TrimSpace(s[1])
				d.Note = t
			}
			data.Datareg = append(data.Datareg, d)
			d.T = ""
			d.Mid = model.Date{}
			d.Final = model.Date{}
			d.Note = ""
			d.DayTime = dt
			stetuss = 1
		}

		if stetuss == 1 {
			adddata(band)
			//fmt.Println("two", i)
			if strings.Contains(band, "อาจารย์") {
				s := strings.Split(band, "อาจารย์:")
				t := strings.TrimSpace(s[1])
				d.T = t
			} else if strings.Contains(band, "สอบกลางภาค") {
				s := strings.Split(band, "สอบกลางภาค:")
				t := strings.TrimSpace(s[1])
				d.Mid = model.Date{t, ""}
			} else if strings.Contains(band, "สอบประจำภาค:") {
				s := strings.Split(band, "สอบประจำภาค:")
				t := strings.TrimSpace(s[1])
				d.Final = model.Date{t, ""}
				stetuss = 2
			}
		}

	})
}

func adddata(band string) {
	if strings.Contains(band, "จันทร์") {
		s := strings.Split(band, "จันทร์")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "จันทร์")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[1], time[1]})

	} else if strings.Contains(band, "อังคาร") {
		s := strings.Split(band, "อังคาร")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "อังคาร")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "พุธ") {
		s := strings.Split(band, "พุธ")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "พุธ")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "พฤหัสบดี") {
		s := strings.Split(band, "พฤหัสบดี")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "พฤหัสบดี")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "ศุกร์") {
		s := strings.Split(band, "ศุกร์")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "ศุกร์")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "เสาร์") {
		s := strings.Split(band, "เสาร์")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "เสาร์")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "อาทิตย์") {
		s := strings.Split(band, "อาทิตย์")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		v := strings.Split(band, "F")
		vv := strings.Split(v[0], "B")
		vvv := strings.Split(vv[0], "L")
		v2 := strings.Split(vvv[0], s[0])

		time := strings.Split(v2[1], "อาทิตย์")
		day := strings.Split(v2[1], time[1])

		d.DayTime = append(d.DayTime, model.DayTime{day[0], time[1]})
	} else if strings.Contains(band, "ไม่มีข้อมูล") {
		s := strings.Split(band, "ไม่มีข้อมูล")
		t := strings.TrimSpace(s[0])
		if t != "" {
			d.Sec = t
		}
		d.DayTime = append(d.DayTime, model.DayTime{"ไม่มีข้อมูล", "ไม่มีข้อมูล"})
	}
}
func ch(band string) bool {
	if strings.Contains(band, "จันทร์") {
		return true
	} else if strings.Contains(band, "อังคาร") {
		return true
	}
	return false
}

// GetDataReg is schedule data from reg
func GetDataReg(cid string, c echo.Context) error {
	defer deleteDatax()

	exampleScrape("http://reg4.sut.ac.th/registrar/class_info_2.asp?backto=home&option=0&courseid=" + cid + "&acadyear=2561&semester=1&avs972184082=6")
	//fmt.Println("#!", data)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&data)

}

func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		log.Println(diff)
	}
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]), string(runes[nn:]))
		break
	}
	return chunks
}
