package card

import "testing"

func Test_NewDeck(t *testing.T) {

	tests := []struct {
		name          string
		wantSize      int
		amountOfDecks int
	}{
		{
			name:          "create",
			wantSize:      52,
			amountOfDecks: 1,
		},
		{
			name:          "create two decks",
			wantSize:      104,
			amountOfDecks: 2,
		},
		{
			name:          "empty deck",
			wantSize:      0,
			amountOfDecks: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeck(false, tt.amountOfDecks)

			if tt.wantSize != len(d.Cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(d.Cards))
			}
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name      string
		wantSize  int
		setup     func(d *Deck) Card
		wantSuite Suite
		wantRank  Rank
	}{
		{
			name:     "draw a card",
			wantSize: 51,
			setup: func(d *Deck) Card {
				return d.Draw()
			},
			wantSuite: Hearts,
			wantRank:  Ace,
		},
		{
			name:     "draw two cards",
			wantSize: 50,
			setup: func(d *Deck) Card {
				d.Draw()
				return d.Draw()
			},
			wantSuite: Hearts,
			wantRank:  King,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeck(false, 1)
			card := tt.setup(d)

			if tt.wantSize != len(d.Cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(d.Cards))
			}

			if tt.wantSuite != card.Suite {
				t.Errorf("want suite %s, got %s", tt.wantSuite, card.Suite)
			}

			if tt.wantRank != card.Rank {
				t.Errorf("want rank %s, got %s", tt.wantRank, card.Rank)
			}
		})
	}
}
