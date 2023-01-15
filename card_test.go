package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Spade, Ace})

	// Output:
	// Ace of Spades
}

func TestNew(t* testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t* testing.T) {
	cards := New(DefaultSort)
	card := Card{Spade, Ace}

	if cards[0] != card {
		t.Error("Expected Ace of Spades as first card.")
	}
}

func TestSort(t* testing.T) {
	cards := New(Sort(Less))
	card := Card{Spade, Ace}

	if cards[0] != card {
		t.Error("Expected Ace of Spades as first card.")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, Received:", count)
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(c Card) bool {
		return c.Rank == Ace
	}))

	for _, card := range cards {
		if card.Rank == Ace {
			t.Error("Expected no Aces, Received:", card)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	if len(cards) != 156 {
		t.Error("Expected 156 cards, Received:", len(cards))
	}
}