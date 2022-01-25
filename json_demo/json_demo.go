package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"name"`
	ServerIp   string `json:"ip"`
	ServerPort int    `json:"port"`
}

func TestSerializeStruct() {
	server := new(Server)
	server.ServerPort = 80
	server.ServerIp = "localhost"
	server.ServerName = "json"

	by, err := json.Marshal(server)
	if err != nil {
		fmt.Println(" Marshall error : ", err.Error())
	}
	fmt.Println(" marshal json ", string(by))
}

func TestSerializeMap() {
	mapSerial := make(map[string]string)
	mapSerial["11"] = "11"
	mapSerial["22"] = "22"
	mapSerial["33"] = "33"

	mapByte, err := json.Marshal(mapSerial)
	if err != nil {
		fmt.Println("map marshall error ", err.Error())
	}

	fmt.Println("mapr serialize : ", string(mapByte))
}

// 反序列化
func TestDeserializeStruct() {
	jsonStr := "{\"name\":\"json\",\"ip\":\"localhost\",\"port\":80}"
	server := new(Server)
	err := json.Unmarshal([]byte(jsonStr), server)
	if err != nil {
		fmt.Println(" json unmarshall error", err.Error())
	}

	fmt.Println("Unmarshal struct : ", server)
}

func TestDeSerializeMap() {
	jsonStr := "{\"11\":\"11\",\"22\":\"22\",\"33\":\"33\"}"
	mMap := new(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), mMap)

	if err != nil {
		fmt.Println("map unmarshal ", err.Error())
	}

	fmt.Println("map unmarshal ", mMap)
}
