package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {

	expected := []struct {
		value string
		exp   string
		r     bool
	}{
		{
			"myapp_password",
			`^myapp*`,
			true,
		},
		{
			"xxx_password",
			`^myapp*|^xxx*`,
			true,
		},
		{
			"app_password",
			`^myapp*|^xxx*`,
			false,
		},
	}

	gcp := &GCPSecretManager{}

	for _, e := range expected {
		r := gcp.Filter(e.value, e.exp)
		assert.Equal(t, e.r, r)
	}
}
