package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

func Test_ToPointer(t *testing.T) {
	name := "test"
	assert.Equal(t, name, *utils.ToPointer(name))

	age := 33
	assert.Equal(t, age, *utils.ToPointer(age))

	alive := true
	assert.Equal(t, alive, *utils.ToPointer(alive))

	height := 172.7
	assert.Equal(t, height, *utils.ToPointer(height))

	assert.Panics(t, func() {
		pointer := utils.ToPointer("hi")
		utils.ToPointer(pointer)
	})
}
