package models

// Data Model
type TData_req struct {
	User string `gorm:"column:first_name" json:"user"`
}

type TData_resp struct {
	Status  string `gorm:"column:status_id" json:"status"`
	RunTime int64  `gorm:"column:runTime" json:"runTime"`
	Value   int    `gorm:"column:value" json:"value"`
}
