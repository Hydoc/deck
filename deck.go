package deck

import "math/rand/v2"

var Suites = []Suit{Spade, Diamond, Club, Heart}
var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func Draw(cards []Card) (Card, []Card) {
	return cards[len(cards)-1], cards[:len(cards)-1]
}

func Shuffle(cards []Card) []Card {
	out := make([]Card, len(cards))
	rand.Shuffle(len(cards), func(i, j int) {
		out[i], out[j] = cards[j], cards[i]
	})
	return out
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, s := range Suites {
		for _, r := range Ranks {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func Jokers(amount int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for range amount {
			cards = append(cards, Card{
				Suit: Joker,
			})
		}
		return cards
	}
}

func Decks(amount int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var out []Card
		for range amount {
			out = append(out, cards...)
		}
		return out
	}
}
