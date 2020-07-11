# Ejercicio realizado por Mercado Pago S.A.
![https://www.mercadopago.com.ar/](https://noticiasmercedinas.com/site/wp-content/uploads/2020/04/qmOFm9ar_400x400.jpg)


API RESTful instanciada en [Google Cloud Platform](https://cloud.google.com/) para obtener información de cadenas ribonucleicas
sobre ADN de mutantes.
Consta de dos (2):
[POST] `/mutant` para postear una cadena y conocer si es un mutante y [GET] `/stats` para obtener el total
de humanos cargados.

La base de datos es un MongoDb instanciado en [Cloud Mongo](https://cloud.mongodb.com/)

## Documentation

Disponible en [https://go-meli-2020.rj.r.appspot.com/api/doc/index.html](https://go-meli-2020.rj.r.appspot.com/api/doc/index.html)

## Requirements

Debe tener las siguientes dependencias instaladas:

*   Golang 1.14+ 


### Building

Para realizar el bildeo ejecutar en consola los siguientes comandos:

```
go mod tidy
export MONGO_CONNECTION="mongodb+srv://gadiel:dalvigay@mercadolibre.edcuz.gcp.mongodb.net/test?retryWrites=true&w=majority"
go run main.go
```


### Testing

Para correr los test unitarios correr el siguiente comando:

```
go test ./... -cover
```


### Documentation

Para actualizar la documentación ejectuar:

```
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

### Using

Para llamar al endpoint `/mutant` enviar un json similar a:

```
{
"dna": [
        "TTGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAAGG",
        "CCCCTA",
        "TCACTG"
    ]
}
```
