package include

import (
	"fmt"
	"context"
	"html/template"
	"net/http"
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
			 templ , err := template.ParseFiles("templates/index.html")
             if err != nil {
				return
			 }
			 templ.Execute(w,nil)
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
				templ , err := template.ParseFiles("templates/index.html")
				if err != nil {
				   return
				}
				templ.Execute(w,nil)
		   }

			fmt.Println("Welcome Back",username)


		}


}


