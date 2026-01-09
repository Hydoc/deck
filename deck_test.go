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
			cards := New(Decks(tt.amountOfDecks))

			if tt.wantSize != len(cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(cards))
			}
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name     string
		wantSize int
		setup    func(cards []Card) (Card, []Card)
		wantSuit Suit
		wantRank Rank
	}{
		{
			name:     "draw a card",
			wantSize: 51,
			setup: func(cards []Card) (Card, []Card) {
				return Draw(cards)
			},
			wantSuit: Heart,
			wantRank: King,
		},
		{
			name:     "draw two cards",
			wantSize: 50,
			setup: func(cards []Card) (Card, []Card) {
				_, r := Draw(cards)
				return Draw(r)
			},
			wantSuit: Heart,
			wantRank: Queen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := New()
			card, remaining := tt.setup(cards)

			if tt.wantSize != len(remaining) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(remaining))
			}

			if tt.wantSuit != card.Suit {
				t.Errorf("want suit %d, got %d", tt.wantSuit, card.Suit)
			}

			if tt.wantRank != card.Rank {
				t.Errorf("want rank %d, got %d", tt.wantRank, card.Rank)
			}
		})
	}
}
