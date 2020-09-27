package types

// представляет собой денежную сумму в минимальных единицах (центы, дирамы, копейки и т.д)
type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

// Category представляет собой категорию, в которой был совершен платёж (авто, аптеки, рестораны и т.д)
type Category string

// Status представляет с собой статус платежа.
type Status string

// Переопределенные статусы платежей.
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment представляет с собой информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
