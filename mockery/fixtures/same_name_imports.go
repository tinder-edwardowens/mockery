package test

import (
	"net/http"

	my_http "github.com/tinder-edwardowens/mockery/mockery/fixtures/http"
)

// Example is an example
type Example interface {
	A() http.Flusher
	B(fixtureshttp string) my_http.MyStruct
}
