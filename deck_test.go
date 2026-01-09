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

func TestDeck_Jokers(t *testing.T) {
	tests := []struct {
		name           string
		wantSize       int
		amountOfJokers int
	}{
		{
			name:           "with one joker",
			wantSize:       53,
			amountOfJokers: 1,
		},
		{
			name:           "with two jokers",
			wantSize:       54,
			amountOfJokers: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := New(Jokers(tt.amountOfJokers))

			if tt.wantSize != len(cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(cards))
			}
		})
	}
}

func TestDeck_Filter(t *testing.T) {
	tests := []struct {
		name     string
		wantSize int
		filter   func(Card) bool
	}{
		{
			name:     "filter one suit",
			wantSize: 13,
			filter: func(card Card) bool {
				return card.Suit == Heart
			},
		},
		{
			name:     "filter only aces",
			wantSize: 4,
			filter: func(card Card) bool {
				return card.Rank == Ace
			},
		},
		{
			name:     "get a specific card",
			wantSize: 1,
			filter: func(card Card) bool {
				return card.Rank == Ten && card.Suit == Spade
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := New(Filter(tt.filter))

			if tt.wantSize != len(cards) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(cards))
			}
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name          string
		wantSize      int
		amountOfDraws int
		wantSuit      Suit
		wantRank      Rank
	}{
		{
			name:          "draw a card",
			wantSize:      51,
			amountOfDraws: 1,
			wantSuit:      Heart,
			wantRank:      King,
		},
		{
			name:          "draw five cards",
			wantSize:      47,
			amountOfDraws: 5,
			wantSuit:      Heart,
			wantRank:      Queen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawnCards, remaining := Draw(tt.amountOfDraws)(New(Shuffle))

			if tt.wantSize != len(remaining) {
				t.Errorf("want size %d, got %d", tt.wantSize, len(remaining))
			}

			if len(drawnCards) != tt.amountOfDraws {
				t.Errorf("want drawn cards %d, got %d", tt.amountOfDraws, len(drawnCards))
			}
		})
	}
}
