package main

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
	"path"
)

type Area struct {
	Country  string
	Province string
	City     string
}

var (
	db          *ipdb.City
	protectCity = map[string]bool{
		"北京": true,
		"上海": true,
		"广州": true,
		"深圳": true,
	}
)

func Init(basicPath string) (err error) {
	db, err = ipdb.NewCity(path.Join(basicPath, "/conf/ipiptest.ipdb"))
	return
}

func GetLocation(ip string) (area *Area, err error) {
	area = new(Area)
	locations, err := db.Find(ip, "CN")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	length := len(locations)
	switch length {
	case 1:
		area.Country = locations[0]
	case 2:
		area.Country, area.Province = locations[0], locations[1]
	case 3:
		area.Country, area.Province, area.City = locations[0], locations[1], locations[2]
	}
	return
}

func CanShare(ip string) bool {
	area, err := GetLocation(ip)
	if err != nil {
		fmt.Println("error", err)
		return false
	}
	if protectCity[area.Province] {
		return false
	}
	if protectCity[area.City] {
		return false
	}
	return true
}
