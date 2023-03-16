// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type SearchResult interface {
	IsSearchResult()
}

type Token interface {
	IsToken()
}

type ListingSource string

const (
	ListingSourceOpensea  ListingSource = "OPENSEA"
	ListingSourceSwapmeet ListingSource = "SWAPMEET"
)

var AllListingSource = []ListingSource{
	ListingSourceOpensea,
	ListingSourceSwapmeet,
}

func (e ListingSource) IsValid() bool {
	switch e {
	case ListingSourceOpensea, ListingSourceSwapmeet:
		return true
	}
	return false
}

func (e ListingSource) String() string {
	return string(e)
}

func (e *ListingSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ListingSource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ListingSource", str)
	}
	return nil
}

func (e ListingSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
