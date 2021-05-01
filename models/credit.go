package models

import "time"

type InvestmentRequest struct {
	Investment int32 `json:"investment"`
}

type CreditType struct {
	Id            int64     `json:"id" gorm:"primary_key"`
	CreditType300 int32     `json:"credit_type_300" gorm:"column:credit_type_300"`
	CreditType500 int32     `json:"credit_type_500" gorm:"column:credit_type_500"`
	CreditType700 int32     `json:"credit_type_700" gorm:"column:credit_type_700"`
	Investment    int32     `json:"Investment" gorm:"column:investment"`
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

/*
CREATE TABLE IF NOT EXISTS `test_db`.`credit_Type` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `credit_type_300` INT NULL,
  `credit_type_500` INT NULL,
  `credit_type_700` INT NULL,
  `invesment` INT NOT NULL,
  `success` TINYINT(1) NULL,
  `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

*/
