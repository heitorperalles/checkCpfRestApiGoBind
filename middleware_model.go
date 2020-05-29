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
