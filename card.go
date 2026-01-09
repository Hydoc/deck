//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit int
type Rank int

const (
	Spade Suit = iota
	Club
	Diamond
	Heart
	Joker
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit Suit
	Rank Rank
}

// Is checks if a card is a specific card by checking Rank and Suit.
func (c Card) Is(rank Rank, suit Suit) bool {
	return c.Rank == rank && c.Suit == suit
}

func (c Card) String() string {
	if c.Suit != Joker {
		return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
	}
	return c.Suit.String()
}
