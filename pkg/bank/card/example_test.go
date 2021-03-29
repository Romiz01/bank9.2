/*
package card

import (
	"bank/pkg/bank/types"

	"fmt"
)

func ExamplePaymentSource(cards []types.Card) []types.PaymentSource {
	card := types.Card{Balance: 10_000_00, Active: true}
	a := PaymentSources(card, "card", '5058 0000 0000 8888', 10_000_00)
	fmt.Fprintln(a)

}
*/

package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:     "2",
		}, {
			Balance: 2_000_00,
			Active:  true,
			PAN:     "3",
		}, {
			Balance: 3_000_00,
			Active:  false,
			PAN:     "4",
		},
	}
	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

	// Output:
	//2
	//3
}
