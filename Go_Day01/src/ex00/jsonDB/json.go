package jsonDB

import (
	"day01/domain"
	"encoding/json"
	"encoding/xml"
)

type JSONIngredient struct {
	Name  string `json:"ingredient_name"`
	Count string `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit"`
}

type JSONCake struct {
	Name        string           `json:"name"`
	Time        string           `json:"time"`
	Ingredients []JSONIngredient `json:"ingredients"`
}

type JSONRecipes struct {
	Cakes []JSONCake `json:"cake"`
}

var recipes JSONRecipes

func (j *JSONRecipes) GetRecipes() domain.DBReader {
	return &recipes
}

func (j *JSONRecipes) ConvertFile(data []byte) error {
	err := json.Unmarshal(data, &recipes)
	if err != nil {
		return err
	}
	return nil
}

func (j *JSONRecipes) GetResult() ([]byte, error) {
	data, err := xml.MarshalIndent(j, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
