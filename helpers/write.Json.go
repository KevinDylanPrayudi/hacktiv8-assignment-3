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
		data := map[string]structs.Data{
			"Status": {
				Water: rand.Intn(100+1-1) + 1,
				Wind:  rand.Intn(100+1-1) + 1,
			},
		}

		datas, _ := json.MarshalIndent(data, "", " ")
		_ = os.WriteFile("test.json", datas, 0644)
	}
}
