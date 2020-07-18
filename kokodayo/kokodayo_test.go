package kokodayo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenHello(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(GenHello("hikalium"), HelloResponse{Message: "kokodayo hikalium!"})
}
