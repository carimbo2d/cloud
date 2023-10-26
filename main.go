package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
)

type ArithService struct{}

func (s *ArithService) Add(a, b float64) float64 {
	return a + b
}

func (s *ArithService) Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func (s *ArithService) Scalar(vector [3]float64, scalar float64) []float64 {
	return []float64{vector[0] * scalar, vector[1] * scalar, vector[2] * scalar}
}

type HelperService struct{}

type ObjectSampleResponse struct {
	Number float64 `json:"number"`
	String string  `json:"string"`
}

type SampleResponse struct {
	Integer int    `json:"integer"`
	Boolean bool   `json:"boolean"`
	String  string `json:"string"`
	Null    *int   `json:"null"`
	// Array   []int                `json:"array"`
	Object ObjectSampleResponse `json:"object"`
}

func (s *HelperService) Sample() SampleResponse {
	return SampleResponse{
		Integer: 1,
		Boolean: true,
		String:  "hello",
		Null:    nil,
		// Array:   []int{1, 2, 3},
		Object: ObjectSampleResponse{
			Number: 1.23,
			String: "world",
		},
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var server = rpc.NewServer()

	if err := server.RegisterName("arith", &ArithService{}); err != nil {
		return err
	}

	if err := server.RegisterName("helper", &HelperService{}); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server)
}
