package main

import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	log "github.com/Sirupsen/logrus"
	"math/rand"
)


// Find_item generates a new item for the hero when they level up
// and notifies the player about the item found
//Parameters: Hero ID, Hero Level, Hero Name and the SQL.DB connection.
//Return: None
func Find_item(heroID int64, hero_level int, hero_name string, conn *sql.DB){

	//log.Info("Find Items called, Hero Level: " + strconv.Itoa(hero_level))

	items :=[10]string{"weapon","tunic","shield","leggins","ring","gloves","boots","helm","charm","amulet"}

	find_chance := [51]float32{ 100.00, 91.93227152,84.51542547,77.69695042,71.42857143,65.66590823,60.36816105,55.49782173,
		51.02040816,46.90422016,43.12011504,39.64130124,36.44314869,33.5030144,	30.80008217,28.31521517,26.03082049,
		23.93072457,22.00005869,20.22515369,18.59344321,17.0933747,	15.71432764,14.44653835,13.28103086,12.20955335,
		11.22451974,10.31895597,9.486450616,8.721109539,8.017514101,7.370682832,6.776036155,6.229363956,5.726795786,
		5.264773452,4.840025825,4.449545683,4.090568419,3.760552466,3.457161303,3.178246916,2.921834585,2.686108904,
		2.469400931,2.270176369,2.087024703,1.918649217,1.763857808,1.621554549,1.490731931	}

	var item_type int
	var item_name string
	var item_level int
	var item_found_chance float32

	for i := hero_level; i > 0; i-- {

		if(i > 50) {
			//After Hero Level of 50, has a 1% chance to find an item.
			item_found_chance = 1.0

		} else {

			item_found_chance = find_chance[i]
		}
		//Start with highest Level Item and subtract a level as it misses the chance
		if(rand.Intn(100) <= int(item_found_chance)){

			item_gain_percentage := float64(rand.Intn(100))
			item_level = int(float64(i) + (float64(i) * (item_gain_percentage/100)))
			item_type = rand.Intn(10)
			item_name = items[item_type]

			var log_msg = "Item Found: " + item_name + " | Hero Level: " + strconv.Itoa(hero_level)+ " found @ Level: " + strconv.Itoa(i)+ " Item Level: " + strconv.Itoa(item_level)

			log.Info(log_msg)

			break
		}
	}

	//check if No items where found

	var current_item_level = Get_Item_By_HeroID(heroID, item_name, conn)

	log.Info("Items: Current: " + strconv.Itoa(current_item_level) + " | New: " + strconv.Itoa(item_level))

	var message string
	var message_plural string

	if (item_name == items[3] || item_name == items[5] || item_name == items[6]) {

		message_plural =  " are only level "

	} else {

		message_plural =  " is only level "
	}

	if item_level > current_item_level {

		//Replace the current item value with the new one
		Update_Item_for_Hero(heroID,item_name, item_level, conn)

		message = "You found a level " + strconv.Itoa(item_level) + " " + item_name + "! Your current " + item_name + message_plural +  strconv.Itoa(current_item_level) + ", so it seems Luck is with you!"

	} else {

		//Message back to player that current item level is better
		message = "You found a level " + strconv.Itoa(item_level) + " " + item_name + "! Your current " + item_name + message_plural +  strconv.Itoa(current_item_level) + ", so it seems Luck is against you. You toss the " + item_name + "."

	}

	log.Info(message)

	//Message back to player that new item is better
	Insert_World_Event_for_Hero(heroID, message, conn)


}