package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func insertTestBadRequestScenario(request []byte, expectedCode int) error {
	req, _ := http.NewRequest("POST", "/cakes", bytes.NewBuffer(request))
	response := ExecuteRequest(router, req)
	ClearTable("cakes")
	return CheckResponseCode(expectedCode, response.Code)
}

func TestInsertWithRightData(t *testing.T) {
	cake := Cake{
		Title:       "Amandine",
		Description: "Chocolate layered cake filled with chocolate, caramel and fondant cream",
		Rating:      10,
		Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/120px-Amandine_cake.jpg",
	}
	payload, err := json.Marshal(cake)
	if err != nil {
		t.Error(err.Error())
	}
	req, _ := http.NewRequest("POST", "/cakes", bytes.NewBuffer(payload))
	response := ExecuteRequest(router, req)
	err = CheckResponseCode(http.StatusCreated, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
	result, err := convertResponseToCake(response)
	if err != nil {
		t.Error(err.Error())
	}
	err = checkDB(result)
	if err != nil {
		t.Error(err.Error())
	}
	ClearTable("cakes")
}

func TestInsertWithBlankTitle(t *testing.T) {
	payload := []byte(`{
				"title": "",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithNullTitle(t *testing.T) {
	payload := []byte(`{
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithNumericTitle(t *testing.T) {
	payload := []byte(`{
				"title": 1,
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithBlankDescription(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWitNullDescription(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWitNumericDescription(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithBlankRating(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"rating": "",
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithNullRating(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithStringRating(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"rating": "as",
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithRatingLessThanZero(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"rating": -2,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithBlankImage(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": ""
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithNullImage(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": 7
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestInsertWithNumericImage(t *testing.T) {
	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": 12
				}`)
	if err := insertTestBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}
