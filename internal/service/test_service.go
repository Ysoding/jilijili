package service

import "fmt"

type TestService struct {
}

func NewTestService() *TestService {
	fmt.Println("test service")
	return &TestService{}
}
