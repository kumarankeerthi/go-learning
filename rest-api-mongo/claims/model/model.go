package model

type Claims struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	EmployeeId  string `json:"employeeId,omitempty"`
}
