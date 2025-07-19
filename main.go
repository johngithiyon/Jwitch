package main

import (
	"fmt"
	"html/template"
     "net/http"
	"jwitch/include"
	
)

func handle_home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        fmt.Println("Error",err)
        http.Error(w, "Template error",500)
        return
    }


    tmpl.Execute(w,nil)
}

 
func handle_form(w http.ResponseWriter,r *http.Request) {
    tmpl,err := template.ParseFiles("templates/form.html")
    if err != nil {
        http.Error(w,"Template Error",500)
        return
    }

    tmpl.Execute(w,nil)
}

func handle_stream(w http.ResponseWriter,r *http.Request) {
    tmpl,err := template.ParseFiles("templates/stream.html")
    if err != nil {
        http.Error(w,"Template Error",500)
        return
    }

    tmpl.Execute(w,nil)
}

func handle_live(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/livevideo.html")
}


func main() {
 
   include.GetConnection()
   

    http.HandleFunc("/", handle_home)
    http.HandleFunc("/signup",include.Handle_signup)
    http.HandleFunc("/login",include.Handle_login)
    http.HandleFunc("/form.html",handle_form)
    http.HandleFunc("/stream.html",handle_stream)
    http.HandleFunc("/livevideo.html",handle_live)
    fmt.Println("Server is Listening .......")
    http.ListenAndServe(":8095", nil)
}
