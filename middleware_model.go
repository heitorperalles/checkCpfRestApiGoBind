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

// SERPRO model
//
// Represents the RESPONSE format of SERPRO API.
//
// Response example:
//
// 		{
// 		"ni":"40442820135",
// 		"nome":"Nome do CPF 404.428.201-35",
// 		"situacao":
//   		{
//     		"codigo":"0",
//     		"descricao":"Regular"
//   		}
// 		}
//
type SerproPerson struct {

	// CPF Identification Number (The CPF)
	//
	// Required: False
	// Example: 40442820135
	NI   		string  						`json:"ni"`

	// CPF Name
	//
	// Required: False
	// Example: Heitor Peralles
	Name    string  						`json:"nome"`

	// CPF Status
	//
	// Required: False
	// Example: [SerproPersonStatus struct]
	Status *SerproPersonStatus  `json:"situacao"`
}

// SERPRO CPF status
//
// Represents the CPF block on SERPRO API Response.
type SerproPersonStatus struct {

	// CPF Status Code
	//
	// One from:
	// 	0: 	Regular
	// 	2: 	Suspensa (Suspended)
	// 	3: 	Titular Falecido (Deceased Holder)
	// 	4: 	Pendente de Regularização (Regularization Pending)
	// 	5: 	Cancelada por Multiplicidade (Canceled by Multiplicity)
	// 	8: 	Nula (Null)
	// 	9: 	Cancelada de Ofício (Registration Canceled)
	//
	// Required: False
	// Example: 0
	Code 			 		string `json:"codigo"`

	// CPF Status Description
	//
	// One from:
	// 	Regular
	// 	Suspensa
	// 	Titular Falecido
	// 	Pendente de Regularização
	// 	Cancelada por Multiplicidade
	// 	Nula
	// 	Cancelada de Ofício
	//
	// Required: False
	// Example: Regular
	Description   string `json:"descricao"`
}
