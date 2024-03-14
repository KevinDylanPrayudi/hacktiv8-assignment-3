package main

import (
	"assignment-3/helpers"
	"assignment-3/routers"
	"assignment-3/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/radovskyb/watcher"
)

var cwd, _ = os.Getwd()
var data map[string]map[string]structs.Data

func init() {
	_, err := os.ReadFile(cwd + "\\test.json")
	if err != nil {
		data := map[string]structs.Data{
			"Status": {
				Water: rand.Intn(100+1-1) + 1,
				Wind:  rand.Intn(100+1-1) + 1,
			},
		}

		datas, _ := json.MarshalIndent(data, "", " ")
		_ = os.WriteFile("test.json", datas, 0644)
	}
	log.Println("test.json exist")
}

func main() {
	go helpers.WriteJson()
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-w.Event:
				{
					resp, err := http.Get("http://localhost:8080/ping")
					if err != nil {
						log.Println(err)
					}
					defer resp.Body.Close()
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Println(err)
					}
					err = json.Unmarshal(body, &data)
					if err != nil {
						log.Println(err)
					}
					contents := `
	Water level: %d m %s
	Wind Speed: %d m/s %s
					`
					var waterStat, windStat string
					if data["data"]["Status"].Water > 8 {
						waterStat = "Danger"
					} else if data["data"]["Status"].Water > 5 {
						waterStat = "StandBy"
					} else {
						waterStat = "Safe"
					}
					if data["data"]["Status"].Wind > 8 {
						windStat = "Danger"
					} else if data["data"]["Status"].Wind > 5 {
						windStat = "StandBy"
					} else {
						windStat = "Safe"
					}
					fmt.Printf(contents, data["data"]["Status"].Water, waterStat, data["data"]["Status"].Wind, windStat)
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	go func() {
		if err := w.Add(cwd + "\\test.json"); err != nil {
			log.Fatalln(err)
		}
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()
	routers.Router()
}
