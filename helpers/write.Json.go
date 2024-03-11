package helpers

import (
	"assignment-3/structs"
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func WriteJson() {
	for range time.Tick(time.Second * 15) {
		data := structs.Data{
			Water: rand.Intn(10),
			Wind:  rand.Intn(10),
		}
		datas, _ := json.MarshalIndent(data, "", " ")
		_ = os.WriteFile("test.json", datas, 0644)
	}
}
