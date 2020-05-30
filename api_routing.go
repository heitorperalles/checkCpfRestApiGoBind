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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index Route
//
// Params according to HandleFunc pattern: func(ResponseWriter, *Request))
func index(indexResponse http.ResponseWriter, indexRequest *http.Request) {
	fmt.Fprintf(indexResponse, "Check-CPF API by Heitor Peralles!")
}

// Routing service start function
func startRouting() {

	log.Print("Initializing Check-CPF API")

	// Init router
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api/v1/verify", verify).Methods("POST")

	// Start server
	port := ":8000"
	log.Print("Starting to listen on port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
