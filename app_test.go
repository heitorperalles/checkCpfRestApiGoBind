package main

import (
		"log"
		"testing"
		"net/http"
    "net/http/httptest"
		"time"
		"bytes"
)

// Main TEST Function
//
// This function is executed before all tests
//
// Param according to testing package pattern [func TestMain(m *testing.M)]
func TestMain(m *testing.M) {

		log.Println("Starting TESTS...")
		m.Run()
}

// Test startRouting function
//
// This test consists in rise the API router and test Verify endpoint
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestStartRouting(t *testing.T) {

	log.Println("Testing startRouting...")

	// Calling startRouting asynchronously to make the test possible
	go startRouting()

	// Giving some time to the server start...
	time.Sleep(3 * time.Second)

	// Calling Verify Route with no data

	response, errRequest := http.Post("http://localhost:8000/api/v1/verify", "", nil)
	if (errRequest != nil) {
		t.Errorf("Problem calling Verify route: [%s]", errRequest)
	} else if response.StatusCode == http.StatusNotFound {
		t.Errorf("Problem calling Verify route: 404, Not Found!")
	}	else {
		log.Println("Success by calling [POST] Index: StatusCode ", response.StatusCode)
	}

	// TODO: startRouting function will keep running in background. Would be nice to stop it.
}

// Test validateCpf function
//
// This test consists of calling the function by passing a CPF without numbers
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfLetters(t *testing.T) {

	log.Println("Testing validateCpf passing a CPF without numbers...")

	err := validateCpf("INVALID")

	if err != http.StatusBadRequest {
		t.Errorf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a CPF without numbers.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing an empty CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfEmpty(t *testing.T) {

	log.Println("Testing validateCpf passing an empty CPF...")

	err := validateCpf("")

	if err != http.StatusBadRequest {
		t.Errorf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with an empty CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a malformed  CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfMalformed(t *testing.T) {

	log.Println("Testing validateCpf passing a malformed CPF...")

	err := validateCpf("1234567")

	if err != http.StatusBadRequest {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a malformed CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a canceled  CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfCanceled(t *testing.T) {

	log.Println("Testing validateCpf passing a canceled CPF...")

	err := validateCpf("64913872591")

	if err != http.StatusForbidden {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than Forbidden: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a canceled CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing an inexistant CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfInexistant(t *testing.T) {

	log.Println("Testing validateCpf passing an inexistant CPF...")

	err := validateCpf("11334739706")

	if err != http.StatusForbidden {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than Forbidden: %d", err)
	} else {
		log.Println("Success by calling validateCpf with an inexistant CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a regular CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfRegular(t *testing.T) {

	log.Println("Testing validateCpf passing a regular CPF...")

	err := validateCpf("40442820135")

	if err != http.StatusOK {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than OK: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a regular CPF.")
	}
}

// Test verify function
//
// This test consists of calling the function by passing an empty body
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestVerifyEmpty(t *testing.T) {

	log.Println("Testing verify function by passing an empty Request...")

	var jsonStr = []byte(`" "`)

	request, errRequest := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if (errRequest != nil) {
			t.Skipf("Problem creating Request for the test: %s", errRequest)
	} else {
			responseRecorder := httptest.NewRecorder()
			verify(responseRecorder, request)
			response := responseRecorder.Result()

			if (response.StatusCode != http.StatusBadRequest) {
				t.Errorf("Expecting BadRequest by calling Verify route with an empty Request, received: %d", response.StatusCode)
			} else {
				log.Println("Success by testing Verify Route with an empty Request")
			}
	}
}

// Test verify function
//
// This test consists of calling the function by not passing CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestVerifyNoCPF(t *testing.T) {

	log.Println("Testing verify function by not passing required field CPF...")

	var jsonStr = []byte(`{"name":"Heitor Peralles"}`)

	request, errRequest := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if (errRequest != nil) {
			t.Skipf("Problem creating Request for the test: %s", errRequest)
	} else {
			responseRecorder := httptest.NewRecorder()
			verify(responseRecorder, request)
			response := responseRecorder.Result()

			if (response.StatusCode != http.StatusBadRequest) {
				t.Errorf("Expecting BadRequest by calling Verify route with no CPF, received: %d", response.StatusCode)
			} else {
				log.Println("Success by testing Verify Route with no CPF")
			}
	}
}

// Test verify function
//
// This test consists of calling the function by passing full data
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestVerifyComplete(t *testing.T) {

	log.Println("Testing verify function by passing full data...")

	var jsonStr = []byte(`{"name":"Heitor Peralles", "cpf":"40442820135",
	"rg": { "number":"209921899", "issued":"May 20, 2020", "entity":"DETRAN" }}`)

	request, errRequest := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if (errRequest != nil) {
			t.Skipf("Problem creating Request for the test: %s", errRequest)
	} else {
			responseRecorder := httptest.NewRecorder()
			verify(responseRecorder, request)
			response := responseRecorder.Result()

			if (response.StatusCode != http.StatusOK) {

				// Not rising an error because this test depends on SERPRO API response

				t.Skipf("Expecting OK by calling Verify route with full data: %d", response.StatusCode)
			} else {
				log.Println("Success by testing Verify Route with full data")
			}
	}
}
