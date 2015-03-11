package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"

	// database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// routing
	"github.com/gorilla/mux"

	// web server
	"github.com/codegangsta/negroni"

	// handle different response types
	"github.com/unrolled/render"

	// custom imports
	"github.com/shiloa/apidemo/models"
)

var (
	db gorm.DB
	r  *render.Render
)

func initDB() gorm.DB {
	db, err := gorm.Open("mysql", "quixey:1234@/devportal?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		println("Error:", err)
		panic(err)
	}

	return db
}

// just a place holder for the root URL
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	r.JSON(w, http.StatusOK, map[string]string{"status": "it's alllllive!!!1"})
}

// GET /users/:id:w
func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	users := []models.User{}
	if id, ok := vars["id"]; ok {
		db.Where("id = ?", id).Find(&users)
		if len(users) > 0 {
			r.JSON(w, http.StatusOK, users[0])
			//r.XML(w, http.StatusOK, users[0])
		} else {
			r.JSON(w, http.StatusOK, map[string]string{})
		}
	}
}

/**
curl -H "Content-Type: application/json" -X PATCH -d '{"subscribed": true}' http://localhost:3000/users/02b41766-2292-473b-bb7d-0bfd9f4f10f5
*/
// PATCH /users/:id
func PatchUserHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// read request body
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	// handle any request errors
	if err != nil || body == nil {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false"})
	}

	// fetch relevant user
	users := []models.User{}
	db.Where("id = ?", vars["id"]).Find(&users)
	if len(users) > 0 {
		user := users[0]
		err = json.Unmarshal(body, &user)
		if err != nil {
			r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "message": err.Error()})
		}

		// update the user and return
		db.Save(&user)
		r.JSON(w, http.StatusOK, user)
	}
}

func GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	users := []models.User{}
	db.Find(&users)
	r.JSON(w, http.StatusOK, users)
}

func main() {

	// connect to db
	db = initDB()

	// initialize renderer
	r = render.New()

	// establish router
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")

	router.HandleFunc("/users/{id}", PatchUserHandler).Methods("PATCH")

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":3001")
}
