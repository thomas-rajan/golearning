package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type DnpPointMap struct {
	XMLName   xml.Name   `xml:"dnpPointMap"`
	Version   string     `xml:"version,attr"`
	DnpPoints []DnpPoint `xml:"dnpPoint"`
}

type DnpPoint struct {
	XMLName xml.Name `xml:"dnpPoint"`
	Type    string   `xml:"type,attr"`
	Index   string   `xml:"index,attr"`
	Name    string   `xml:"name,attr"`
}

func main() {
	byteValue, err := os.ReadFile("cmd/dnp.xml")

	// we initialize our Users array
	var dnpPointMap DnpPointMap
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &dnpPointMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Dnp Point map version", dnpPointMap.Version)
	for _, dnpPoint := range dnpPointMap.DnpPoints {
		fmt.Printf("type %v, index: %v, name: %v \n", dnpPoint.Type, dnpPoint.Index, dnpPoint.Name)
	}
}
