package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var CheckORSToken = checkORSToken

func TestCheckORSToken(t *testing.T) {
	os.Setenv("ORS_TOKEN", "abc")
	assert.NotPanics(t, CheckORSToken, " the code should not panic")

	os.Setenv("ORS_TOKEN", "")
	assert.Panics(t, CheckORSToken, "the code should panic")

}
