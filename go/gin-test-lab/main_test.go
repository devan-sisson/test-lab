package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-test-lab/models"
)

type mockBookModel struct{}

func (m *mockBookModel) All() ([]models.Book, error) {
	var bks []models.Book

	bks = append(bks, models.Book{"978-1503261969", "Emma", "Jayne Austen", 9.44})
	bks = append(bks, models.Book{"978-1505255607", "The Time Machine", "H. G. Wells", 5.99})

	return bks, nil
}

func TestBooksIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	env := Env{books: &mockBookModel{}}

	http.HandlerFunc(env.books.All()).ServeHTTP(rec, req)

	expected := "978-1503261969, Emma, Jayne Austen, £9.44\n978-1505255607, The Time Machine, H. G. Wells, £5.99\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}