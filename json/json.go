package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	pF  = fmt.Printf
	pLn = fmt.Println
)

type Person1 struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	intVal, _ := json.Marshal(500)
	pF("intVal: %v, type: %T \n", intVal, intVal)
	pF("intVal: %#v, type: %T \n", intVal, intVal)
	pF("intVal: %v, type: %T \n", string(intVal), string(intVal))

	pLn()

	strVal, _ := json.Marshal("Galatasaray")
	pF("strVal: %v, type: %T \n", strVal, strVal)
	pF("strVal: %#v, type: %T \n", strVal, strVal)
	pF("strVal: %v, type: %T \n", string(strVal), string(strVal))

	pLn()

	boolVal, _ := json.Marshal(true)
	pF("boolVal: %v, type: %T \n", boolVal, boolVal)
	pF("boolVal: %#v, type: %T \n", boolVal, boolVal)
	pF("boolVal: %v, type: %T \n", string(boolVal), string(boolVal))

	pLn()

	fltVal, _ := json.Marshal(10.50)
	pF("fltVal: %v, type: %T \n", fltVal, fltVal)
	pF("fltVal: %#v, type: %T \n", fltVal, fltVal)
	pF("fltVal: %v, type: %T \n", string(fltVal), string(fltVal))

	pLn()

	sliceVal, _ := json.Marshal([]string{"yellow", "purple"})
	pF("sliceVal: %v, type: %T \n", sliceVal, sliceVal)
	pF("sliceVal: %#v, type: %T \n", sliceVal, sliceVal)
	pF("sliceVal: %v, type: %T \n", string(sliceVal), string(sliceVal))

	pLn()

	mapVal, _ := json.Marshal(map[string]int{"yellow": 5, "purple": 10})
	pF("mapVal: %v, type: %T \n", mapVal, mapVal)
	pF("mapVal: %#v, type: %T \n", mapVal, mapVal)
	pF("mapVal: %v, type: %T \n", string(mapVal), string(mapVal))

	pLn()

	structVal1, _ := json.Marshal(&Person1{Name: "John", Age: 35})
	pF("structVal1: %v, type: %T \n", structVal1, structVal1)
	pF("structVal1: %#v, type: %T \n", structVal1, structVal1)
	pF("structVal1: %v, type: %T \n", string(structVal1), string(structVal1))

	pLn()

	structVal2, _ := json.Marshal(&Person2{Name: "John", Age: 35})
	pF("structVal2: %v, type: %T \n", structVal2, structVal2)
	pF("structVal2: %#v, type: %T \n", structVal2, structVal2)
	pF("structVal2: %v, type: %T \n", string(structVal2), string(structVal2))

	pLn()

	byt := []byte(`{"name": "Product 1", "price": 10.50}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pF("dat: %v, type: %T \n", dat, dat)
	pF("dat['name']: %v, type: %T \n", dat["name"], dat["name"])
	pF("dat['price']: %v, type: %T \n", dat["price"], dat["price"])
	pF("dat['price']: %v, type: %T \n", dat["price"].(float64), dat["price"].(float64))

	pLn()

	str := `{"name": "John", "age": 35}`
	inter := Person2{}
	if err := json.Unmarshal([]byte(str), &inter); err != nil {
		panic(err)
	}
	pF("dat: %+v, type: %T \n", inter, inter)
	pF("dat: %v, type: %T \n", inter.Name, inter.Name)
	pF("dat: %v, type: %T \n", inter.Age, inter.Age)

	pLn()

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
