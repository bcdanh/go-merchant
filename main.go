package main

import (
	"log"
	"merchant/hunter"
	"merchant/librarian"
	"merchant/transmitter"
	"os"
	"sync"

	"github.com/robfig/cron"
)

var barie sync.WaitGroup
var cronbarie sync.WaitGroup
var dbFileID = "1AqR6URNNnWPxdFxkd-tBrP2Uv1CAJAjy"
var logFileID = "1UhSUwq0uz6dUsXeHfA7GTRDvzmcMhtZ7"

func main() {
	defer librarian.Close()

	if os.Args[1] == "test" {
		log.Println("init...")
		return
	}

	if os.Args[1] == "onetime" {
		getRealtimeData()
		return
	}

	if os.Args[1] == "upload" {
		transmitter.UpdateFile("./scrolls.db3", dbFileID, "RealtimeData.db")
		transmitter.UpdateFile("./nohup.out", logFileID, "log")
		log.Println("__")
		return
	}

	composer := cron.New()

	composer.AddFunc(os.Args[1], getRealtimeData)
	cronbarie.Add(1)
	composer.Start()
	cronbarie.Wait()
}

func retrieveRealtimeData(start int, limit int) {

	defer barie.Done()
	list := hunter.GetTickers(start, limit)
	for _, value := range list {
		librarian.PushRealtimeData(value)
	}

}

func getRealtimeData() {
	coins := hunter.GetCoinList()
	n := len(coins)
	for i := 1; i < n; i = i + 100 {
		barie.Add(1)
		retrieveRealtimeData(i, 100)
	}
	barie.Wait()
	log.Println("Retrieved ", n, " currencies.")
	transmitter.UpdateFile("./scrolls.db3", dbFileID, "RealtimeData.db")
	transmitter.UpdateFile("./nohup.out", logFileID, "log")
}

func writeData(data []byte, filename string) error {
	os.MkdirAll("output", os.FileMode(0666))
	out, err := os.Create("output/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.Write(data)
	if err != nil {
		return err
	}
	return nil
}
