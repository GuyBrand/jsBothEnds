// main.go
package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type generalRet struct {
	Status   string      `json:"status"`
	Value    interface{} `json:"value"`
}

func main() {
	http.HandleFunc("/shippingRate/", folderHandler)
	http.HandleFunc("/getShippingRate", getShippingRate)

	http.ListenAndServe(":2014", nil)
}

func folderHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "."+request.URL.Path)
}

func readPostData(w http.ResponseWriter, req *http.Request,location string)(postData []byte,err error){
	if postData, err = ioutil.ReadAll(req.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logError("reaPostData " + location + " err :",err)
	}
	return
}

func writeGeneralResponse(location string, value interface{},origData []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	gr := generalRet{Status:"ok",Value:value}

	if ret,err := json.Marshal(gr);err!=nil{
		logError("write resp marshal",err)
		w.WriteHeader(http.StatusInternalServerError)
	} else if ret,err = runHook(location,"post",origData,ret);err!=nil {
		logError("run post hook for " + location, err)
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		w.Write(ret)
	}
}

