package main

import (
	"fmt"
	"html/template"
     "net/http"
	"jwitch/include"
	
)

type homedata struct {
    Steam_info [] include.Stream
    Steamtype  string
}

func handle_home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")

    data := homedata{
           Steam_info: include.Fetch_form(),
           Steamtype: include.Steamtypevar,
    }

    if err != nil {
        fmt.Println("Error",err)
        http.Error(w, "Template error",500)
        return
    }


    tmpl.Execute(w,data)
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

func handle_screen(w http.ResponseWriter,r *http.Request) {
    http.ServeFile(w,r,"templates/livescreen.html")
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
    http.HandleFunc("/stream_screen",screen_stream)
    http.HandleFunc("/steamtype",include.Steamtype)
    http.HandleFunc("/livescreen.html",handle_screen)
    fmt.Println("Server is Listening .......")
    http.ListenAndServe(":8095", nil)
}
