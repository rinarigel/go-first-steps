package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money "+
		"equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bobb", 25, -50, 4.2, 0.8, []string{"Dance", "Films"}}
	//fmt.Fprintf(w, "<b>Main text</b>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	//bob = User{name: "Bob", age: 25, money: -50}
	//bob = User{"Bob", 25, -50}

	handleRequest()
}
