package services

import (
	"errors"
	"fmt"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditYoFio struct {
}

func (credit *CreditYoFio) Assign(investment int32) (int32, int32, int32, error) {
	if investment%100 != 0 {
		return 0, 0, 0, errors.New("Quantity must be multiple of 100")
	}

	if investment < 1500 {
		return 0, 0, 0, errors.New("Quantity must be greater than 1500")
	}

	investment2 := investment
	var creditType700, creditType500, creditType300 int32

	if investment >= 700 {
		creditType700 = int32(investment / 700)
		investment = investment - (creditType700 * 700)
		if investment < 800 {
			creditType700--
			investment += 700
		}
	}

	for investment >= 300 && creditType700 >= 1 {
		if investment >= 500 {
			can := int32(investment / 500)
			creditType500 += can
			investment -= can * 500
			if investment < 300 {
				creditType500--
				investment += 500
			}
		}

		if investment >= 300 {
			can := int32(investment / 300)
			creditType300 += int32(investment / 300)
			investment -= can * 300
		}

		if (investment > 0 && creditType700 > 1) || creditType700 > creditType300 {
			creditType700--
			investment += 700
		}
	}

	fmt.Printf("investment received %d result %d \n", investment2, (creditType300*300 + creditType500*500 + creditType700*700))
	if investment > 0 {
		return 0, 0, 0, errors.New(fmt.Sprintf("It's impossible to make an assign with %d", investment2))
	}

	return creditType300, creditType500, creditType700, nil

}

func GetCreditAssigner() CreditAssigner {
	return &CreditYoFio{}
}
