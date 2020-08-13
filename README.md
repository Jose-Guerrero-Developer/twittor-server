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
> GET Info Profile http://twittorbackend.herokuapp.com/api/profiles?id=5f21b569c826360941d243c9

> GET Avatar http://twittorbackend.herokuapp.com/api/profiles/5f21b569c826360941d243c9/avatar

> GET Banner http://twittorbackend.herokuapp.com/api/profiles/5f21b569c826360941d243c9/banner

> POST Store http://twittorbackend.herokuapp.com/api/users
```json
{
    "name": "Joe",
    "lastName": "Doe",
    "dateBirth": "1993-01-01T00:00:00.000+00:00",
    "email": "joe@gmail.com",
    "password": "123456",
    "avatar": "",
    "banner": "",
    "biography": "Desarrollador de software",
    "location": "",
    "website": "http://joe.com"
}
```

> PUT Update http://twittorbackend.herokuapp.com/api/users
```json
{
     "name": "Joe",
    "lastName": "Doe",
    "dateBirth": "1993-01-01T00:00:00.000+00:00",
    "email": "joe@gmail.com",
    "password": "123456",
    "avatar": "http://joe/asstes//img/avatars/001.png",
    "banner": "http://joe/assets/img/banners/001.png",
    "biography": "Desarrollador de software",
    "location": "Todo el mundo",
    "website": "http://joe.com"
}
```

###### Tweets
> GET ID http://twittorbackend.herokuapp.com/api/tweets/5f2c795d181587349d099e57

> GET ALL http://twittorbackend.herokuapp.com/api/tweets

> GET ALL Profile http://twittorbackend.herokuapp.com/api/tweets/profile/5f21b569c826360941d243c9

> GET ALL Follow Profile http://twittorbackend.herokuapp.com/api/followers/tweets

> POST Store http://twittorbackend.herokuapp.com/api/tweets

```json
{
    "message": "Publicación sobre temas de artes y ciencias",
    "idProfile": "5f32b1927d0a043587783a81"
}
```

> PUT Update http://twittorbackend.herokuapp.com/api/tweets/5f2c795d181587349d099e57

```json
{
    "message": "Publicación sobre temas de Golang"
}
```

> DELETE Remove http://twittorbackend.herokuapp.com/api/tweets/5f31a814fd5aa58df521afc7

###### Follow
> GET ID Profile http://twittorbackend.herokuapp.com/api/followers/profile/5f21b569c826360941d243c9/follow/5f2c73c255767c0f4f1d72f6

> GET ALL Profile http://twittorbackend.herokuapp.com/api/followers/profile/5f21b569c826360941d243c9

> GET ALL Followerd Profile http://twittorbackend.herokuapp.com/api/followers

> POST Store http://twittorbackend.herokuapp.com/api/followers

```json
{
    "idProfile": "5f32b1927d0a043587783a81",
    "idFollow": "5f21b569c826360941d243c9"
}
```

> DELETE Remove http://twittorbackend.herokuapp.com/api/followers

###### Uploads

> POST Avatar http://twittorbackend.herokuapp.com/api/profiles/5f21b569c826360941d243c9/upload/avatar

form file field avatar

> POST Banner http://twittorbackend.herokuapp.com/api/profiles/5f21b569c826360941d243c9/upload/banner

form file field banner
