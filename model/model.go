package model

type RPN struct {
	Name   string `json:"name"`
	Result int    `json:"result"`
}

type RPNs struct {
	RPN []RPN `json:"rpn"`
}

type RPNInput struct {
	Input string `json:"input"`
}

type RPNInputs struct {
	Inputs []RPNInput `json:"rpns"`
}

type Health struct {
	Status     string `json:"status"`
	StatusCode int    `json:"code"`
}
