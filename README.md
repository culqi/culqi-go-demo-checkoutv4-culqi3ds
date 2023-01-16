# DEMO - Culqi Go + Checkout V4 + Culqi 3DS

La demo integra Culqi Go, Checkout V4 , Culqi 3DS y es compatible con la v2.0 del Culqi API, con esta demo podrás generar tokens, cargos, clientes, cards.

## Requisitos

* Go 1.6+
* Afiliate [aquí](https://afiliate.culqi.com/).
* Si vas a realizar pruebas obtén tus llaves desde [aquí](https://integ-panel.culqi.com/#/registro), si vas a realizar transacciones reales obtén tus llaves desde [aquí](https://panel.culqi.com/#/registro) (1).

> Recuerda que para obtener tus llaves debes ingresar a tu CulqiPanel > Desarrollo > ***API Keys***.

![alt tag](http://i.imgur.com/NhE6mS9.png)

> Recuerda que las credenciales son enviadas al correo que registraste en el proceso de afiliación.

## Configuración

Dentro del archivo **main.go** coloca tus llaves pk y sk.

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
