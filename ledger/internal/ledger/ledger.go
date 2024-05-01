package ledger

import "database/sql"

func Insert(db *sql.DB, orderId string, userID string, amount int64, operation string, transactionDate string) error {
	stmt, err := db.Prepare("INSERT INTO ledger (order_id, user_id, amount, operation, transaction_date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(orderId, userID, amount, operation, transactionDate)
	if err != nil {
		return nil
	}

	return nil
}