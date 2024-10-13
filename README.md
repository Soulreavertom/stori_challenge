# stori_challenge
Proyecto codificado en Golang

## How to test?
En la siguiente url cambiar el email al que quieras enviar el reporte, el lambda y api-gateway estan configurados para recibir el parametro del email como parte de la misma ruta de forma dinamica, cambia el valor "youremail@gmail.com" por uno deseado: 

```bash
https://6bl55z8lz6.execute-api.us-west-2.amazonaws.com/dev/accounts/youremail@gmail.com
```

Cuando el proceso se ejecute va a hacer el calculo solicitado, enviara un reporte al e-mail ingresado en la url y guardara en una base de datos todas las solicitudes que se vayan solicitando, al terminar el proceso como forma de respuesta en objeto json regresara los valores previamente guardados en la base de datos.

##Instalar dependencias
Antes de compilar es necesario tener tolas las dependencias instalas, lo podemos hacer con el siguiente comando:

```bash
go mod tidy
```

## ¿Como construir el paquete para deployar en aws lambda?
En la raiz del proyecto hay un archivo Makefile con los comandos necesarios, solo tendrias que ejecutar el comando siguiente en la misma raiz del proyecto, eso te daria como resultado el archivo lambda-handler.zip, ese zip lo subimos al lamba previamente configurado y listo.

```bash
make package_lambda
```

## Estructura del proyecto 

### cmd 
En esta carpeta encontraras el paquete main y responsable de la ejecuion y ademas un archivo de tipo test para pruebas locales.

### db
Paquete encargado de crear la configuración con la base de datos, ests paquete usa variables de ambiente para obtener el usuario, password, url y nombre de la base de datos, los cuales debemos configurar en lambda.

### models
Paquete para crear modelos relacionados con las tablas de las bases de datos mediante tipos de estructuras ademas de crear a esos modelos metodos para guardar o leer la base de datos.

### services
Paquete para alojar funciones mas genericas como el envio de e-mail o la lectura del archivo fuente de datos tipo csv

### tmp 
Carpeta para alojar archivos de forma temporal, en este caso ahi se encuentra una copia del archivo .csv alojado en un bucket s3













