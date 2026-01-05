package card

import "fmt"

type Suite string
type Rank string

const (
	Spades   Suite = "Spades"
	Clubs    Suite = "Clubs"
	Diamonds Suite = "Diamonds"
	Hearts   Suite = "Hearts"
)

const (
	Ace   Rank = "Ace"
	Two   Rank = "Two"
	Three Rank = "Three"
	Four  Rank = "Four"
	Five  Rank = "Five"
	Six   Rank = "Six"
	Seven Rank = "Seven"
	Eight Rank = "Eight"
	Nine  Rank = "Nine"
	Ten   Rank = "Ten"
	Jack  Rank = "Jack"
	Queen Rank = "Queen"
	King  Rank = "King"
)

type Card struct {
	Suite Suite
	Rank  Rank
}

func (c Card) Is(rank Rank, suite Suite) bool {
	return c.Rank == rank && c.Suite == suite
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suite)
}
