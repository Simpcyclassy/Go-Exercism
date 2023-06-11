package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "ace":
    	return 11
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
    case "ten", "jack", "queen", "king":
    	return 10
    default:
    	return 0
    }
}


// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parseCardMap := map[string]int{
		"ace": 11,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"ten": 10,
		"jack": 10,
		"queen": 10,
		"king": 10,
	}
    summedParsedCard := parseCardMap[card1] + parseCardMap[card2]
    parsedDealerCard := parseCardMap[dealerCard]
    if card1 == "ace" && card2 == "ace" {
        return "P"
    }
    if summedParsedCard == 21 {
        if parsedDealerCard == 11 || parsedDealerCard == 10 {
            return "S"
        }
        return "W"
    }
    if summedParsedCard >= 17 && summedParsedCard <= 20 {
        return "S"
    }
    if summedParsedCard >= 12 && summedParsedCard <= 16 {
        if parsedDealerCard >= 7 {
            return "H"
        }
        return "S"
    }
    if summedParsedCard <= 11 {
        return "H"
    }
    return ""
}