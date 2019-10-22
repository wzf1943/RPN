package api

import (
	"RPN/model"
	"RPN/parser"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// RpnHandler get the reverse polish notaiton request from client
func RpnHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(string(body))

	var inputs model.RPNInputs
	err = json.Unmarshal(body, &inputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("input %v", inputs)

	var resRPN model.RPNs
	for _, input := range inputs.Inputs {
		strs := strings.Split(input.Input, " ")
		log.Printf("%q\n", strs)
		res, err := parser.EvalRPN(strs)
		if err != nil {
			log.Printf("erro: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("res = %v", res)

		rpn := model.RPN{
			Name:   input.Input,
			Result: res,
		}
		resRPN.RPN = append(resRPN.RPN, rpn)
	}

	if err := json.NewEncoder(w).Encode(resRPN); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
