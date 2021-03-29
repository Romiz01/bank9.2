package card

import (
	"bank/pkg/bank/types"
)

/*
// isue card
func IssueCard(currency types.Currency,
	color string, name string) types.Card {

	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
*/

// PaymentSources

// func PaymentSources(cards []types.Card) []types.PaymentSource {
// 	var operations []types.PaymentSource
// 	i := 0
// 	for _, card := range cards {
// 		if card.Balance <= 0 && !card.Active {
// 			continue
// 		}
// 		i++
// 	}
// 	return operations

// }

// PaymentSources
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payments []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			payment := types.PaymentSource{
				Type:    "card",
				Balance: card.Balance,
				Number:  string(card.PAN),
			}
			payments = append(payments, payment)
		}
	}
	return payments
}
