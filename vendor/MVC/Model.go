package MVC

import (
	"io/ioutil"
)

type Model struct {
	Title string
	Body []byte
}
/*
func (m Model) GetFromDatabase() {

}

func (m Model) PutInDatabase() {

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