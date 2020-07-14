package kokodayo

import (
	"encoding/json"
	"fmt"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func GenHello(name string) string {
	resBytes, _ := json.Marshal(HelloResponse{Message: fmt.Sprintf("kokodayo %s!", name)})
	return string(resBytes)
}
