package kokodayo

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenHello(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(GenHello("hikalium"), gin.H{"message": "kokodayo hikalium!"})
}
