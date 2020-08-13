# Twittor server

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Twittor server es una API REST, creada con **Golang** para ser utilizada como sistema de **Tweets**.

**Librerías implementadas:**
- **JWT**:               [github.com/dgrijalva/jwt-go](github.com/dgrijalva/jwt-go)
- **Cors**:              [github.com/rs/cors](github.com/rs/cors)
- **Crypto**:            [golang.org/x/crypto](golang.org/x/crypto)
- **Go Delve**:          [github.com/go-delve/delve](github.com/go-delve/delve)
- **Go Dotenv**:         [github.com/joho/godotenv](github.com/joho/godotenv)
- **Gorilla Mux**:       [github.com/gorilla/mux](github.com/gorilla/mux)
- **Mongodb Driver**:    [go.mongodb.org/mongo-driver](go.mongodb.org/mongo-driver)

**Galex** es un marco interno creado para brindar soporte en tareas como:
- Registro de **Rutas y Middlewares**.
- Registro de **Controladores**.
- Registro de **Modelos de Datos**.
- **Helper de mongodb** para ser implementado en los modelos de datos.
- Paquete de utilidades para el manejo de solicitudes y respuestas (**Writer & Request**).

### Instalación

Twittor server require [Golang](https://golang.org/) v1.14.
Para realizar instalación debe seguir las siguientes instrucciones.

```sh
$ git clone https://github.com/devJGuerrero/twittor-server.git
```
```sh
$ cd twittor-server
```
```sh
$ go run main.go
```

### Versión de vista previa

Credenciales: **Email**: demo@twittor.com - **Password**: 123456

###### Autentificación
> GET Sign http://twittorbackend.herokuapp.com/api/sign

###### Perfiles
> GET Info Profile http://twittorbackend.herokuapp.com/api/profiles?id=123456

> GET Avatar http://twittorbackend.herokuapp.com/api/profiles/123456/avatar

> GET Banner http://twittorbackend.herokuapp.com/api/profiles/123456/banner

> POST Banner http://twittorbackend.herokuapp.com/api/users
```json
{
    "name": "string required",
    "lastName": "string",
    "dateBirth": "date 0000-00-00T00:00:00.000+00:00",
    "email": "string required",
    "password": "string required",
    "avatar": "string",
    "banner": "string",
    "biography": "string",
    "location": "string",
    "website": "string"
}
```
