package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten","jack","queen", "king":
        return 10
    case "ace":
        return 11
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerTotal := ParseCard(card1) + ParseCard(card2)
    dealerTotal := ParseCard(dealerCard)
	if playerTotal == 22 {
        return "P"
    }
    if playerTotal == 21 {
        if dealerTotal >= 10 {
            return "S"
        } else {
            return "W"
        }
    }
    if playerTotal >= 17 && playerTotal < 21  {
        return "S"
    }
    if playerTotal >= 12 && playerTotal <= 16 {
        if dealerTotal >= 7 {
            return "H"
        }
        return "S"
    }
    if playerTotal <= 11 {
        return "H"
    }
    return "S"
}
