package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Subject verify route
//
// Request: Subject struct
// Response: Veridict struct
//
// Params according to HandleFunc pattern: func(ResponseWriter, *Request))
func verify(verdictResponse http.ResponseWriter, subjectRequest *http.Request) {

	log.Print("Receiving a request...")

	// Treating Request...
	var subject Subject
	decoder := json.NewDecoder(subjectRequest.Body)
	errsubject := decoder.Decode(&subject)
	if (errsubject != nil){
			log.Println("Problem trying to decode received JSON:")
			log.Println(errsubject)

			// Composing Error Response

			var verdict Verdict
			verdict.Status = "False"
			verdict.Message = INVALID_JSON_FORMAT_MESSAGE
			verdictResponse.Header().Set("Content-Type", "application/json")
			verdictResponse.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(verdictResponse)
			errVerdictResponse := encoder.Encode(verdict)
			if (errVerdictResponse != nil){
					log.Println("Problem trying to write response JSON:")
					log.Println(errVerdictResponse)
			}
			return
	}

	// Checking received fields

	if (subject.Name != "") {
		log.Print("subject's Name: " + subject.Name)
	}
	if (subject.Cpf != "") {
		log.Print("subject's CPF: " + subject.Cpf)
	} else {
		log.Print("CPF not provided!")

		// Composing Error Response

		var verdict Verdict
		verdict.Status = "False"
		verdict.Message = CPF_NOT_PROVIDED_MESSAGE
		verdictResponse.Header().Set("Content-Type", "application/json")
		verdictResponse.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(verdictResponse)
		errVerdictResponse := encoder.Encode(verdict)
		if (errVerdictResponse != nil){
				log.Println("Problem trying to write response JSON:")
				log.Println(errVerdictResponse)
		}
		return
	}
	if (subject.RG != nil) {
		if (subject.RG.Number != "") {
			log.Print("subject's rg Number: " + subject.RG.Number)
		}
		if (subject.RG.Issued != "") {
			log.Print("subject's rg Issued: " + subject.RG.Issued)
		}
		if (subject.RG.Entity != "") {
			log.Print("subject's rg CVV: " + subject.RG.Entity)
		}
	}

	// Verifying CPF...

	cpfValidationCode := validateCpf(subject.Cpf)

	// Creating Response...

	var verdict Verdict
	switch cpfValidationCode {
		case http.StatusOK:
			verdict.Status = "True"
		case http.StatusForbidden:
			verdict.Status = "False"
			verdict.Message = SUBJECT_REJECTED_MESSAGE
		case http.StatusBadRequest:
			verdict.Status = "False"
			verdict.Message = INVALID_CPF_FORMAT_MESSAGE
		case http.StatusInternalServerError:
			verdict.Status = "False"
			verdict.Message = EXTERNAL_SERVER_ERROR_MESSAGE
		default:
			verdict.Status = "False"
			verdict.Message = UNKNOWN_ERROR_MESSAGE
	}

	verdictResponse.Header().Set("Content-Type", "application/json")
	verdictResponse.WriteHeader(cpfValidationCode)
	encoder := json.NewEncoder(verdictResponse)
	errVerdictResponse := encoder.Encode(verdict)
	if (errVerdictResponse != nil){
			log.Println("Problem trying to write response JSON:")
			log.Println(errVerdictResponse)
	}
}
