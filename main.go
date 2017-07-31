package main

import (
	src "./src"
)

func main() {
	src.StartMartini();
	/*data,err := src.RsaEncrypt([]byte(src.Pay(1, "123456", 1.00)));
	if(err != nil){
		fmt.Printf("%s", err);
	}
	origData,err := src.RsaDecrypt(data)
	if(err != nil){
		fmt.Printf("%s", err);
	}
	fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	fmt.Println("rsa decrypt :" + string(origData))*/
}