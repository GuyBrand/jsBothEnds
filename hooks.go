// main.go
package main

import (
	"io/ioutil"
	"os"
)

var hooksCache map[string]string //cache files instead of reading each time

func init() {
	hooksCache = make(map[string]string)
}

func addToHookCache(key, value string) {
	//TODO : sync.Map{}
	//Simple : hooksCache[key] = value
	//Currently : no cache - better for demo
}

func getHook(location, stage string) string {
	filename := GetCurrentDir() + "/scripts/" + location + "/" + stage + ".js"

	if content, exists := hooksCache[filename]; exists {
		return content
	} else if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if data, err := ioutil.ReadFile(filename); err == nil { //TODO : what if err
			content = string(data)
			addToHookCache(filename, content)
			return content
		}
	}
	addToHookCache(filename, "")
	return ""
}

func runHook(location, stage string, origData, curReply []byte) ([]byte, error) {
	if script := getHook(location, stage); script == "" {
		if stage == "post" {
			return curReply, nil
		}
		return origData, nil
	} else {
		return scriptRet(script, location, stage, origData, curReply)
	}
}

func bytesToJson(b []byte) string {
	if len(b) == 0 {
		return "{}"
	}
	return string(b) //no need to escape, client side already did !
}

func scriptRet(script, location, stage string, origData []byte, curReply []byte) (reply []byte, err error) {
	vm := GetDbEnabledJsVm()
	withReply := len(curReply) > 0
	preScript := "var origData = " + bytesToJson(origData) + "\n"
	if withReply {
		preScript += "var curReply = " + bytesToJson(curReply) + "\n"
	}
	if ret, err := vm.Run(preScript + script); err != nil {
		return origData, err
	} else if bt, err := convertOttoValToBytes(ret); err != nil {
		return origData, err
	} else {
		return bt, err
	}
}
