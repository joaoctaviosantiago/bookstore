package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "For the love of Go",
		Author: "John Arundel",
		Copies: 99,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	want := 1
	result, err := bookstore.Buy(b)
	got := result.Copies

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of %d, got %d", want, b.Copies, got)
	}
}

func TestBuyErrorsIfNoBookLeft(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	got := catalog.GetAllBooks()

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := bookstore.Book{
		ID: 2, Title: "The Power of Go: Tools",
	}

	got, err := catalog.GetBook(2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookInvalidId(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{}

	_, err := catalog.GetBook(999)

	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	want := 3000
	got := b.NetPriceCents()

	if want != got {
		t.Errorf("NetPrice(%d, %d): want %d, got %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}

	want := 6000
	err := b.SetPriceCents(want)
	got := b.PriceCents

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsNegativePrice(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{}
	err := b.SetPriceCents(-1)

	if err == nil {
		t.Errorf("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "For the Love of Go",
	}

	want := "Autobiography"
	err := b.SetCategory(want)
	got := b.Category()

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want category %q, got %q", want, got)
	}

}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "For the Love of Go",
	}

	err := b.SetCategory("bogus")

	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}

}
