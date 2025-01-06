package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData_Success(t *testing.T) {
	h, err := GetData("data.json")

	assert.Nil(t, err)
	assert.NotNil(t, h)
}
