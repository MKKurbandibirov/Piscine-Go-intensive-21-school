package xmlDB

import (
	"encoding/json"
	"encoding/xml"
)

type XMLIngredient struct {
	XMLName xml.Name `xml:"item" json:"-"`
	Name    string   `xml:"itemname"`
	Count   string   `xml:"itemcount"`
	Unit    string   `xml:"itemunit"`
}

type XMLIngredients struct {
	XMLName     xml.Name        `xml:"ingredients" json:"-"`
	Ingredients []XMLIngredient `xml:"item"`
}

type XMLCake struct {
	XMLName     xml.Name       `xml:"cake" json:"-"`
	Name        string         `xml:"name"`
	Time        string         `xml:"stovetime"`
	Ingredients XMLIngredients `xml:"ingredients"`
}

type XMLRecipes struct {
	XMLName xml.Name  `xml:"recipes" json:"-"`
	Cakes   []XMLCake `xml:"cake"`
}

var recipes XMLRecipes

func (x *XMLRecipes) GetRecipes() *XMLRecipes {
	return &recipes
}

func (x *XMLRecipes) ConvertFile(data []byte) error {
	err := xml.Unmarshal(data, &recipes)
	if err != nil {
		return err
	}
	return nil
}

func (x *XMLRecipes) GetResult() ([]byte, error) {
	data, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
