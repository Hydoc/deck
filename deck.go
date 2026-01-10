package deck

import (
	"math/rand/v2"
	"sort"
)

var Suits = []Suit{Spade, Diamond, Club, Heart}
var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Draw removes the given amount of cards from the end and returns it as the first return value.
// The second return value are the remaining cards in the deck.
func Draw(amount int) func(cards []Card) ([]Card, []Card) {
	return func(cards []Card) ([]Card, []Card) {
		return cards[len(cards)-amount:], cards[:len(cards)-amount]
	}
}

// Shuffle randomizes the deck.
// Can be used as an option when creating a deck using New.
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

// Jokers puts the passed amount of jokers, with the Suit Joker, in the deck without a Rank.
// Can be used as an option when creating a deck using New.
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

// Sort takes a less function and applies it to the cards.
// Can be used as an option when creating a deck using New.
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// DefaultSort sorts the cards using the Less function
// Can be used as an option when creating a deck using New.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Less sorts the cards ascending by their corresponding rank
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].rank() < cards[j].rank()
	}
}

// Filter applies the given filter to the deck and returns it.
// Can be used as an option when creating a deck using New.
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

// Decks puts the given amount of decks in a deck (52 * amount).
// Can be used as an option when creating a deck using New.
func Decks(amount int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var out []Card
		for range amount {
			out = append(out, cards...)
		}
		return out
	}
}
