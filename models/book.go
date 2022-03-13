//my first API based on the project of https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
