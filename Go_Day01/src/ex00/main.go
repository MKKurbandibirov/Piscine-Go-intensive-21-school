package main

import (
	"bufio"
	"day01/domain"

	"day01/jsonDB"
	"day01/xmlDB"
	"flag"
	"io"
	"log"
	"os"
	"strings"
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
	name := flag.String("f", "", "Enter a DB file name")
	flag.Parse()

	var recipes domain.DBReader
	var err error
	var outName string
	if strings.HasSuffix(*name, ".json") {
		recipes, err = JSONHandler(*name)
		if err != nil {
			log.Fatalln(err)
		}
		outName = strings.TrimSuffix(*name, ".json")
		outName += ".xml"
	} else if strings.HasSuffix(*name, ".xml") {
		recipes, err = XMLHandler(*name)
		if err != nil {
			log.Fatalln(err)
		}
		outName = strings.TrimSuffix(*name, ".xml")
		outName += ".json"
	}
	err = Exec(recipes.GetRecipes(), outName)
	if err != nil {
		log.Fatalln(err)
	}
}
