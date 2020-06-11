package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConsumption(t *testing.T) {
	consumption, err := GetConsumption("small-diesel-car")

	assert.NoError(t, err, "valid case - small-diesel-car & no error")
	assert.Equal(t, consumption, 142)

	consumption, err = GetConsumption("train")

	assert.NoError(t, err, "valid case -train & no error")
	assert.Equal(t, consumption, 6)

	consumption, err = GetConsumption("myway")

	assert.EqualError(t, err, "transportation method myway does not exist", "invalid case - myway & error")

}
