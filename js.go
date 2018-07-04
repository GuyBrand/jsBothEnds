// main.go
package main

import (
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
)

func GetDbEnabledJsVm() *otto.Otto {
	vm := otto.New()

	vm.Set("now", NowTimeStr)
	vm.Set("logToFile", logFromJs)
	//vm.Set("ret",nil)
	vm.Run(`function ret(value){
					return JSON.stringify(value);
				 }`)
	//	vm.Set("getData", getData)

	//for key, obj := range JsObjects {
	//	vm.Run(key + "=" + SprintAsJson(obj))
	//}

	return vm
}

func logFromJs(call otto.FunctionCall) otto.Value { // allows to write the js log to go log file
	if parms, err := getArgumentList(call, 0); err != nil {
		return reterr(err)
	} else {
		fmt.Println(parms)
	}
	return reterr("")
}

func getArgumentList(call otto.FunctionCall, fromParm int) ([]interface{}, error) {
	var parms []interface{}
	for i := fromParm; i < len(call.ArgumentList); i++ {
		if val, err := call.ArgumentList[i].Export(); err != nil {
			return parms, err
		} else {
			parms = append(parms, val)
		}
	}
	return parms, nil

}

func reterr(v interface{}) otto.Value {
	val, _ := otto.ToValue(v)
	return val
}

func convertOttoValToBytes(value otto.Value) (bt []byte, err error) {
	if !value.IsString() {
		return bt, errors.New(" val is not an string")
	} else if retStr, err := value.ToString(); err != nil {
		return bt, err
	} else {
		return []byte(retStr), nil
	}
}

/*
func getData(call otto.FunctionCall) otto.Value {
	if script, err := call.Argument(0).ToString(); err != nil {
		return reterr(err)
	} else {
		if parms, err := getArgumentList(call, 1); err != nil {
			return reterr(err)
		} else {
			if len(parms) == 1 && reflect.TypeOf(parms[0]).Kind() != reflect.String &&
				reflect.TypeOf(parms[0]).String() == "[]interface {}" {
				parms = parms[0].([]interface{})
			}
			if ret, err := NewGetAsJson(script, parms...); err != nil {
				fmt.Println("getData err: ", err.Error(), "\nscript: ", script, "\nparms: ", parms)
				//PrintParamsSlice(parms)
				return reterr(err)
			} else {
				return setRetValue(ret)
			}

		}
	}
}
*/
