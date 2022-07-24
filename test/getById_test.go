package test

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestGetByIdNotFound(t *testing.T) {
	id := rand.Intn(20) + 100
	req, _ := http.NewRequest("GET", "/cakes/"+strconv.Itoa(id), nil)
	response := ExecuteRequest(router, req)
	err := CheckResponseCode(http.StatusNotFound, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetByIdFound(t *testing.T) {
	log.Print("Inserting Data....")
	cake, err := insertDefaultCake()
	if err != nil {
		t.Errorf(err.Error())
	}
	req, _ := http.NewRequest("GET", "/cakes/"+strconv.FormatUint(uint64(cake.ID), 10), nil)
	response := ExecuteRequest(router, req)
	err = CheckResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
	result, err := convertResponseToCake(response)
	if err != nil {
		t.Error(err.Error())
	}
	err = checkBodyResponse(cake, result)
	if err != nil {
		t.Error(err.Error())
	}
	ClearTable("cakes")
}
