package main

import (
	src "./src"
	"fmt"
	"encoding/base64"
)

func main() {
	//src.ExampleScrape()
	data,_ := src.RsaEncrypt([]byte(src.Pay(1, "123456", 1.00)));
	origData,_ := src.RsaDecrypt(data)
	fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	fmt.Println("rsa decrypt :" + string(origData))
}