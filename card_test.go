package deck

import "testing"

func TestCard_Is(t *testing.T) {
	tests := []struct {
		name  string
		card  Card
		rank  Rank
		suite Suite
		want  bool
	}{
		{
			name: "is requested card",
			card: Card{
				Suite: Spades,
				Rank:  Two,
			},
			rank:  Two,
			suite: Spades,
			want:  true,
		},
		{
			name: "suite does not match",
			card: Card{
				Suite: Spades,
				Rank:  Two,
			},
			rank:  Two,
			suite: Hearts,
			want:  false,
		},
		{
			name: "rank does not match",
			card: Card{
				Suite: Hearts,
				Rank:  Ace,
			},
			rank:  Two,
			suite: Hearts,
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
				Suite: Spades,
				Rank:  Two,
			},
			want: "2 of 0",
		},
		{
			name: "ace of hearts",
			card: Card{
				Suite: Hearts,
				Rank:  Ace,
			},
			want: "1 of 3",
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
