package MVC

import (
	"io/ioutil"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

//var mgoSession mgo.Session

type Model struct {
	Title string
	Body []byte
}
/*
type Category struct {
	ID		bson.ObjectId	`json:"id" bson:"_id"`
	Title 	string			`json:"title" bson:"title"`
}

func ConnectToDB() {
	s, err := mgo.Dial("mongodb://localhost")

	if (err != nil) {
		panic(err)
	}

	//	Consistency mode Strong has the least load distribution, but the most guarantees.
	s.SetMode(mgo.Strong, true)

	mgoSession = *s
}
*/
func Save(m *Model, title string, body []byte) error {
	m.Title = title
	m.Body = body
	fileName := m.Title + ".txt"
	return ioutil.WriteFile(fileName, m.Body, 0600)
}

func LoadPage(m *Model, title string) error {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	m.Title = title
	m.Body = body
	return nil
}