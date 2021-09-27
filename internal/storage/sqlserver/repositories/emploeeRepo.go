package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"payment-service/domain/employees"
	"payment-service/internal/storage/sqlserver"
)

type EmployeeRepository struct {
	db sqlserver.IDatabase
}

func NewEmployeeRepository(database sqlserver.IDatabase) *EmployeeRepository {
	return &EmployeeRepository{

		db: database,
	}
}

// CreateEmployee inserts an employee record
func (sq *EmployeeRepository) CreateEmployee(name string,
	location string) (int64, error) {
	ctx := context.Background()
	var err error

	if sq.db == nil {
		err = errors.New("CreateEmployee: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = sq.db.GetDB().PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO TestSchema.Employees (Name, Location) VALUES (@Name, @Location);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := sq.db.GetDB().Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", name),
		sql.Named("Location", location))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// ReadEmployees reads all employee records
func (sq *EmployeeRepository) ReadEmployees() ([]employees.Employee, error) {
	ctx := context.Background()
	var _employs []employees.Employee
	// Check if database is alive.
	err := sq.db.GetDB().PingContext(ctx)
	if err != nil {
		return _employs, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := sq.db.GetDB().QueryContext(ctx, tsql)
	if err != nil {
		return _employs, err
	}

	defer rows.Close()

	var count int
	var mod employees.Employee
	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return _employs, err
		}

		mod = employees.Employee{
			Id:       id,
			Name:     name,
			Location: location,
		}

		_employs = append(_employs, mod)
		//jsn, err := json.Marshal(mod)

		//fmt.Printf("employe %s \n",  jsn)
		count++
	}

	return _employs, nil
}

// UpdateEmployee updates an employee's information
func (sq *EmployeeRepository) UpdateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := sq.db.GetDB().PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name = @Name")

	// Execute non-query with named parameters
	result, err := sq.db.GetDB().ExecContext(
		ctx,
		tsql,
		sql.Named("Location", location),
		sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func (sq *EmployeeRepository) DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := sq.db.GetDB().PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name = @Name;")

	// Execute non-query with named parameters
	result, err := sq.db.GetDB().ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
