package repository

type Transactionser interface {
	NewTransactions()
	TransactionsRollback()
	TransactionsCommit()
}
