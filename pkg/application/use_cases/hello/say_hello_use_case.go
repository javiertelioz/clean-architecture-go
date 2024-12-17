package hello

import "fmt"

type SayHelloUseCase struct{}

func NewSayHelloUseCase() *SayHelloUseCase {
	return &SayHelloUseCase{}
}

func (s *SayHelloUseCase) Execute(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
