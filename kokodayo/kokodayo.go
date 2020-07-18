package kokodayo

import (
	"fmt"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func GenHello(name string) HelloResponse {
	return HelloResponse{Message: fmt.Sprintf("kokodayo %s!", name)}
}
