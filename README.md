# DEMO - Culqi Go + Checkout V4 + Culqi 3DS

La demo integra Culqi Go, Checkout V4 , Culqi 3DS y es compatible con la v2.0 del Culqi API, con esta demo podrás generar tokens, cargos, clientes, cards.

## Requisitos

* Go 1.6+
* Afiliate [aquí](https://afiliate.culqi.com/).
* Si vas a realizar pruebas obtén tus llaves desde [aquí](https://integ-panel.culqi.com/#/registro), si vas a realizar transacciones reales obtén tus llaves desde [aquí](https://panel.culqi.com/#/registro) (1).

> Recuerda que para obtener tus llaves debes ingresar a tu CulqiPanel > Desarrollo > ***API Keys***.

![alt tag](http://i.imgur.com/NhE6mS9.png)

> Recuerda que las credenciales son enviadas al correo que registraste en el proceso de afiliación.

## Configuración backend

Primero se tiene que modificar los valores del archivo `config/config.go` que se encuentra en al raíz del proyecto. A continuación un ejemplo.

```
var pk string = " Llave pública del comercio (pk_test_xxxxxxxxx)"
var sk string = "Llave secreta del comercio (sk_test_xxxxxxxxx)"
var puerto string = ":3000"
var encrypt = "0" // 1 = activar encriptación
var encryptiondData = []byte(`{		
	"rsa_public_key": "Llave pública RSA que sirve para encriptar el payload de los servicios",
	"rsa_id": "Id de la llave RSA"
}`)
```
## Configuración frontend

Para configurar los datos del cargo, pk del comercio y datos del cliente se tiene que modificar en el archivo `js/config/index.js`.

```js
export default Object.freeze({
    TOTAL_AMOUNT: monto de pago,
    CURRENCY: tipo de moneda,
    PUBLIC_KEY: llave publica del comercio (pk_test_xxxxx),
    RSA_ID: Id de la llave RSA,
    RSA_PUBLIC_KEY: Llave pública RSA que sirve para encriptar el payload de los servicios del checkout,
    COUNTRY_CODE: iso code del país(Ejemplo PE)
});

export const customerInfo = {
    firstName: "Fernando",
    lastName: "Chullo",
    address: "Coop. Villa el Sol",
    phone: "945737476",
}
```

## Inicializar la demo

ejecuta los comandos:

```go
go mod init
go mod tidy
go run main.go
```

## Probar la demo

Para poder visualizar el frontend de la demo ingresar a la siguiente URL:

- Para probar cargos: `http://localhost:3000`
- Para probar creación de cards: `http://localhost:3000/index-card`
