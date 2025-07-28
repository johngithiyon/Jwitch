package include

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)


func Handle_signup(w http.ResponseWriter,r *http.Request) {

    if r.Method == http.MethodGet {
		tmpl , err := template.ParseFiles("templates/signup.html")
		if err != nil {
			fmt.Println("Error",err)
			http.Error(w,"Template Error",500)
			return
		}
	
		tmpl.Execute(w,nil)
		return
	}
	

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirm_password := r.FormValue("confirm_password")

		if password != confirm_password {
             http.Error(w,"Confirm Password Wrong",400)
		} 

		_, err := Conn.Exec(context.Background(),
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		username, email, password)


		if err != nil {
              http.Error(w,"DB Error",500)
			  return
		} else {
			http.Redirect(w,r,"templates/index.html",http.StatusSeeOther)
			return
		}
	}
}

	func Handle_login(w http.ResponseWriter,r *http.Request) {

		if r.Method == http.MethodGet {
			tmpl,err := template.ParseFiles("templates/login.html")
			if err != nil {
				fmt.Println("Error",err)
				http.Error(w,"Template Error",500)
				return
			} 

			tmpl.Execute(w,nil)
			return
	    }

		if r.Method == http.MethodPost {
            var uname string

			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			err := Conn.QueryRow(context.Background(),"SELECT username FROM users WHERE email=$1 AND password=$2",
			email, password).Scan(&uname)

			if err != nil {
				http.Error(w,"DB Error",500)
				return
			} else {
				http.Redirect(w,r,"templates/index.html",http.StatusSeeOther)
		   }

			fmt.Println("Welcome Back",username)


		}
}

func Handle_form(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			fmt.Println("Error", err)
			http.Error(w, "Template Error", 500)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")


		if strings.TrimSpace(title) == "" || strings.TrimSpace(description) == "" {
			http.Error(w, "Please fill in all fields", http.StatusBadRequest)
			return
		}

		_, err := Conn.Exec(context.Background(),
			"INSERT INTO streams (title, description) VALUES ($1, $2)", title, description)
		if err != nil {
			http.Error(w, "DB Error", 500)
			return
		}

		
		http.Redirect(w, r, "http://localhost:8095/stream-options", http.StatusSeeOther)
	}
}


type Stream struct 
{
	 Id int
	 Title string 
	 Description string
}

var streamer_id int

func Fetch_form() ([] Stream) {

     var stream [] Stream
	 var title , description string
     var id int

	  rows, err := Conn.Query(context.Background(), "SELECT id,title, description FROM streams")

	  if err != nil {
		fmt.Println("Fetch error",err)
      }
     
	 defer func()  {
		rows.Close()
	} ()
	  
	  for rows.Next() {


		    err := rows.Scan(&id,&title,&description)

			if err != nil {
				fmt.Println("Scan error",err)
			}


		streamer_id = id

	    stream = append(stream, Stream{
			Id : id,
			Title: title,
			Description: description,
		}) 
	  }


      
	  return  stream
}


func Stopstream(w http.ResponseWriter,r *http.Request) {

	fmt.Println("stop endpoint is hit")
       
	if r.Method == http.MethodPost {

		 var streamID = streamer_id
           
		  body , err := io.ReadAll(r.Body)

		 if err != nil {
			 fmt.Println("Error in the read body",err)
		 }

		message := string(body)

		if message == "stopping" {
			_, err := Conn.Exec(context.Background(), "DELETE FROM streams WHERE id = $1", streamID)
			fmt.Println("It is sucessfully delete the stream")

		if err != nil {
			   fmt.Println("It did not return proper body noway",err)
		}


	}
	  
	}
}

var Steamtypevar string

func Steamtype(w http.ResponseWriter,r *http.Request) {
      
     body , err := io.ReadAll(r.Body)

	 if err != nil {
		fmt.Println("There is a problem in the type coming",err)
	 }

	 Steamtypevar= string(body)

}
	
    