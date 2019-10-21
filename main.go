package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/gorilla/mux"
)

func evalRPN(tokens []string) (int, error) {
	stk := stack.New()

	for _, token := range tokens {
		if token[0] == '-' && len(token) > 1 || token[0] >= '0' && token[0] <= '9' {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("failed to cast token to int: %v", err)
			}
			stk.Push(num)
			continue
		}

		right := stk.Pop()
		left := stk.Pop()

		if token == "+" {
			stk.Push(left.(int) + right.(int))
		}

		if token == "-" {
			stk.Push(left.(int) - right.(int))
		}

		if token == "*" {
			stk.Push(left.(int) * right.(int))
		}

		if token == "/" {
			stk.Push(left.(int) / right.(int))
		}
	}
	log.Printf("reach here")
	return stk.Peek().(int), nil
}

// RpnHandler get the reverse polish notaiton request from client
func RpnHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url inputs is %v", r.URL.Path[1:])
	rpnstr := r.URL.Path[1:]
	strs := strings.Split(rpnstr, ",")
	log.Printf("%q\n", strs)
	res, err := evalRPN(strs)
	if err != nil {
		log.Printf("erro: %v", err)
		fmt.Fprintf(w, "error = %v", err)
	}
	log.Printf("res = %v", res)

	rpnRes := RPN{
		RPN:    rpnstr,
		Result: res,
	}

	if err := json.NewEncoder(w).Encode(rpnRes); err != nil {
		fmt.Fprintf(w, "error = %v", err)
	}
	fmt.Fprintf(w, "res = %v", res)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{rpn}", RpnHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
