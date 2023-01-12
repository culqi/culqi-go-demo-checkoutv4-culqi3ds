# Demo Go - Checkout v4 + 3DS

La demo integra el checkout v4 y nuestra libreria Culqi 3DS.
Podra simular pruebas:

- Creación de cargo
- Creación de card.

> Recuerda que para usar cualquier plugins necesitas tener tu llave pública y llave privada (test o live), los cuales los puedes generar a través de tu Culqipanel.

## Requisitos

- Go 1.6+
- [Credenciales de Culqi](https://www.culqi.com)

## Configuración

Dentro del archivo main.go puedes colocar tus llaves y el puerto en el cual quiere levantar la demo

## Ejecutar la demo

ejecuta los comandos:

```go
go mod init
go mod tiny
go run main.go
```

