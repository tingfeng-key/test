package src

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
)

var decrypted string

func init() {
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
}

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDmzZenWOqOPdSHM+usjNcyV/1EU0092+CH+yw+g4+FhLrRhVmz
KxF1H3ZLN5gbJs3r6kPPlmdq8Bw6e1hFzgseB0Emmu+rN6y+TMUFogJ+hMbH/2GY
tay/3rxcAwUFID8sahnw0sLYtBnBZ++4WlNNyNaNjBfCiHcr/R6wRQlaiQIDAQAB
AoGAO19qAs+5zwuSz5EIsGOuvDGsYlY3iZE1imWPtSks04EYeeue7ptGq2WDjcfo
ZhuuX2DtPsqHIhhQeq7eP17Ittyrm3H9jMPS4UZLA9YktCq4Z9eZSwKFOhZ2IkFx
8Ba3HTWPvKYf2ODKMTyHUUUjgUxyrLosT+fBuYCWPSf1eLECQQD8K8ZidCPlsf/R
lzl3GjB/NjNqViiHpQWl6h05/gz257c9PD+Q5X9b+qJllKdqj8zw/n5B+wHZp7ii
lWt8HYPtAkEA6k7B4UbHygFZEx0AT0gtOzv2W1sO0WJi7fEa4TLMuR+Ep0lRoC9U
3sdtvxgPD9o+UcNMqyw5iZFXSZgMIKJVjQJBALnfvmxFtUwTB0DDWi1LVaH/Hqjr
9xX88ovYZFChzYaVSADQDv7L2zO+K8P6bejjNCoTPJd/4F5B/ZuBzed/jgUCQQDl
LzQljn5N8ROXxn5PDmJv1i8bHO5ZkMsDwxpvvKjTSykOGIFnFeUGJDwM9xOEhgz/
HPEZXaDKFwgyTC9QnHjhAkEA7M4GIFNSwS7+gCg4z4GskBLE3kWe/3LB9Jf2ZnWh
xHr7ylycGIsjZsfpgXq/HYhrTtoTRlEM7PWBdNdWUkz4jg==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDmzZenWOqOPdSHM+usjNcyV/1E
U0092+CH+yw+g4+FhLrRhVmzKxF1H3ZLN5gbJs3r6kPPlmdq8Bw6e1hFzgseB0Em
mu+rN6y+TMUFogJ+hMbH/2GYtay/3rxcAwUFID8sahnw0sLYtBnBZ++4WlNNyNaN
jBfCiHcr/R6wRQlaiQIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}