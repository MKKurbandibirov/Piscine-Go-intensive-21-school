package main

import (
	"bufio"
	"day01/compare"
	"day01/domain"

	"day01/jsonDB"
	"day01/xmlDB"
	"flag"
	"io"
	"log"
	"os"
)

func Exec(r domain.DBReader, name string) error {
	data, err := r.GetResult()
	if err != nil {
		return err
	}
	outFile, err := os.Create(name)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(data)
	writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func JSONHandler(name string) (*jsonDB.JSONRecipes, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var recipes = new(jsonDB.JSONRecipes)
	err = recipes.ConvertFile(data)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func XMLHandler(name string) (*xmlDB.XMLRecipes, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var recipes = new(xmlDB.XMLRecipes)
	err = recipes.ConvertFile(data)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func main() {
	oldName := flag.String("old", "", "Enter a old DB file name")
	newName := flag.String("new", "", "Enter a new DB file name")
	flag.Parse()

	var oldRecipes *xmlDB.XMLRecipes
	var newResipes *jsonDB.JSONRecipes
	var err error
	oldRecipes, err = XMLHandler(*oldName)
	if err != nil {
		log.Fatalln(err)
	}
	newResipes, err = JSONHandler(*newName)
	if err != nil {
		log.Fatalln(err)
	}
	compare.Compare(oldRecipes.GetRecipes(), newResipes.GetRecipes())
}
