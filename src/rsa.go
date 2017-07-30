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
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDY37aK+J1VFPj0
ms1zwWhtcVFZgp+pQ60NEnpHnaXRAkJ9Vh5lH8sVbl60ITBLrDeVDIIsFCIzbsOc
gzWUWCX6YzCWoVV+z2y9hP948dWH+IuS6vWyaDIvK2Se/OrQ9eklxUgD1anB3w2N
bDuP5yX5/uZQ8qNtQPR2XfJwbzvPF1wMjgFdIQqDFK8VYPkozNZDpEnXoyeMKSNv
CwhQVEdmy5pKxt2SCU04KVqkioibY/bkYFsNL+xDKFS3XP+NBuwbSsSKSfRCvQTM
bx0bk2p8mXP/EbXfamgw8dCHLKjFH8SyNJTL7UHAOU1VA0igS4hoZ64vPrtObXo1
x4VgoH/HAgMBAAECggEAPwUjIlcEQQkLHuks8ootMkBPh1cx1hnYW4gXCGCDIq0p
bQptYq1OcjfURTuvsQ2mSSWNXgmrto5lQUHPe//xaUURhXP0SBk2XAwdwKNljVFo
HZUxOFqW2qYdYyXThk0vJs6hEpwv7CpI4eOlmzdsDa1IeBVJ8CGD6DF8lbd9BbwX
DTeKEgR+lWU58j4vMHjA6J3Wi5vHNREj7/1hej67IB2HO2sdw9tZpxhBGPxbzFjh
IvLet009jM8pmWZ32I87b4yA/Ppwtt7SewJH3M7hbCkGEhD3nH37RgUaT9QOxQxc
md5IjH87rnSz6HbeBzbkwhAiV7lz+C1fQ83lxeyLQQKBgQDt0+YWhLajhsDknDfP
7xlYQ5S3vJFb1j27kjmGkQeZhZps/alQx5S5gRNqfWX/FwVgpFh2WcqAzVCNLgJ5
6sooTKn1HXlMcSJuoGGQRknTUKgMmlEIkEpq/mNfdDyyulv6bBInB5eDHi6L64ad
NyjkH3KWA02wTnRCMsuvv3p12QKBgQDpcfCtlIXsr9CZtfkf/rdKk9aV52s7RcAN
rLrnUwDmCAFiTlIRS2GhLbseHBmqQO65lFeqfgDAHvDN0fFG5Z+zvwkeBGam1VH6
iVWChYJlSSk9u9+SgwuqxImZEQYGlqIAEWahdN2Q2hC++gms42V6fh/IUySPF2Y3
8SuPrUT+nwKBgCuM7whWmx9M0MvMM4g3L3vzPCt9M10O2VUczL3wluStC39D4Bpz
wa9KjKSme+P/cjOlRLG/dY6qikKgF0H+FbDVgRXqpMizuutpIRSLWn+AuJU+OPKJ
gacgGFBld6v+5lLwMU7GXB9dioaCxa7KaAmGHLSNvE85vp5UsOno9WLZAoGBAIU9
0V9hrGcBmpCBOuY2Hivk4c/BvISiNRIZCrzJtXOWEiQoC9NvMRQr4+c1prh0Q02N
C9R2ArulbVAsiggrO9H+MZf+bxC+JQJYrIPI7DofXfisaNlZDqmzb83713KfO2vK
gseYynXI4VH/qNITIHKf1kginHVHp8I6VuC0xvKFAoGAE2CTRJLTPHB8R2q81rCD
MxSaQOxhqlS3d0UIPhc22ioqwtPCH09YT/D1N1YaBu270eCbJ90psWMG6TzXbZZ5
pFwlIf0xdE5J0thEIyKlckp6VC6jsJ55J9zHgx7WpwaAAjbaMx226zsA0XcGNjbF
8jHhWprkc0XsWU1pf4FjMX0=
-----END PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2N+2ividVRT49JrNc8Fo
bXFRWYKfqUOtDRJ6R52l0QJCfVYeZR/LFW5etCEwS6w3lQyCLBQiM27DnIM1lFgl
+mMwlqFVfs9svYT/ePHVh/iLkur1smgyLytknvzq0PXpJcVIA9Wpwd8NjWw7j+cl
+f7mUPKjbUD0dl3ycG87zxdcDI4BXSEKgxSvFWD5KMzWQ6RJ16MnjCkjbwsIUFRH
ZsuaSsbdkglNOClapIqIm2P25GBbDS/sQyhUt1z/jQbsG0rEikn0Qr0EzG8dG5Nq
fJlz/xG132poMPHQhyyoxR/EsjSUy+1BwDlNVQNIoEuIaGeuLz67Tm16NceFYKB/
xwIDAQAB
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
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priKey := priv.(*rsa.PrivateKey)
	return rsa.DecryptPKCS1v15(rand.Reader, priKey, ciphertext)
}