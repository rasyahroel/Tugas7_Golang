package common

type Users struct {
	ID           int    `json:"id,omitempty"`
	NamaDepan    string `json:"nama_depan,omitempty"`
	NamaBelakang string `json:"nama_belakang,omitempty"`
	Email        string `json:"email, omitempty"`
	Username     string `json:"username, omitempty"`
	Password     string `json:"password, omitempty"`
}
type Customers struct {
	CustomerID   string `json:"CustomerID, omitempty"`
	CompanyName  string `json:"CompanyName, omitempty"`
	ContactName  string `json:"ContactName, omitempty"`
	ContactTitle string `json:"ContactTitle, omitempty"`
	Address      string `json:"Address, omitempty"`
	City         string `json:"City, omitempty"`
	Country      string `json:"Country, omitempty"`
	Phone        string `json:"Phone, omitempty"`
	PostalCode   string `json:"PostalCode, omitempty"`
}
type Employees struct {
	EmployeeID string `json:"EmployeeID"`
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	Title      string `json:"Title"`
	Address    string `json:"Address"`
}

type Suppliers struct {
	SupplierID   string `json:"SupplierID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
}
