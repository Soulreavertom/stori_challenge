package main

import (
	"os"
	"testing"
	"time"

	"github.com/Soulreavertom/stori_challenge/services/s3service"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	t.Logf(time.Now().Format("2006-01-02 15:04:05"))

	csvfile := "../../tmp/transactions.csv"
	f, err := os.Open(csvfile)

	if err != nil {
		t.Logf("failed can not open the file: %s", err.Error())
	}

	transactions, totalBalance, err := s3service.ProcessCSV(f)

	t.Logf("Total amount: %f", totalBalance)
	t.Logf("transactions: %+v", transactions)

	/*if err != nil {
		t.Logf("error processcsv: %s", err.Error())
	}

	type MontlyData struct {
		NumberDebit  int32
		NumberCredit int32
		TotalDebit   float64
		TotalCredit  float64
	}

	totalBalance := 0.0
	months := make(map[string]MontlyData)
	for key, value := range transactions {
		t.Log("Key:", key, "Value:", value)

		var montlyData MontlyData
		var card_kind string
		var month string
		var strtofloat float64
		monthday := strings.Split(value.Date, "/")
		if len(monthday) > 0 {
			month = monthday[0]          // Extract the month
			fmt.Println("Month:", month) // Output: Month: 12
		} else {
			fmt.Println("Invalid date format")
		}

		//debit cards
		amountd, isdebit := strings.CutPrefix(value.Trans, "-")

		if isdebit {
			strtofloat, err = strconv.ParseFloat(strings.TrimSpace(amountd), 64)
			if err == nil {
				totalBalance -= strtofloat
			}
			card_kind = "debit"
		}

		//credit cards
		amountc, iscredit := strings.CutPrefix(value.Trans, "+")
		if iscredit {
			strtofloat, err = strconv.ParseFloat(strings.TrimSpace(amountc), 64)
			if err == nil {
				totalBalance += strtofloat
			}
			card_kind = "credit"
		}

		_, ok := months[month]
		if ok {
			t.Logf("Exists")
			montlyData = months[month]
			if card_kind == "debit" {
				montlyData.NumberDebit++
				montlyData.TotalDebit += strtofloat
			} else {
				montlyData.NumberCredit++
				montlyData.TotalCredit += strtofloat
			}

		} else {
			t.Logf("Not exists")
			if card_kind == "debit" {
				montlyData.NumberDebit = 1
				montlyData.TotalDebit = strtofloat
			} else {
				montlyData.NumberCredit = 1
				montlyData.TotalCredit = strtofloat
			}

		}
		months[month] = montlyData

		t.Logf("Month: %s", month)
		t.Logf("Kinnd: %s", card_kind)
		monthint, err := strconv.Atoi(month)
		if err != nil {
			t.Logf("Ivalid number")
		}
		monthname := time.Month(monthint)
		t.Logf("Month name: %s", monthname)

	}

	t.Logf("Total amount: %f", totalBalance)
	t.Logf("procesed months: %+v", months)*/

}
