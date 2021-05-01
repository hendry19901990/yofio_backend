package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/hendry19901990/yofio_backend/models"
)

type CreditStore struct {
	Connection *gorm.DB
}

func (creditStore *CreditStore) Save(creditType *models.CreditType) {
	creditStore.Connection.NewRecord(creditType)
	creditStore.Connection.Create(creditType)
}

type CreditTypeStatisticsResult struct {
	Count   int32
	Average float32
	Success int32
}

func (creditStore *CreditStore) GetCreditTypeStatistics() (*models.CreditTypeStatistics, error) {

	var result []CreditTypeStatisticsResult

	creditStore.Connection.Raw("select count(*) as count, avg(investment) as average, success from credit_Type group by success").Scan(&result)

	statistics := models.CreditTypeStatistics{}

	assignmentsDone := int32(0)
	for _, res := range result {
		if res.Success == 1 {
			statistics.AssignmentsSuccess = res.Count
			statistics.InvestmentAverageSuccess = int32(res.Average)
			assignmentsDone += res.Count
		} else if res.Success == 0 {
			statistics.AssignmentsUnSuccess = res.Count
			statistics.InvestmentAverageUnSuccess = int32(res.Average)
			assignmentsDone += res.Count
		}
	}

	statistics.AssignmentsDone = assignmentsDone
	return &statistics, nil
}
