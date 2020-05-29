# checkCpfRestApiGoBind

RESTful API to verify physical person registers status on Brazilian entities.<br>
A GO (Golang) API solution that dynamically **binds** the library [CURL](https://curl.haxx.se) to manage HTTP requests.

The routing API receives a CPF (Brazilian Physical Person Register) to be checked.<br>
Then the CPF is validated on SERPRO (Brazilian Federal Data Processing Service) that also uses a RESTful API.<br>

The HTTP request implementation uses **go-curl** ([github.com/andelf](https://github.com/andelf/go-curl)) package that dynamically links with the pre-installed **libCurl** ([github.com/curl](https://github.com/curl/curl)).<br>
The libCurl is implemented in C/C++, and go-curl is powered by GO and C languages.

The API may also receives more informative data, like Name and RG (Brazilian General Registry), not being validated now.<br>
The API routing implementation uses the library **gorilla/mux** ([github.com/gorilla/mux](https://github.com/gorilla/mux)).

## Quick Start (on Linux)

### Install GO compiler
```bash
sudo apt install gccgo-go
```
This application was tested with GO release `1.12.2`. To check the installed version of GO compiler, run:
```bash
go version
```

### Install Gorilla Mux router library
```bash
go get -u github.com/gorilla/mux
```

### Install lib Curl for Development
```bash
sudo apt-get install libcurl-dev
```
This application was tested with the library `libcurl4-openssl-dev` on release `7.68.0`

### Install Go-Curl binding library
```bash
go get -u github.com/andelf/go-curl
```

### Build and run checkCpfRestApiGoBind (from the project root folder)
```bash
go build
./checkCpfRestApiGoBind
```

## API description

**URL** : `api/v1/verify`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Required fields** : `cpf`

**Non-required fields** : `name`, `rg` (`number`, `issued`, `entity`)

**Request example:**

```json
 {
  "name":"Heitor Peralles",
  "cpf":"40442820135",
  "rg": { "number":"209921899", "issued":"2020/5/20", "entity":"DETRAN-RJ" }
 }
```

### Success Response

**Code** : `200 Register OK`

**Response example:**

```json
 {
  "status":"True"
 }
```

### Error Responses

**Code** : `400 Invalid format or registry number`

**Code** : `403 Problem with the registry, NOT OK`

**Code** : `500 Server Error`

**Response example:**

```json
 {
  "status":"False",
  "message":"Invalid CPF Format."
 }
```

## Testing the API

### RUN Unit tests (from the project root folder)

The file `app_test.go` contains a collection of unit tests. Run the tests by calling the command below.

``` bash
go test -v
```

### Testing manually

#### A simple way to test the API on Linux:

Install CURL, the recommended library to deal with requests.

``` bash
sudo apt-get install curl
```

Trigger a request to the primary endpoint as follows.

``` bash
curl localhost:8000/api/v1/verify -i -X POST -d '{"cpf":"40442820135"}'
```

Remembering that checkCpfRestApiGoBind must be running to receive the request.

#### List of testing CPFs provided by SERPRO:

The application uses a service provided by SERPRO, so any test depends on this service being online.<br>
The CPFs provided by SERPRO for testing are:

```
40442820135: Regular (Register OK)
63017285995: Regular (Register OK)
91708635203: Regular (Register OK)
58136053391: Regular (Register OK)
40532176871: Suspended (Problem with the registry)
47123586964: Suspended (Problem with the registry)
07691852312: Regularization Pending (Problem with the registry)
10975384600: Regularization Pending (Problem with the registry)
01648527949: Canceled by Multiplicity (Problem with the registry)
47893062592: Canceled by Multiplicity (Problem with the registry)
98302514705: Null (Problem with the registry)
18025346790: Null (Problem with the registry)
64913872591: Registration Canceled (Problem with the registry)
52389071686: Registration Canceled (Problem with the registry)
05137518743: Deceased Holder (Problem with the registry)
08849979878: Deceased Holder (Problem with the registry)
```

## Implementation

### File structure

The application is composed by a set of source code files:

**app.go** : Main file, with `main()` function.<br>
**app_test.go** : File containing all application tests.

**api_model.go** : API structures.<br>
**api_routing.go** : API routing implementation using Gorilla Mux.<br>
**api.go** : Code to handle provided API requests and responses.

**middleware_model.go** : SERPRO service structures.<br>
**middleware.go** : Core implementation of the request using go-curl library.

**README.md** : This file, explaining how the application works.<br>
**.gitignore** : List of non-versioned files (such as the compiled binary).

After the compilation process, the non-versioned binary file `checkCpfRestApiGoBind` will be generated.

### API Credentials

The application is configured with a personal TOKEN, please change it.<br>
Its possible to generate a new TOKEN on SERPRO service page: [servicos.serpro.gov.br](https://servicos.serpro.gov.br/inteligencia-de-negocios-serpro/biblioteca/consulta-cpf/teste.html).

The in use TOKEN is specified on the top of `middleware.go` file.<br>
If going to production the SERPRO service URL must also be changed (at the same file).

## Application Info

### Author

Heitor Peralles<br>
[heitorgp@gmail.com](mailto:heitorgp@gmail.com)<br>
[linkedin.com/in/heitorperalles](https://www.linkedin.com/in/heitorperalles)

### Source

[github.com/heitorperalles/checkCpfRestApiGoBind](https://www.github.com/heitorperalles/checkCpfRestApiGoBind)

### License

MIT licensed. See the **LICENSE** file for details.
