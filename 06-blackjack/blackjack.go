package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch {
	case card == "ace":
		value = 11
	case card == "two":
		value = 2
	case card == "three":
		value = 3
	case card == "four":
		value = 4
	case card == "five":
		value = 5
	case card == "six":
		value = 6
	case card == "seven":
		value = 7
	case card == "eight":
		value = 8
	case card == "nine":
		value = 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		value = 10
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var Karar string
	switch {
	case card1 == "ace" && card2 == "ace":
		Karar = "P"
	case ParseCard(card1)+ParseCard(card2) == 21:
		if dealerCard == "ten" || dealerCard == "queen" || dealerCard == "king" || dealerCard == "jack" || dealerCard == "ace" {
			Karar = "S"
		} else {
			Karar = "W"
		}
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		Karar = "S"
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			Karar = "H"
		} else {
			Karar = "S"
		}
	case ParseCard(card1)+ParseCard(card2) <= 11:
		Karar = "H"
	}
	return Karar
}
