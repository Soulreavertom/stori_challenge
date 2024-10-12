package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Soulreavertom/stori_challenge/db"
	"github.com/Soulreavertom/stori_challenge/models"
	"github.com/Soulreavertom/stori_challenge/services/emailservice"
	"github.com/Soulreavertom/stori_challenge/services/s3service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Println("Hello world")
	db.InitDB()
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Request: %+v\n", request)

	email := request.PathParameters["email"]
	fmt.Printf("email: %s\n", email)

	transactions, totalbalance, err := s3service.ReadCSV("twads", "transactions.csv")

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, err
	}

	subject := "Stori account"
	email_body := ""

	email_body_h := "<div style='width: 100%; background: #003a40; padding: 30px 0;'><div style='width:100%; background:#ffffff;'>" +
		"<div style='max-width: 600px; background: #ffffff; margin: auto; position: absolute; top: 50px; left: 0; right: 0; bottom: 0px;'>" +
		"<div style='text-align: center;'><img style='max-width: 250px;' src='https://api-redeco.certificacionpld.com/images/stori_card.png'></div>" +
		"<div style='text-align: center;'><h5>BALANCE</h5></div>" +
		"<div style='padding: 10px;'>" +
		"<table width='100%' cellspacing='3' cellpadding='3' border='1' style='border: 1px solid black; border-collapse: collapse;'>" +
		"<tr><th>Transactions</th><th>Montly average</th></tr>"

	email_body_b := ""
	ttd := 0.0
	tnd := 0.0

	ttc := 0.0
	tnc := 0.0
	for key, value := range transactions {
		monthint, err := strconv.Atoi(key)
		if err != nil {
			fmt.Printf("ivalid number")
		}
		monthname := time.Month(monthint)

		fmt.Printf("Month name: %s", monthname)

		numtrans := fmt.Sprintf("%d", value.NumberCredit+value.NumberDebit)
		averagedeb := fmt.Sprintf("%.2f", value.TotalDebit/float64(value.NumberDebit))
		averagecred := fmt.Sprintf("%.2f", value.TotalCredit/float64(value.NumberCredit))
		email_body_b += "<tr style='border-top: solid #000000 1px;'>" +
			"<td>Number of transactions in " + monthname.String() + ": " + numtrans + "</td>" +
			"<td>" +
			"Average debit amount: -" + averagedeb + "<br>" +
			"Average credit amount: " + averagecred +
			"</td>" +
			"</tr>"

		ttd += value.TotalDebit
		tnd += float64(value.NumberDebit)

		ttc += value.TotalCredit
		tnc += float64(value.NumberCredit)
	}

	twodecimalstb := fmt.Sprintf("%.2f", totalbalance)

	twodecimalTAD := fmt.Sprintf("%.2f", ttd/tnd)
	twodecimalTAC := fmt.Sprintf("%.2f", ttc/tnc)

	email_body_f := "</table>" +
		"<h4><b>Total balance is " + twodecimalstb + "</b></h4>" +
		"<h4><b>Total average debit amount is -" + twodecimalTAD + "</b></h4>" +
		"<h4><b>Total average credit amount is " + twodecimalTAC + "</b></h4>" +
		"</div></div></div></div>"
	email_body = email_body_h + email_body_b + email_body_f

	//This is not working because the way that works lambda, after you return a response lambda finish all process
	// Asynchronously send email
	/*go func() {
		emailservice.SendEmail(subject, email_body, email)
	}()*/

	for key, value := range transactions {
		fmt.Println("Key:", key, "Value:", value)
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	savereport := models.TransactionReport{
		Email:              email,
		TotalBalance:       totalbalance,
		TotalAverageDebit:  ttd / tnd,
		TotalAverageCredit: ttc / tnc,
		CreatedAt:          now,
	}

	saved, err := savereport.Save()
	if err != nil {
		fmt.Printf("Error saving: %s", err.Error())
	}
	fmt.Printf("last inserted id: %d", saved)

	emailservice.SendEmail(subject, email_body, email)

	reports, err := savereport.GetAll()
	if err != nil {
		fmt.Printf("Get errors failed: %s", err.Error())
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(reports)}, nil
}
