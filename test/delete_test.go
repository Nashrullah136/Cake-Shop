package test

import (
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestDeleteNotFound(t *testing.T) {
	id := rand.Intn(20) + 100
	req, _ := http.NewRequest("DELETE", "/cakes/"+strconv.Itoa(id), nil)
	response := ExecuteRequest(router, req)
	err := CheckResponseCode(http.StatusNotFound, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDeleteWithFoundData(t *testing.T) {
	cake, err := insertDefaultCake()
	if err != nil {
		t.Error(err.Error())
	}
	req, _ := http.NewRequest("DELETE", "/cakes/"+strconv.Itoa(int(cake.ID)), nil)
	response := ExecuteRequest(router, req)
	err = CheckResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
}
