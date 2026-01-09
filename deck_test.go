package deck

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
			cards := New(Deck(tt.amountOfDecks))

			if tt.wantSize != len(cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(cards))
			}
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name      string
		wantSize  int
		setup     func(cards []Card) (Card, []Card)
		wantSuite Suite
		wantRank  Rank
	}{
		{
			name:     "draw a card",
			wantSize: 51,
			setup: func(cards []Card) (Card, []Card) {
				return Draw(cards)
			},
			wantSuite: Hearts,
			wantRank:  King,
		},
		{
			name:     "draw two cards",
			wantSize: 50,
			setup: func(cards []Card) (Card, []Card) {
				_, r := Draw(cards)
				return Draw(r)
			},
			wantSuite: Hearts,
			wantRank:  Queen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := New()
			card, remaining := tt.setup(cards)

			if tt.wantSize != len(remaining) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(remaining))
			}

			if tt.wantSuite != card.Suite {
				t.Errorf("want suite %d, got %d", tt.wantSuite, card.Suite)
			}

			if tt.wantRank != card.Rank {
				t.Errorf("want rank %d, got %d", tt.wantRank, card.Rank)
			}
		})
	}
}
