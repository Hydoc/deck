package card

import "math/rand/v2"

type Deck struct {
	Cards []Card
}

func (d *Deck) Draw() Card {
	card, cards := d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	d.Cards = cards
	return card
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func NewDeck(shuffle bool, amount int) *Deck {
	suites := []Suite{Spades, Clubs, Diamonds, Hearts}
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	cards := make([]Card, 0)

	for i := 0; i < amount; i++ {
		for _, s := range suites {
			for _, r := range ranks {
				cards = append(cards, Card{Suite: s, Rank: r})
			}
		}

	}

	deck := &Deck{
		Cards: cards,
	}

	if shuffle {
		deck.Shuffle()
	}

	return deck
}
