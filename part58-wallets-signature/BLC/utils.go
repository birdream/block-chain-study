package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

// 将int64转字节数组
func IntToHex(d int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, d)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// JSON字符串转成数组
func JSONToArray(jsonStr string) []string {
	var sArr []string
	if err := json.Unmarshal([]byte(jsonStr), &sArr); err != nil {
		log.Panic(err)
	}

	return sArr
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
