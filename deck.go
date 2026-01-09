package deck

import "math/rand/v2"

var Suits = []Suit{Spade, Diamond, Club, Heart}
var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Draw draws the given amount of cards and returns it as the first return value.
// The second return value are the remaining cards in the deck.
func Draw(amount int) func(cards []Card) ([]Card, []Card) {
	return func(cards []Card) ([]Card, []Card) {
		return cards[len(cards)-amount:], cards[:len(cards)-amount]
	}
}

// Shuffle is an option when creating a deck using New.
// It randomizes the deck.
func Shuffle(cards []Card) []Card {
	out := make([]Card, len(cards))
	rand.Shuffle(len(cards), func(i, j int) {
		out[i], out[j] = cards[j], cards[i]
	})
	return out
}

// New creates a new 52 card deck using the default Suits Spade, Diamond, Club and Heart
// Can take options to configure it as desired.
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, s := range Suits {
		for _, r := range Ranks {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// Jokers is an option when creating a deck using New.
// It puts the passed amount of jokers, with the Suit Joker, in the deck without a Rank.
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

// Filter is an option when creating a deck using New.
// It filters the deck using the passed function.
func Filter(f func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var out []Card
		for _, card := range cards {
			if f(card) {
				out = append(out, card)
			}
		}
		return out
	}
}

// Decks is an option when creating a deck using New.
// It creates the passed amount of decks, so 52 * amount.
func Decks(amount int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var out []Card
		for range amount {
			out = append(out, cards...)
		}
		return out
	}
}
