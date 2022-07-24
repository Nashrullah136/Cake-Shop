package test

import (
	"bytes"
	"encoding/json"
	"github.com/imdario/mergo"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func update(request []byte, id uint) (*httptest.ResponseRecorder, error) {
	req, _ := http.NewRequest("PATCH", "/cakes/"+strconv.FormatUint(uint64(id), 10), bytes.NewBuffer(request))
	response := ExecuteRequest(router, req)
	return response, nil
}

func UpdateSuccessScenario(request []byte, expectedCode int) error {
	log.Print("Inserting Data....")
	defaultCake, err := insertDefaultCake()
	if err != nil {
		return err
	}
	response, err := update(request, defaultCake.ID)
	if err != nil {
		return err
	}
	if err = CheckResponseCode(expectedCode, response.Code); err != nil {
		return err
	}
	var newCake Cake
	if err = json.Unmarshal(request, &newCake); err != nil {
		return err
	}
	if err = mergo.Merge(&newCake, defaultCake); err != nil {
		return err
	}
	err = checkDB(newCake)
	ClearTable("cakes")
	return err
}

func UpdateBadRequestScenario(newCake []byte, expectedCode int) error {
	log.Print("Inserting Data....")
	defaultCake, err := insertDefaultCake()
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("PATCH", "/cakes/"+strconv.FormatUint(uint64(defaultCake.ID), 10), bytes.NewBuffer(newCake))
	response := ExecuteRequest(router, req)
	ClearTable("cakes")
	return CheckResponseCode(expectedCode, response.Code)
}

func TestUpdateWithAllField(t *testing.T) {
	newCake := []byte(`{
		"title":       "Amandine",
		"description": "Chocolate layered cake filled with chocolate, caramel and fondant cream",
		"rating":      10,
		"image":       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/120px-Amandine_cake.jpg"
	}`)
	if err := UpdateSuccessScenario(newCake, http.StatusOK); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithTitleOnly(t *testing.T) {
	newCake := []byte(`{
		"title": "Title"
	}`)
	if err := UpdateSuccessScenario(newCake, http.StatusOK); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithDescriptionOnly(t *testing.T) {
	newCake := []byte(`{
		"description": "Decryption"
	}`)
	if err := UpdateSuccessScenario(newCake, http.StatusOK); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithRatingOnly(t *testing.T) {
	newCake := []byte(`{
		"rating": 11
	}`)
	if err := UpdateSuccessScenario(newCake, http.StatusOK); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithImageOnly(t *testing.T) {
	newCake := []byte(`{
		"image": "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/120px-Amandine_cake.jpg"
	}`)
	if err := UpdateSuccessScenario(newCake, http.StatusOK); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithoutData(t *testing.T) {
	payload := []byte(`{
				"title": "Banana Foster",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	req, _ := http.NewRequest("PATCH", "/cakes/1", bytes.NewBuffer(payload))
	response := ExecuteRequest(router, req)
	err := CheckResponseCode(http.StatusNotFound, response.Code)
	ClearTable("cakes")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithBlankTitle(t *testing.T) {

	payload := []byte(`{
				"title": "",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithNumericTitle(t *testing.T) {

	payload := []byte(`{
				"title": 1,
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithBlankDescription(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "",
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWitNumericDescription(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": 12,
				"rating": 7,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithBlankRating(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": "",
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithStringRating(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": "as",
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithRatingLessThanZero(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": -2,
				"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithBlankImage(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": ""
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateWithNumericImage(t *testing.T) {

	payload := []byte(`{
				"title": "Lemon cheesecake",
				"description": "A cheesecake made of lemon",
				"rating": 7,
				"image": 12
				}`)
	if err := UpdateBadRequestScenario(payload, http.StatusBadRequest); err != nil {
		t.Error(err.Error())
	}
}
