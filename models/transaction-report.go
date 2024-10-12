package models

import (
	"encoding/json"
	"fmt"

	"github.com/Soulreavertom/stori_challenge/db"
)

type TransactionReport struct {
	Email              string
	TotalBalance       float64
	TotalAverageDebit  float64
	TotalAverageCredit float64
	CreatedAt          string
}

func (tr TransactionReport) Save() (int64, error) {

	query := `
	INSERT INTO requested_reports(email,total_balance,total_average_debit,total_average_credit,created_at)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(tr.Email, tr.TotalBalance, tr.TotalAverageDebit, tr.TotalAverageCredit, tr.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	fmt.Printf("last inserted id:%d", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (tr TransactionReport) GetAll() ([]byte, error) {
	query := `
		SELECT email, total_balance, total_average_debit, total_average_credit, created_at
		FROM requested_reports
	`
	stmt, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer stmt.Close()

	var reports []TransactionReport

	for stmt.Next() {
		var report TransactionReport
		err := stmt.Scan(&report.Email, &report.TotalBalance, &report.TotalAverageDebit, &report.TotalAverageCredit, &report.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		reports = append(reports, report)
	}

	err = stmt.Err()
	if err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	jsonData, err := json.Marshal(reports)
	if err != nil {
		return nil, fmt.Errorf("error marshaling to JSON: %v", err)
	}

	return jsonData, nil
}
