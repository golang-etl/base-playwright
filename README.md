# Golang ETL / Base Playwright

Proyecto base para iniciar un servicio basado en Playwright.

## Prerequisitos

Si se esta iniciando un nuevo proyecto, se recomienda crear un nuevo repositorio y clonar el repositorio base.
Adicionalmente, se recomienda ejecutar estos comandos reemplazando los siguientes valores:

```bash
export BASE_MS_MODULE_NAME=github.com/golang-etl/base-playwright
export BASE_MS_EXAMPLE_PROVIDER=login
export BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE=login
export BASE_MS_EXAMPLE_PROVIDER_PASCAL_CASE=Login

LC_CTYPE=C find . -type f -not -path "./.git/*" -exec sed -i '' "s|github.com/golang-etl/base-playwright|$BASE_MS_MODULE_NAME|g" {} +
LC_CTYPE=C find . -type f -not -path "./.git/*" -exec sed -i '' "s|login|$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE|g" {} +
LC_CTYPE=C find . -type f -not -path "./.git/*" -exec sed -i '' "s|Login|$BASE_MS_EXAMPLE_PROVIDER_PASCAL_CASE|g" {} +

mv src/controllers/http/login.go src/controllers/http/$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE.go
mv src/providers/login src/providers/$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE
mv src/providers/$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE/login.go src/providers/$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE/$BASE_MS_EXAMPLE_PROVIDER_CAMEL_CASE.go
```

## Dependencias privadas

Si se quiere instalar dependencias privadas, recomiendo modificar el Dockerfile para incluir las credenciales de acceso a los repositorios privados.
De la siguiente manera, se modifica el Stage 1:

```bash
#############################
# Stage 1: Modules caching
#############################
FROM golang:1.24 AS modules

COPY go.mod go.sum /modules/

WORKDIR /modules

ARG GITHUB_TOKEN

RUN git config --global credential.helper store
RUN git config --global url."https://$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"
RUN go mod download
```

Y se le pasa la variable de entorno `GITHUB_TOKEN` a la hora de construir la imagen:

```bash
docker build --build-arg GITHUB_TOKEN=your_github_token -t your_image_name .
```

## Variables de entorno

Las variables de entorno se pueden definir en un archivo `.env` en la raíz del proyecto, base en el archivo `.env.example`.

# Generación de secretos

Es recomendable generar un secreto para cada servicio que se implemente. Este secreto se utilizará para encriptar y desencriptar. Para generar un secreto, se puede utilizar el siguiente comando en MacOS:
```
LC_ALL=C tr -dc 'A-Za-z0-9!@#%^*-_=+' < /dev/urandom | head -c 32
```
