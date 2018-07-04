// main.go
package main

import (
	"encoding/json"
	"net/http"
)

type shipRate struct {
	Vendor   string `json:"vendor"`
	Customer string `json:"customer"`
}

func getShippingRate(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		var sr shipRate
		if data, err := readPostData(w, req, "getShippingRate"); err != nil {
			logError("get shipping rate postData", err)
		} else if dataAfterHook, err := runHook("getShippingRate", "pre", data, nil); err != nil {
			logError("shipping rate pre hook ", err)
		} else if err := json.Unmarshal(dataAfterHook, &sr); err != nil {
			logError("decode shipping rate", err)
		} else if ret, err := getShipRateFromFakeDb(sr); err != nil {
			logError("get shipping rate from db", err)
		} else {
			writeGeneralResponse("getShippingRate", ret, dataAfterHook, w)
		}
	}
}

func createRateReply(rate float64) generalRet {
	return generalRet{Status: "ok", Value: rate}
}


func getShipRateFromFakeDb(rate shipRate) (float64, error) {
	rates := map[string]float64{
		"Dominos Ritza": 10,
		"Coca Yolla":    12,
		"Live-Ice":      18,
		"LA-Beard":      8,
		"Not-Ella":      11,
		"Mercy-Days":    9,
	}
	return rates[rate.Vendor], nil //select rate from shipping rates where vandor = ....
}

