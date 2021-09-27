package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"payment-service/api/sharedKernel"
	"payment-service/domain/employees"
	"payment-service/internal/storage/sqlserver"
	"payment-service/internal/storage/sqlserver/repositories"
	"strconv"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

var repo employees.IEmployeeGateway

type EmployeeController struct {
	db sqlserver.IDatabase
}

func NewEmployeeController(db sqlserver.IDatabase) *EmployeeController {
	return &EmployeeController{
		db: db,
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func enableJson(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAll")

	enableJson(w)
	logger := &sharedKernel.Logger{}
	client := sharedKernel.NewHttpClient(logger)
	concatService := sharedKernel.NewConcatService(logger, client)

	result := concatService.GetAll(
		"https://github.com",
		"https://khelpix.org",
	)

	repo = repositories.NewEmployeeRepository(c.db)
	allEmployees, err := repo.ReadEmployees()

	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
	}

	fmt.Printf("Read %d row(s) successfully.\n", len(allEmployees))
	fmt.Println(result)

	json.NewEncoder(w).Encode(allEmployees)

}

func (c *EmployeeController) GetOne(w http.ResponseWriter, r *http.Request) {
	enableJson(w)
	vars := mux.Vars(r)
	key := vars["id"]

	repo = repositories.NewEmployeeRepository(c.db)
	empl, _ := repo.ReadEmployees()

	for _, article := range empl {
		if strconv.Itoa(article.Id) == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func (c *EmployeeController) CreateNew(w http.ResponseWriter, r *http.Request) {

	repo = repositories.NewEmployeeRepository(c.db)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var employ employees.Employee
	json.Unmarshal(reqBody, &employ)

	repo.CreateEmployee(employ.Name, employ.Location)
	//Articles = append(Articles, employ)

	json.NewEncoder(w).Encode(employ)
}

func (c *EmployeeController) DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}
