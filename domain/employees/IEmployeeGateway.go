package employees

type IEmployeeGateway interface {
	CreateEmployee(name string, location string) (int64, error)
	ReadEmployees() ([]Employee, error)
	UpdateEmployee(name string, location string) (int64, error)
	DeleteEmployee(name string) (int64, error)
}
