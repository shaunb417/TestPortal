package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date	string
	time	string
}

func main(){
	http.HandleFunc("/",HomePage)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

	now :=time.Now() //finds the time right now
	HomePageVars := PageVariables { //stores the date and time struct on load
	Date: now.Format("02-01-2006"),
	Time: now.Format ("15:04:05"),
	}
	
	t, err :=template.ParseFiles("homepage.html") //parse the html file homepate.html
	if err := nil { // if there is an error
		log.Print("template parsing error: ", err)	
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct.
	if err != nil {
		log.Print{"template executing error: ", err}
	}

}