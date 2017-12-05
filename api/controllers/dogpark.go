package controllers


import (
     "encoding/json"
     "fmt"
     "net/http"

     "github.com/julienschmidt/httprouter"
	 "github.com/user/goWebRestApp/api/models"
	"strings"

)

var Dogparks = []models.Dogpark{
	{
		"Boulder Dog Park", "Twin Pines", "1",
	},{
		"Boulder Dog Park2", "Twin Pines2", "2",
	},
}

type (
	// UserController represents the controller for operating on the User resource
	DogparkController struct{}
)

func NewDogparkController() *DogparkController {
	return &DogparkController{}
}

// GetDogpark retrieves an individual user resource
func (uc DogparkController) GetDogparks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
	u := Dogparks
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// GetDogpark retrieves an individual user resource
func (uc DogparkController) GetDogpark(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user

	u := models.Dogpark{}

	//decoder := json.NewDecoder(r.Body).Decode(&u)
	id := strings.TrimPrefix(r.URL.Path, "/dogpark/")

	for i := range Dogparks {

		if Dogparks[i].Id == id {
			// Found!
			u = Dogparks[i]
		}
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateDogpark creates a new user resource
func (uc DogparkController) CreateDogpark(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.Dogpark{}

	// Populate the user data
	//decoder := json.NewDecoder(r.Body).Decode(&u)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//log.Println(u.Name, u.Address, u.Id)

	Dogparks = append(Dogparks,u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)


	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveDogpark removes an existing user resource - not implemented
func (uc DogparkController) RemoveDogpark(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status for now
	id := strings.TrimPrefix(r.URL.Path, "/dogpark/")
	for i := range Dogparks {
		fmt.Printf("%d",i)
		if Dogparks[i].Id == id {
			Dogparks = append(Dogparks[:i], Dogparks[i+1:]...)
		}
	}


	w.WriteHeader(200)
}

