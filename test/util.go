package test

import (
	"Pretests/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"time"
)

type Cake struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ExecuteRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	log.Print(recorder.Body)
	return recorder
}

func ClearTable(tableName string) {
	dbConnection := database.GetDbConnection()
	log.Print("Clearing Table....")
	_, err := dbConnection.Exec("DELETE FROM " + tableName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func CheckResponseCode(expected, actual int) error {
	if expected != actual {
		return fmt.Errorf("wrong Response Code. Expected %d got %d", expected, actual)
	}
	return nil
}

func checkBodyResponse(expected, actual Cake) error {
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("Unexpected Body Response. Expected : \n %+v \n Actual : \n %+v \n", expected, actual)
	}
	return nil
}

func setGormDbConnection() *gorm.DB {
	dbConnection, err := gorm.Open(mysql.New(mysql.Config{
		Conn: database.GetDbConnection(),
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return dbConnection
}

func insertCake(newCake Cake) (uint, error) {
	dbConnection := setGormDbConnection()
	result := dbConnection.Create(&newCake)
	if result.Error != nil {
		log.Printf(result.Error.Error())
		return 0, fmt.Errorf("fail to insert data")
	}
	return newCake.ID, nil
}

func getCake(id int) Cake {
	dbConnection := setGormDbConnection()
	var cake Cake
	dbConnection.First(&cake, id)
	return cake
}

func convertResponseToCake(response *httptest.ResponseRecorder) (Cake, error) {
	var cake Cake
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return cake, fmt.Errorf("Failed to read response. Error: " + err.Error())
	}
	err = json.Unmarshal(responseByte, &cake)
	if err != nil {
		return cake, fmt.Errorf("Failed to convert response to construct. Error: " + err.Error())
	}
	return cake, nil
}

func insertDefaultCake() (Cake, error) {
	cake := Cake{
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Title:       "Amandine",
		Description: "Chocolate layered cake filled with chocolate, caramel and fondant cream",
		Rating:      10,
		Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/120px-Amandine_cake.jpg",
	}
	log.Print("Inserting Data....")
	id, err := insertCake(cake)
	if err != nil {
		return cake, err
	}
	cake.ID = id
	return cake, nil
}

func checkDB(expectedCake Cake) error {
	actualCake := getCake(int(expectedCake.ID))
	expectedCake.CreatedAt = actualCake.CreatedAt
	expectedCake.UpdatedAt = actualCake.UpdatedAt
	expectedCake.ID = actualCake.ID
	log.Printf("Expected : \n %+v \n Actual : \n %+v \n", expectedCake, actualCake)
	return checkBodyResponse(expectedCake, actualCake)
}
