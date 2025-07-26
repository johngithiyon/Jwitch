package main

import (
	"fmt"
	"html/template"
     "net/http"
	"jwitch/include"
	
)

func handle_home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")

    steam_info := include.Fetch_form()

    if err != nil {
        fmt.Println("Error",err)
        http.Error(w, "Template error",500)
        return
    }


    tmpl.Execute(w,steam_info)
}


func handle_live(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/livevideo.html")
}

func  streamoptions(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w,r,"templates/stream-options.html")
}

func cam_stream(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w,r,"templates/stream.html")
}

func screen_stream(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w,r,"templates/stream_screen.html")
}

func main() {
 
   include.GetConnection()
   

    http.HandleFunc("/home", handle_home)
    http.HandleFunc("/signup",include.Handle_signup)
    http.HandleFunc("/login",include.Handle_login)
    http.HandleFunc("/form.html",include.Handle_form)
    http.HandleFunc("/livevideo.html",handle_live)
    http.HandleFunc("/stop-stream",include.Stopstream)
    http.HandleFunc("/stream-options",streamoptions)
    http.HandleFunc("/stream.html",cam_stream)
    http.HandleFunc("/stream_screen.html",screen_stream)
    fmt.Println("Server is Listening .......")
    http.ListenAndServe(":8095", nil)
}
