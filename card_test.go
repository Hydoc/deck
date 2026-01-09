package deck

import "testing"

func TestCard_Is(t *testing.T) {
	tests := []struct {
		name  string
		card  Card
		rank  Rank
		suite Suit
		want  bool
	}{
		{
			name: "is requested card",
			card: Card{
				Suit: Spade,
				Rank: Two,
			},
			rank:  Two,
			suite: Spade,
			want:  true,
		},
		{
			name: "suite does not match",
			card: Card{
				Suit: Spade,
				Rank: Two,
			},
			rank:  Two,
			suite: Heart,
			want:  false,
		},
		{
			name: "rank does not match",
			card: Card{
				Suit: Heart,
				Rank: Ace,
			},
			rank:  Two,
			suite: Heart,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.card.Is(tt.rank, tt.suite)
			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want string
	}{
		{
			name: "two of spades",
			card: Card{
				Suit: Spade,
				Rank: Two,
			},
			want: "Two of Spades",
		},
		{
			name: "ace of hearts",
			card: Card{
				Suit: Heart,
				Rank: Ace,
			},
			want: "Ace of Hearts",
		},
		{
			name: "joker",
			card: Card{
				Suit: Joker,
				Rank: Two,
			},
			want: "Joker",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.card.String()
			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}
