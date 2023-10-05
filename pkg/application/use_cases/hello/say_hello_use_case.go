package hello

import (
	"fmt"
)

func SayHelloUseCase(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
