package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// handle different response types
	"github.com/unrolled/render"

	"github.com/shiloa/apidemo/models"
)

var r = render.New()
var DB gorm.DB

// just a place holder for the root URL
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	r.JSON(w, http.StatusOK, map[string]string{"status": "it's alllllive!!!1"})
}

// GET /users/:id:w
func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	users := []models.User{}
	if id, ok := vars["id"]; ok {
		DB.Where("id = ?", id).Find(&users)
		if len(users) > 0 {
			r.JSON(w, http.StatusOK, users[0])
			//r.XML(w, http.StatusOK, users[0])
		} else {
			r.JSON(w, http.StatusOK, map[string]string{})
		}
	}
}

type SignupForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /users
func PostUserHandler(w http.ResponseWriter, req *http.Request) {
	// read request body
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	// handle any request errors
	if err != nil || body == nil {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false"})
	}

	// get form Data
	form := SignupForm{}
	err = json.Unmarshal(body, &form)

	// unmarshall errors
	if err != nil {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "error": err.Error()})
	}

	// handle form erros
	if form.Name == "" {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "error": "name cannot be empty"})
	}

	if form.Email == "" {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "error": "email cannot be empty"})
	}

	if form.Password == "" {
		r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "error": "password cannot be empty"})
	}

	user := models.CreateUser(form.Name, form.Email, form.Password, DB)
	r.JSON(w, http.StatusOK, user)
}

/**
curl -H "Content-Type: application/json" \
     -X PATCH -d '{"subscribed": true}' \
	 http://localhost:3000/users/02b41766-2292-473b-bb7d-0bfd9f4f10f5
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
	user := models.FindUser(vars["id"], DB)
	if user != nil {

		// update attributes from the request
		err = json.Unmarshal(body, user)

		if err != nil {
			r.JSON(w, http.StatusInternalServerError, map[string]string{"success": "false", "message": err.Error()})
		}

		// update the user and return
		DB.Save(&user)
		r.JSON(w, http.StatusOK, user)
	}
}

func GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	users := models.FindUsers(DB)
	r.JSON(w, http.StatusOK, users)
}
