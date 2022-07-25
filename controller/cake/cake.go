package cake

import (
	modelCake "Pretests/model/cake"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	enPackage "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func validateCake(cake modelCake.Cake) (map[string]string, error) {
	en := enPackage.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}
	if err := validate.Struct(cake); err != nil {
		return err.(validator.ValidationErrors).Translate(trans), nil
	}
	return nil, nil
}

func GetAll(c *gin.Context) {
	query := c.Request.URL.Query()
	if len(query["page"]) != 0 && len(query["items"]) != 0 {
		page, pageError := strconv.Atoi(query["page"][0])
		items, itemsError := strconv.Atoi(query["items"][0])
		if itemsError == nil && pageError == nil {
			cakes, paginationError := modelCake.FindPagination(page, items)
			if paginationError == nil {
				c.IndentedJSON(http.StatusOK, cakes)
				return
			}
		}
	}
	cakes, err := modelCake.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, cakes)
}

func GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "404 Not Found")
		return
	}
	cake, err := modelCake.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, cake)
}

func Create(c *gin.Context) {
	var cake modelCake.Cake
	if err := c.BindJSON(&cake); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	msg, err := validateCake(cake)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if msg != nil {
		c.IndentedJSON(http.StatusBadRequest, msg)
		return
	}
	newCake, err := modelCake.Insert(cake)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, newCake)
}

func Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "404 Not Found")
		return
	}
	oldCake, err := modelCake.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "404 Not Found")
		return
	}
	var requestMap map[string]interface{}
	jsonRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.Unmarshal(jsonRequest, &requestMap); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err = mergeCakes(&oldCake, requestMap); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	msg, err := validateCake(oldCake)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if msg != nil {
		c.IndentedJSON(http.StatusBadRequest, msg)
		return
	}
	err = modelCake.Update(id, oldCake)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, oldCake)
}

func mergeCakes(dst *modelCake.Cake, src map[string]interface{}) error {
	var jsonToStruct = map[string]string{
		"title":       "Title",
		"description": "Description",
		"rating":      "Rating",
		"image":       "Image",
	}
	dstReflect := reflect.ValueOf(dst).Elem()
	for key, value := range src {
		structName, isExist := jsonToStruct[key]
		if !isExist {
			continue
		}
		dstField := dstReflect.FieldByName(structName)
		if dstField.CanSet() {
			if reflect.ValueOf(value).Kind() == dstField.Kind() {
				dstField.Set(reflect.ValueOf(value))
			} else {
				return fmt.Errorf("%s wrong type", key)
			}
		}
	}
	return nil
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "404 Not Found")
		return
	}
	_, err = modelCake.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "404 Not Found")
		return
	}
	err = modelCake.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted"))
}
