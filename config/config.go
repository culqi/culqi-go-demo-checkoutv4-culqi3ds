package config

import (
	//"testing"
	culqi "github.com/culqi/culqi-go"
)

const pk string = "INSERTE_SU_LLAVE_PUBLICA"
const sk string = "INSERTE_SU_LLAVE_PRIVADA"
const rsa_id = "INSERTE_SU_RSA_ID"
const rsa_public_key = "INSERTE_SU_LLAVE_RSA"

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
   	   		"rsa_public_key": "INSERTE_SU_LLAVE_RSA",
   	   		"rsa_id":  "` + rsa_id + `"
   	   	}`)
}
