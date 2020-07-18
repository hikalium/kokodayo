package kokodayo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GenHello(name string) gin.H {
	return gin.H{"message": fmt.Sprintf("kokodayo %s!", name)}
}
