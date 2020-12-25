package models

import (
	"echo-rest/common"
	"echo-rest/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

// FetchEmployees ...
func FetchEmployees() (res Response, err error) {
	var errs *ex.AppError
	var emp common.Employees
	var empObj []common.Employees
	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()
	con := db.CreateCon()
	sqlQuery := `SELECT 
					EmployeeID, 
					IFNULL(LastName,'') LastName,
					IFNULL(FirstName,'') FirstName,
					IFNULL(Title,'') Title,
					IFNULL(Address,'') Address
				FROM employees`
	rows, err := con.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return res, err
	}
	empObj = nil
	for rows.Next() {
		err = rows.Scan(&emp.EmployeeID, &emp.LastName, &emp.FirstName, &emp.Title, &emp.Address)
		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}
		empObj = append(empObj, emp)
	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = empObj
	return res, nil
}

// StoreEmployee ...
func StoreEmployee(e echo.Context) (res Response, err error) {
	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()
	req := new(common.Employees)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := `INSERT INTO employees (LastName,FirstName, Title, Address) VALUES (?,?,?,?)`
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	result, err := stmt.Exec(req.LastName, req.FirstName, req.Title, req.Address)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]string{
		"Employees ADD": req.FirstName,
	}
	return res, nil
}

// UpdateEmployee ...
func UpdateEmployee(e echo.Context) (res Response, err error) {
	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()
	req := new(common.Employees)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := `UPDATE employees SET LastName = ?, FirstName = ?, Title = ?, Address = ? WHERE EmployeeID = ?`
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	result, err := stmt.Exec(req.LastName, req.FirstName, req.Title, req.Address, req.EmployeeID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]string{
		"row affected :": req.EmployeeID,
	}
	return res, nil
}

// DeleteEmployee ...
func DeleteEmployee(e echo.Context) (res Response, err error) {
	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()
	req := new(common.Employees)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := `DELETE FROM employees WHERE  EmployeeID = ?`
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.EmployeeID)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]string{
		"row deleted :": req.EmployeeID,
	}
	return res, nil
}
