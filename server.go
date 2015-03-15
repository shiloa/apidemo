package main

import (
	// database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// routing
	"github.com/gorilla/mux"

	// web server
	"github.com/codegangsta/negroni"

	// custom imports
	v1 "github.com/shiloa/apidemo/api.v1"
)

func initDB() gorm.DB {
	db, err := gorm.Open("mysql", "quixey:1234@/devportal?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		println("Error:", err)
		panic(err)
	}

	return db
}

func main() {

	// connect to db
	v1.DB = initDB()

	// establish router
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/", v1.HomeHandler).Methods("GET")
	router.HandleFunc("/users", v1.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", v1.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", v1.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", v1.PatchUserHandler).Methods("PATCH")

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":3001")
}
