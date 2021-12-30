package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleRequest(w *httptest.ResponseRecorder, r *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, r)
}

func createTestAlbum() album {
	testAlbum := album{ID: "2", Title: "test", Artist: "Test", Price: 1.00}
	storage.Create(testAlbum)
	return testAlbum
}

func TestAlbumList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not OK!")
	}
}

func TestAlbumDetail(t *testing.T) {
	testAlbum := createTestAlbum()
	request, _ := http.NewRequest("GET", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not ok!", w.Code)
	}
}

func TestAlbumNotFound(t *testing.T) {
	albumId := "9999"
	request, _ := http.NewRequest("GET", "/albums/"+albumId, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestDeleteAlbum(t *testing.T) {
	testAlbum := createTestAlbum()
	request, _ := http.NewRequest("DELETE", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status must be 204")
	}
}

func TestDeleteAlbumNotFound(t *testing.T) {
	albumId := "999"
	request, _ := http.NewRequest("DELETE", "/albums/"+albumId, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	albumId := "9999"
	request, _ := http.NewRequest("PUT", "/albums/"+albumId, strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404", w.Code)
	}
}

func TestUpdateAlbum(t *testing.T) {
	testAlb := createTestAlbum()
	request, _ := http.NewRequest("PUT", "/albums/"+testAlb.ID, strings.NewReader(`{"title": "tEst"}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok", w.Code)
	}
}

func TestCreateAlbumBadStructure(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusBadRequest {
		t.Fatal("status must be 400", w.Code)
	}
}

func TestCreateAlbum(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201", w.Code)
	}
}
