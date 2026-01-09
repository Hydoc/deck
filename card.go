package deck

import "fmt"

type Suite int
type Rank int

const (
	Spades Suite = iota
	Clubs
	Diamonds
	Hearts
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
	Suite Suite
	Rank  Rank
}

func (c Card) Is(rank Rank, suite Suite) bool {
	return c.Rank == rank && c.Suite == suite
}

func (c Card) String() string {
	return fmt.Sprintf("%d of %d", c.Rank, c.Suite)
}
