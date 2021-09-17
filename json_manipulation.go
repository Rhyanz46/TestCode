package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

//Your tasks to create functions:
//1. Find items in the Meeting Room.
//2. Find all electronic devices.
//3. Find all the furniture.
//4. Find all items were purchased on 16 Januari 2020.
//5. Find all items with brown color.

type Filter struct {
	RoomId int
	Name   string
	Type   string
	Date   time.Time
}

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

func Find(data []Data, filter Filter) (result []Data) {
	for _, value := range data {
		fmt.Println(value)
	}
	for _, value := range data {
		if filter.Name != "" && !strings.Contains(strings.ToLower(value.Name), strings.ToLower(filter.Name)) {
			continue
		}
		if filter.Type != "" && !strings.Contains(strings.ToLower(value.Type), strings.ToLower(filter.Type)) {
			continue
		}
		if filter.Date != (Filter{}).Date {
			var DateFormat = "2 January 2006"
			current := time.Unix(int64(value.PurchasedAt), 0).Format(DateFormat)
			filterDate := filter.Date.Format(DateFormat)
			if current != filterDate {
				continue
			}
		}
		result = append(result, value)
	}
	return result
}

func main() {
	plan, _ := ioutil.ReadFile("data.json")
	var data []Data
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error")
	}

	//1. Find items in the Meeting Room.
	for _, value := range Find(data, Filter{
		RoomId: 1,
	}) {
		unixTimeUTC := time.Unix(int64(value.PurchasedAt), 0).Format(time.RFC3339)
		fmt.Println(value.Name, " - ", unixTimeUTC)
	}
}
