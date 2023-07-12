package config

import (
	//"testing"
	culqi "github.com/culqi/culqi-go"
)

const pk string = "pk_test_e94078b9b248675d"
const sk string = "sk_test_c2267b5b262745f0"
const rsa_id = "de35e120-e297-4b96-97ef-10a43423ddec"
const rsa_public_key = "-----BEGIN PUBLIC KEY-----\n" +
	"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDswQycch0x/7GZ0oFojkWCYv+gr5CyfBKXc3Izq+btIEMCrkDrIsz4Lnl5E3FSD7/htFn1oE84SaDKl5DgbNoev3pMC7MDDgdCFrHODOp7aXwjG8NaiCbiymyBglXyEN28hLvgHpvZmAn6KFo0lMGuKnz8HiuTfpBl6HpD6+02SQIDAQAB\n" +
	"-----END PUBLIC KEY-----"

var Puerto string = ":3000"

const Encrypt = "1"

var EncryptionData = []byte(`{}`)

var (
	publicKey, secretKey string
	encryptionData       []byte
	encrypt              string
)

func init() {
	culqi.Key(pk, sk)

	EncryptionData = []byte(`{
   	   		"rsa_public_key": "` + rsa_public_key + `",
   	   		"rsa_id":  "` + rsa_id + `"
   	   	}`)
}
