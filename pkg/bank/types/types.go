package types

type PAN string
type Currency string
type Money int64

// cost
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Payment
type Payment struct {
	ID     int
	Amount Money
}

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
