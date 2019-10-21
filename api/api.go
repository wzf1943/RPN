package api

import (
	"RPN/model"
	"RPN/parser"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// RpnHandler get the reverse polish notaiton request from client
func RpnHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url inputs is %v", r.URL.Path[1:])
	rpnstr := r.URL.Path[1:]
	strs := strings.Split(rpnstr, ",")
	log.Printf("%q\n", strs)
	res, err := parser.EvalRPN(strs)
	if err != nil {
		log.Printf("erro: %v", err)
		fmt.Fprintf(w, "error = %v", err)
	}
	log.Printf("res = %v", res)

	rpnRes := model.RPN{
		RPN:    rpnstr,
		Result: res,
	}

	if err := json.NewEncoder(w).Encode(rpnRes); err != nil {
		fmt.Fprintf(w, "error = %v", err)
	}
	fmt.Fprintf(w, "res = %v", res)

}
