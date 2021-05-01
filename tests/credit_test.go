package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hendry19901990/yofio_backend/models"
	"github.com/hendry19901990/yofio_backend/services"
)

func TestCreditAssigner(t *testing.T) {
	c := services.GetCreditAssigner()
	creditType300, creditType500, creditType700, _ := c.Assign(3000)
	resultSuccess := creditType300*300 + creditType500*500 + creditType700*700
	assert.Equal(t, int32(3000), resultSuccess, "The three values should be the same.")

	creditType300, creditType500, creditType700, _ = c.Assign(1600)
	resultUnSuccess := creditType300*300 + creditType500*500 + creditType700*700
	assert.Equal(t, int32(0), resultUnSuccess, "The result should be zero.")
}

func TestSuccessCreditAssignment(t *testing.T) {
	investOk := `{"investment": 3000}`
	credit, err := MakeRequest("http://localhost:9090/api/credit-assignment", []byte(investOk))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	resultSuccess := credit.CreditType300*300 + credit.CreditType500*500 + credit.CreditType700*700
	assert.Equal(t, int32(3000), resultSuccess, "The three values should be the same.")
}

func TestUnSuccessCreditAssignment(t *testing.T) {
	investOk := `{"investment": 1600}`
	credit, err := MakeRequest("http://localhost:9090/api/credit-assignment", []byte(investOk))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	resultUnSuccess := credit.CreditType300*300 + credit.CreditType500*500 + credit.CreditType700*700
	assert.Equal(t, int32(0), resultUnSuccess, "The result should be zero.")
}

func TestStatistics(t *testing.T) {
	statisticsTest, err := MakeRequest("http://localhost:9090/api/statistics", []byte("{}"))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	assert.NotNil(t, statisticsTest, "The result shouldn't be empty.")
}

func MakeRequest(url string, jsonStr []byte) (*models.CreditType, error) {
	fmt.Println("URL:> ", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	credit := models.CreditType{}
	if resp.StatusCode == 200 {
		if errParse := json.Unmarshal(body, &credit); errParse != nil {
			return nil, errParse
		}
	}

	return &credit, nil
}
