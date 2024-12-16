package configs_test

import (
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeSuccess_WhenNewConfig(t *testing.T) {
	c := configs.NewConfig()
	assert.NotNil(t, c)
}
