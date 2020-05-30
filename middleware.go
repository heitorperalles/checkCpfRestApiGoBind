//------------------------------------------------------------------------------
// From http://github.com/heitorperalles/checkCpfRestApiGoBind
//
// Distributed under The MIT License (MIT) <http://opensource.org/licenses/MIT>
//
// Copyright (c) 2020 Heitor Peralles <heitorgp@gmail.com>
//
// Permission is hereby  granted, free of charge, to any  person obtaining a copy
// of this software and associated  documentation files (the "Software"), to deal
// in the Software  without restriction, including without  limitation the rights
// to  use, copy,  modify, merge,  publish, distribute,  sublicense, and/or  sell
// copies  of  the Software,  and  to  permit persons  to  whom  the Software  is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE  IS PROVIDED "AS  IS", WITHOUT WARRANTY  OF ANY KIND,  EXPRESS OR
// IMPLIED,  INCLUDING BUT  NOT  LIMITED TO  THE  WARRANTIES OF  MERCHANTABILITY,
// FITNESS FOR  A PARTICULAR PURPOSE AND  NONINFRINGEMENT. IN NO EVENT  SHALL THE
// AUTHORS  OR COPYRIGHT  HOLDERS  BE  LIABLE FOR  ANY  CLAIM,  DAMAGES OR  OTHER
// LIABILITY, WHETHER IN AN ACTION OF  CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE  OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//------------------------------------------------------------------------------
package main

import (
	"log"
	"encoding/json"
	"regexp"
	"bytes"
	"io"
	"net/http"

	curl "github.com/andelf/go-curl"
)

// SERPRO API information
//
// URL and Token to validate CPF on SERPRO API
const (
		// Public URL to request CPF status on SERPRO API
    SERPRO_URL 			string = "https://apigateway.serpro.gov.br/consulta-cpf-df-trial/v1/cpf/"

		// Token to be used during requests on SERPRO API
	  SERPRO_TOKEN		string = "4e1a1858bdd584fdc077fb7d80f39283"

		// TODO Read these configurations from a config file
)

// CPF pre-processing function
//
// Param: Cpf
// Return:
//		Treated CPF (empty if invalid)
func treatCpf(cpf string) string {

	log.Print("Verifying CPF " + cpf)

	// Removing non-numbers...

	regex, errRegex := regexp.Compile("[^0-9]+")
	if errRegex != nil {
			log.Print("Couldn't create regex to verify CPF")
			log.Fatal(errRegex)
	}
	treatedCpf := regex.ReplaceAllString(cpf, "")

	if treatedCpf == "" {
		log.Print("Invalid CPF format: " + cpf)
	} else {
		log.Print("Post-processed CPF: " + treatedCpf)
	}

	return treatedCpf
}

// Function to convert HTTP code of SERPRO response
//
// Param: code
// Return:
//		http.StatusOK (CPF OK)
//		http.StatusBadRequest (Invalid CPF format)
//		http.StatusForbidden (CPF not regular or not existant)
//		http.StatusInternalServerError (Communication problem)
func convertHttpCode(code int) int {
	switch code {
		case 200:
				log.Println("[SERPRO] Status code 200: Request has been succeeded")
		case 206:
				log.Println("[SERPRO] Status code 206: Warning, Partial content returned")
		case 400:
				log.Println("[SERPRO] Status code 400: Invalid CPF format")
				return http.StatusBadRequest
		case 401:
				log.Println("[SERPRO] Status code 401: Unauthorized, please review the app TOKEN")
				return http.StatusInternalServerError
		case 404:
				log.Println("[SERPRO] Status code 404: Not existant CPF")
				return http.StatusForbidden
		case 500:
				log.Println("[SERPRO] Status code 500: Internal Server error")
				return http.StatusInternalServerError
		default:
				log.Println("[SERPRO] Unknown Status code:", code)
				return http.StatusInternalServerError
	}
	return http.StatusOK
}

// Function to treat received JSON on SERPRO response
//
// Param: body
// Return:
//		http.StatusOK (CPF OK)
//		http.StatusForbidden (CPF not regular or not existant)
//		http.StatusInternalServerError (Communication problem)
func treatResponseData(body io.Reader) int {
	var person SerproPerson
	decoder := json.NewDecoder(body)
	errSerproPerson := decoder.Decode(&person)
	if (errSerproPerson != nil){
			log.Println("[SERPRO] Problem trying to decode received JSON from SERPRO:")
			log.Println(errSerproPerson)
			return http.StatusInternalServerError
	}
	if (person.NI != "") {
		log.Print("[SERPRO] CPF: " + person.NI)
	}
	if (person.Name != "") {
		log.Print("[SERPRO] CPF Name: " + person.Name)
	}
	if (person.Status != nil) {

		// Status Codes and Descriptions

		if (person.Status.Code != "") {
			log.Print("[SERPRO] CPF Status Code: " + person.Status.Code)
		}
		if (person.Status.Description != "") {
			log.Print("[SERPRO] CPF Status Description: " + person.Status.Description)
		}
		if (person.Status.Code != "0") {
			return http.StatusForbidden
		}
	}	else {
		log.Print("[SERPRO] CPF Status not provided!")
		return http.StatusForbidden
	}
	return http.StatusOK
}

// Function to validate CPF
//
// Param: Cpf
// Return:
//		http.StatusOK (CPF OK)
//		http.StatusBadRequest (Invalid CPF format)
//		http.StatusForbidden (CPF not regular or not existant)
//		http.StatusInternalServerError (Communication problem)
func validateCpf(cpf string) int {

	treatedCpf := treatCpf(cpf)
	if treatedCpf == "" {
		return http.StatusBadRequest
	}

	log.Print("[CURL] Creating Request...")

	// Setting Curl properties

	curlInstance := curl.EasyInit()
	defer curlInstance.Cleanup()
	curlInstance.Setopt(curl.OPT_HTTPAUTH, curl.AUTH_ANY)
	curlInstance.Setopt(curl.OPT_DNS_CACHE_TIMEOUT, -1)
	curlInstance.Setopt(curl.OPT_SSL_VERIFYHOST, 0)
	curlInstance.Setopt(curl.OPT_SSL_VERIFYPEER, 0)

	// Setting URL and Authentication TOKEN

	curlInstance.Setopt(curl.OPT_URL, SERPRO_URL+treatedCpf)
	var headers []string
	headers =	append(headers, "Authorization: Bearer " + SERPRO_TOKEN)
	curlInstance.Setopt(curl.OPT_HTTPHEADER,headers)

	// Setting Function to get the returned Body

	var curlOutput []byte
	writeFunc := func (buf []byte, userdata interface{}) bool {
			log.Println("[CURL] Received bytes: ", len(buf))
			log.Println("[CURL] Received content: ", string(buf))
			curlOutput = append(curlOutput, buf...)
			return true
	}
	curlInstance.Setopt(curl.OPT_WRITEFUNCTION, writeFunc)

	// Actually calling the URL

	if err := curlInstance.Perform(); err != nil {
			log.Println("[CURL] Problem performing REQUEST: ", err)
			return http.StatusInternalServerError
	}

	// Treating Response HTTP Code

	curlCode,err := curlInstance.Getinfo(curl.INFO_RESPONSE_CODE)
	if err != nil {
		log.Println("[CURL] Problem obtaining HTTP code: ", err)
		return http.StatusInternalServerError
	}
	log.Println("[CURL] HTTP Response code: ", curlCode)
	if code := convertHttpCode(curlCode.(int)); code != http.StatusOK {
		return code
	}

	// Treating Response JSON data

	if body := bytes.NewReader(curlOutput); body != nil {
		return treatResponseData(body)
	}
	log.Println("Problem converting Body string to buffer!")
	return http.StatusInternalServerError
}
