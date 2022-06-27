package main

import "fmt"

type Item struct {
	title, body string
}

var database []Item

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func GetByName(title string) Item {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteItem(item Item) Item {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {
	fmt.Println("initial database: ", database)

	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("second database: ", database)
}
