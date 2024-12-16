package adapter_test

import (
	"context"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/adapter"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeSuccess_WhenInstanceNewCoordinateServiceAdapter(t *testing.T) {
	r := adapter.NewCoordinateServiceAdapter(configs.NewConfig())
	c, e := r.GetByCep(context.Background(), "09130220")

	assert.NotNil(t, r)
	assert.Nil(t, e)
	assert.NotNil(t, c)

}
