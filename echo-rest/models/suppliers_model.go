package models

import (
	"echo-rest/common"
	"echo-rest/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

var supl common.Suppliers
var suplObj []common.Suppliers

// FetchSuppliers ...
func FetchSuppliers() (res Response, err error) {
	con := db.CreateCon()
	sqlQuery := `SELECT
					SupplierID,
					IFNULL(CompanyName,''),
					IFNULL(ContactName,''),
					IFNULL(ContactTitle,''),
					IFNULL(Address,'')
				FROM suppliers`
	rows, err := con.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return res, err
	}
	suplObj = nil
	for rows.Next() {
		err = rows.Scan(&supl.SupplierID, &supl.CompanyName, &supl.ContactName, &supl.ContactTitle, &supl.Address)
		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}
		suplObj = append(suplObj, supl)
	}
	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = suplObj
	return res, nil
}

// StoreSupplier ...
func StoreSupplier(e echo.Context) (res Response, err error) {
	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := `INSERT INTO suppliers (SupplierID, CompanyName, ContactName, ContactTitle, Address) VALUES (?,?,?,?,?)`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	result, err := stmt.Exec(req.SupplierID, req.CompanyName, req.ContactName, req.ContactTitle, req.Address)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]string{
		"SupplierID ADD":  req.SupplierID,
		"CompanyName ADD": req.CompanyName,
	}
	return res, nil
}

// UpdateSupplier ...
func UpdateSupplier(e echo.Context) (res Response, err error) {
	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := `UPDATE suppliers SET CompanyName = ?, ContactName = ?, ContactTitle = ?, Address = ? WHERE  SupplierID = ?`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address, req.SupplierID)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]string{
		"SupplierID Update": req.SupplierID,
		"CompanyName":       req.CompanyName,
	}
	return res, nil
}

// DeleteSupplier ...
func DeleteSupplier(e echo.Context) (res Response, err error) {
	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := `DELETE FROM suppliers WHERE  SupplierID = ?`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	result, err := stmt.Exec(req.SupplierID)
	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}
	fmt.Println(result)
	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = map[string]string{
		"SupplierID Deleted": req.SupplierID,
	}
	return res, nil
}
