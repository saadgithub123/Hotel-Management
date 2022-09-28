package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Restaurent struct {
	EaterID    string `json:"eater_id"`
	FoodMenuID string `json:"foodmenu_id"`
	FoodName   string `json:"foodname"`
}

type KeyValueMap struct {
	Key   string
	Value int
}

func HotelManagement() error {
	// Open our jsonFile
	jsonFile, err := os.Open("log.json")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened log.json")
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var eater_id []Restaurent
	var foodmenu_id []Restaurent

	// we unmarshal our byteArray which contains
	json.Unmarshal(byteValue, &foodmenu_id)
	json.Unmarshal(byteValue, &eater_id)

	// create a map to store the count of each foodmenu_id and eater_id
	eater_idMap := make(map[string]int)
	foodmenu_idMap := make(map[string]int)

	for i := 0; i < len(foodmenu_id); i++ {
		foodmenu_idMap[foodmenu_id[i].FoodMenuID]++
	}

	for i := 0; i < len(eater_id); i++ {
		eater_idMap[eater_id[i].EaterID]++
	}
	sortedMap := sortMap(foodmenu_idMap)

	// iterate through the foodmenuMap and print the foodmenu_id and its count
	fmt.Println("The top 3 menu items consumed are ")

	for i, kv := range sortedMap {
		if i == 3 {
			break
		}
		fmt.Println("FoodMenuIDs: " + kv.Key + " Count: " + fmt.Sprint(kv.Value))

	}
	var idArray []string
	for _, tempData := range foodmenu_id {
		idArray = append(idArray, tempData.EaterID)
	}
	unique(foodmenu_id)

	return nil
}

// sorting foodmenu_id
func sortMap(foodmenu_idMap map[string]int) []KeyValueMap {

	var ss []KeyValueMap
	for k, v := range foodmenu_idMap {
		ss = append(ss, KeyValueMap{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

// checking if same eater_id has same foodmenu_id more than once
func unique(s []Restaurent) []Restaurent {
	inResult := make(map[Restaurent]bool)
	var result []Restaurent
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		} else {
			fmt.Println("Error - EaterID: " + str.EaterID + " has taken " + str.FoodName + " more than once")
		}
	}
	return result
}

func main() {
	err := HotelManagement()
	if err != nil {
		fmt.Println(err)
	}

}
