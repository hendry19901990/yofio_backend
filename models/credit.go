package models

import "time"

type InvestmentRequest struct {
	Investment int32 `json:"investment"`
}

type CreditType struct {
	Id            int64     `json:"-" gorm:"primary_key"`
	CreditType300 int32     `json:"credit_type_300" gorm:"column:credit_type_300"`
	CreditType500 int32     `json:"credit_type_500" gorm:"column:credit_type_500"`
	CreditType700 int32     `json:"credit_type_700" gorm:"column:credit_type_700"`
	Investment    int32     `json:"-" gorm:"column:investment"`
	Success       int32     `json:"-" gorm:"column:success"`
	DateCreated   time.Time `json:"-" gorm:"column:date_created;default:'current_timestamp'"`
}

func (CreditType) TableName() string {
	return "credit_Type"
}

type CreditTypeStatistics struct {
	AssignmentsDone            int32 `json:"assignments_done"`
	AssignmentsSuccess         int32 `json:"assignments_success"`
	AssignmentsUnSuccess       int32 `json:"assignments_unsuccess"`
	InvestmentAverageSuccess   int32 `json:"investment_average_success"`
	InvestmentAverageUnSuccess int32 `json:"investment_average_unsuccess"`
}
