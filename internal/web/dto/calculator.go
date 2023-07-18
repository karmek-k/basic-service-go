package dto

type CalculatorArguments struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalculatorResult struct {
	Result int `json:"result"`
}