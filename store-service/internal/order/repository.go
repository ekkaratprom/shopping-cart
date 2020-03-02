package order

import "github.com/jmoiron/sqlx"

type OrderRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (orderRepository OrderRepositoryMySQL) CreateShipping(orderID int, submittedOrder SubmitedOrder) (int, error) {
	return 0, nil
}

func (orderRepository OrderRepositoryMySQL) CreateOrder(totalPrice float64) (int, error) {
	tx := orderRepository.DBConnection.MustBegin()
	sqlResult := tx.MustExec("INSERT INTO orders (total_price) VALUE (?)", totalPrice)
	insertedId, err := sqlResult.LastInsertId()
	return int(insertedId), err
}
