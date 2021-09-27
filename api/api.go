package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"payment-service/domain/employees"
	"payment-service/internal/storage/sqlserver"

	"github.com/gorilla/mux"
	"net/http"

	"payment-service/internal/storage/sqlserver/repositories"
)

type IController interface {
	returnAllArticles(w http.ResponseWriter, r *http.Request)
}

type EmployeeController struct {
	db sqlserver.IDatabase
}

func NewEmployeeController(db sqlserver.IDatabase) *EmployeeController {
	return &EmployeeController{
		db: db,
	}
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

var repo employees.IEmployeeGateway

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (c *EmployeeController) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	repo = repositories.NewEmployeeRepository(c.db)
	count, err := repo.ReadEmployees()
	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", len(count))

	json.NewEncoder(w).Encode(count)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func HandleRequests(c IController) {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", c.returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}
