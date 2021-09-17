package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

//Your tasks to create functions:
//1. Find items in the Meeting Room.
//2. Find all electronic devices.
//3. Find all the furniture.
//4. Find all items were purchased on 16 Januari 2020.
//5. Find all items with brown color.

//filter format
//map[string]interface{}{
//	"room_id" : 1,
//	"name": "brown",
//	"type": "tableware",
//	"date": "16 January 2020",
//	"tags": "brown",
//}

type Data struct {
	InventoryId int      `json:"inventory_id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	PurchasedAt int      `json:"purchased_at"`
	Placement   struct {
		RoomId int    `json:"room_id"`
		Name   string `json:"name"`
	} `json:"placement"`
}

var DateFormat = "2 January 2006"

func Find(data []Data, filter map[string]interface{}) (result []Data) {
	for _, value := range data {
		// find by room_id
		if filter["room_id"] != "" && filter["room_id"] != nil {
			if reflect.TypeOf(filter["room_id"]).String() != "int" {
				panic("room_id must be integer")
			}
			if filter["room_id"].(int) != value.Placement.RoomId {
				continue
			}
		}

		// find by name
		if filter["name"] != "" && filter["name"] != nil {
			if reflect.TypeOf(filter["name"]).String() != "string" {
				panic("name must be integer")
			}
			if !strings.Contains(strings.ToLower(value.Name), strings.ToLower(filter["name"].(string))) {
				continue
			}
		}

		// find by type
		if filter["type"] != "" && filter["type"] != nil {
			if reflect.TypeOf(filter["type"]).String() != "string" {
				panic("type must be integer")
			}
			if !strings.Contains(strings.ToLower(value.Type), strings.ToLower(filter["type"].(string))) {
				continue
			}
		}

		// find by date
		if filter["date"] != "" && filter["date"] != nil {
			time1, err := time.Parse(DateFormat, filter["date"].(string))
			if err != nil {
				panic("date format is not valid")
			}
			current := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
			filterDate := time1.Format(DateFormat)
			if current != filterDate {
				continue
			}
		}

		// find by tags
		if filter["tags"] != "" && filter["tags"] != nil {
			var found bool
			if reflect.TypeOf(filter["tags"]).String() != "string" {
				panic("type must be integer")
			}
			filterTag := strings.ToLower(filter["tags"].(string))
			for _, tag := range value.Tags {
				if filterTag == strings.ToLower(tag) {
					found = true
				}
			}
			if !found {
				continue
			}
		}
		result = append(result, value)
	}
	return result
}

func main() {
	plan, _ := ioutil.ReadFile("data.json")
	var data, result []Data
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("1. Find items in the Meeting Room with room_id : 3")
	result = Find(data, map[string]interface{}{"room_id": 3})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

	fmt.Println("\n2. Find all electronic devices")
	result = Find(data, map[string]interface{}{"type": "electronic"})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

	fmt.Println("\n3. Find all the furniture")
	result = Find(data, map[string]interface{}{"type": "furniture"})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

	fmt.Println("\n4. Find all items were purchased on 16 Januari 2020")
	result = Find(data, map[string]interface{}{"date": "16 January 2020"})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

	fmt.Println("\n4. Find all items were purchased on 16 Januari 2020")
	result = Find(data, map[string]interface{}{"date": "16 January 2020"})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

	fmt.Println("\n5. Find all items with brown color")
	result = Find(data, map[string]interface{}{"tags": "brown"})
	for _, value := range result {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
		fmt.Println(fmt.Sprintf("Room id : %d, PurchasedAt : %v, Item name : %s ", value.Placement.RoomId, unixTimeUTC, value.Name))
	}

}
