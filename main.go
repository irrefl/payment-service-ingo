package main

import (
	"bytes"
	"payment-service/api"
	"payment-service/internal/storage/sqlserver"

	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "payment-service/internal/storage/sqlserver"
)

func Init() {

	/*
		var Db = sqlserver.InitDatabase()

		d := repositories.New(Db)

		createID, err := d.CreateEmployee("Jake", "United States")
		if err != nil {
			log.Fatal("Error creating Employee: ", err.Error())
		}
		fmt.Printf("Inserted ID: %d successfully.\n", createID)


		count, err := d.ReadEmployees()
		if err != nil {
			log.Fatal("Error reading Employees: ", err.Error())
		}
		fmt.Printf("Read %d row(s) successfully.\n", count.Location)

	*/

}

func InitServer() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("error in fiber")
			}

			return errors.New("Managed error")
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	//app.routes(app)

	app.Listen(":3000")
}

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}

type HttpClient struct {
	logger *Logger
}

func (client *HttpClient) Get(url string) string {
	client.logger.Log("Getting " + url)

	// make an HTTP request
	return "my response from " + url
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{logger}
}

type ConcatService struct {
	logger *Logger
	client *HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	service.logger.Log("Running GetAll")

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url))
	}

	return result.String()
}

func NewConcatService(logger *Logger, client *HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}

func main() {
	/*logger := &Logger{}
	client := NewHttpClient(logger)
	service := NewConcatService(logger, client)

	result := service.GetAll(
		"http://google.com",
		"https://drewolson.org",
	)

	fmt.Println(result)

	*/

	api.Articles = []api.Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	var s sqlserver.IDatabase = sqlserver.SqlServer{}
	s.InitDatabase()

	//d := repositories.NewEmployeeRepository(Db)

	controller := api.NewEmployeeController(s)
	api.HandleRequests(controller)

	//Init()
}
