package main

import (
	"fmt"
	"log"
	"time"

	"github.com/navneetshukl/bolt/db"
)

func init() {
	db.ConnectToBoltDB()
}

func main() {

	start := time.Now()

	task := []string{"Navneet", "Yatinjal", "Rohan", "Rohit", "Dhoni"}
	err := db.InsertToBoltDB("Watch Youtube")
	if err != nil {
		log.Println("Error in inserting to BOLTDB ", err)
		return
	}

	data, err := db.GetFromBoltDB("Watch Youtube")
	if err != nil {
		log.Println("Error in getting the data from BOLTDB ", err)
		return
	}

	err = db.ArrayInsertToBoltDB(task)
	if err != nil {
		log.Println("Error in inserting to BOLTDB ", err)
		return
	}

	datas, err := db.ArrayGetFromBoltDB(task)
	if err != nil {
		log.Println("Error in getting the data from BOLTDB ", err)
		return
	}

	end := time.Since(start)
	fmt.Println("Total time taken is ", end)
	fmt.Println(data)

	fmt.Println("***********************************************************************************")

	fmt.Println(datas)

}
