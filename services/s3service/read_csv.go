package s3service

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Transaction struct {
	Id    string
	Date  string
	Trans string
}

func ReadCSV(bucket string, key string) (map[string]MontlyData, float64, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, 0, fmt.Errorf("failed s3 : %w", err)
	}

	process, totalbalance, err := ProcessCSV(obj.Body)

	if err != nil {
		return nil, 0, fmt.Errorf("error: %w", err)
	}

	return process, totalbalance, nil
}

func ProcessCSV(content io.ReadCloser) (map[string]MontlyData, float64, error) {
	defer content.Close()

	reader := csv.NewReader(content)
	var transactions []Transaction

	//read the first line and ignore it
	_, err := reader.Read()
	if err != nil {
		return nil, 0, fmt.Errorf("failde to read header: %w", err)
	}

	for {
		raw, err := reader.Read()
		//if last line error breaks
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, 0, fmt.Errorf("failed to read raw: %w", err)
		}

		transaction := Transaction{
			Id:    raw[0],
			Date:  raw[1],
			Trans: raw[2],
		}

		transactions = append(transactions, transaction)
	}

	fmt.Printf("Transactions: %s\n", transactions)
	processedData, totalBalance, err := ProcessData(transactions)
	if err != nil {
		return nil, 0, err
	}
	return processedData, totalBalance, nil
}

type MontlyData struct {
	NumberDebit  int32
	NumberCredit int32
	TotalDebit   float64
	TotalCredit  float64
}

func ProcessData(transactions []Transaction) (map[string]MontlyData, float64, error) {

	totalBalance := 0.0
	months := make(map[string]MontlyData)
	for key, value := range transactions {
		fmt.Println("Key:", key, "Value:", value)

		var montlyData MontlyData
		var card_kind string
		var month string
		var strtofloat float64
		var err error
		monthday := strings.Split(value.Date, "/")
		if len(monthday) > 0 {
			month = monthday[0]
			fmt.Println("Month:", month)
		} else {
			fmt.Println("Invalid date format")
			return nil, 0, fmt.Errorf("invalid format")

		}

		//debit cards
		amountd, isdebit := strings.CutPrefix(value.Trans, "-")

		if isdebit {
			strtofloat, err = strconv.ParseFloat(strings.TrimSpace(amountd), 64)
			if err != nil {
				return nil, 0, err
			}
			totalBalance -= strtofloat
			card_kind = "debit"
		}

		//credit cards
		amountc, iscredit := strings.CutPrefix(value.Trans, "+")
		if iscredit {
			strtofloat, err = strconv.ParseFloat(strings.TrimSpace(amountc), 64)
			if err != nil {
				return nil, 0, err
			}
			totalBalance += strtofloat
			card_kind = "credit"
		}

		_, ok := months[month]
		if ok {
			fmt.Println("Exists")
			montlyData = months[month]
			if card_kind == "debit" {
				montlyData.NumberDebit++
				montlyData.TotalDebit += strtofloat
			} else {
				montlyData.NumberCredit++
				montlyData.TotalCredit += strtofloat
			}

		} else {
			fmt.Println("Not exists")
			if card_kind == "debit" {
				montlyData.NumberDebit = 1
				montlyData.TotalDebit = strtofloat
			} else {
				montlyData.NumberCredit = 1
				montlyData.TotalCredit = strtofloat
			}

		}
		months[month] = montlyData

		fmt.Printf("Month: %s", month)
		fmt.Printf("Kinnd: %s", card_kind)
		monthint, err := strconv.Atoi(month)
		if err != nil {
			return nil, 0, err
		}
		monthname := time.Month(monthint)
		fmt.Printf("Month name: %s", monthname)

	}

	fmt.Printf("Total amount: %f", totalBalance)
	fmt.Printf("procesed months: %+v", months)

	return months, totalBalance, nil

}
