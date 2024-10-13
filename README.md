# stori_challenge

Proyecto desarrollado en Golang.

## Cómo probar

Para probar la funcionalidad, utiliza la siguiente URL. Cambia el correo electrónico al que deseas enviar el reporte. El Lambda y API Gateway están configurados para recibir el parámetro del email como parte de la ruta de forma dinámica. Sustituye "youremail@gmail.com" por el correo deseado:

```bash
https://6bl55z8lz6.execute-api.us-west-2.amazonaws.com/dev/accounts/youremail@gmail.com
```

Cuando se ejecute el proceso, se realizará el cálculo solicitado, se enviará un reporte al email ingresado en la URL y se guardarán en una base de datos todas las solicitudes realizadas. Al finalizar el proceso, se devolverá un objeto JSON con los valores previamente guardados en la base de datos.

## Instalación de dependencias

Antes de compilar el proyecto, es necesario tener instaladas todas las dependencias. Esto se puede hacer con el siguiente comando:

```bash
go mod tidy
```

## Cómo construir el paquete para desplegar en AWS Lambda

En la raíz del proyecto hay un archivo `Makefile` que contiene los comandos necesarios. Solo necesitas ejecutar el siguiente comando en la raíz del proyecto, lo que generará el archivo `lambda-handler.zip`. Este archivo se sube al Lambda previamente configurado y estará listo para usarse.

```bash
make package_lambda
```

## Estructura del proyecto

### cmd 
En esta carpeta encontrarás el paquete principal, responsable de la ejecución, así como un archivo de prueba para realizar pruebas locales.

### db
Paquete encargado de configurar la conexión con la base de datos. Este paquete utiliza variables de entorno para obtener el usuario, contraseña, URL y nombre de la base de datos, los cuales debemos configurar en Lambda.

### models
Paquete diseñado para crear modelos relacionados con las tablas de la base de datos mediante estructuras. Además, se incluyen métodos para guardar y leer datos en la base de datos.

### services
Paquete que alberga funciones más genéricas, como el envío de correos electrónicos o la lectura de archivos de datos en formato CSV.

### tmp 
Carpeta destinada a almacenar archivos de forma temporal; en este caso, contiene una copia del archivo .csv alojado en un bucket S3.

| Id | Date  | Transaction |
|----|-------|-------------|
| 0  | 7/15  | +60.5      |
| 1  | 7/16  | +10.2      |
| 2  | 7/28  | -10.3      |
| 3  | 7/29  | -9.7       |
| 4  | 8/12  | -20.46     |
| 5  | 8/13  | +11.5      |
| 4  | 8/14  | -5.52      |
| 5  | 8/15  | +10        |

## Variables de Entorno

Es necesario configurar las siguientes variables de entorno en AWS Lambda:

| Variable de Entorno | Descripción                                   |
|----------------------|-----------------------------------------------|
| `DBNAME`             | Nombre de la base de datos                    |
| `DBPASS`             | Contraseña para acceder a la base de datos    |
| `DBPORT`             | Puerto utilizado para la conexión             |
| `DBURL`              | URL o dirección IP para la conexión          |
| `DBUSER`             | Nombre de usuario para acceder a la base de datos |
| `EMAIL`              | Cuenta de Gmail utilizada para enviar correos  |
| `KEYM`               | Contraseña de la cuenta de correo electrónico  |
