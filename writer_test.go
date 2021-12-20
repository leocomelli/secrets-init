package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteToStdin(t *testing.T) {
	var buf bytes.Buffer

	w, err := NewWriter(&buf, templates["plaintext"])
	assert.Nil(t, err)

	s := []*SecretData{
		&SecretData{
			Path: "secrets/mysecret",
			Name: "mysecret_user",
			Data: "root",
		},
		&SecretData{
			Path: "secrets/mysecret",
			Name: "mysecret_password",
			Data: "s3cr3t",
		},
	}

	_ = w.Write(s...)

	expected := `
export mysecret_user="root"
export mysecret_password="s3cr3t"`

	assert.Equal(t, expected, buf.String())
}

func TestWriteToFile(t *testing.T) {
	file, err := os.CreateTemp("", "test")

	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	w, err := NewWriter(file, templates["plaintext"])
	assert.Nil(t, err)

	s := []*SecretData{
		&SecretData{
			Path: "secrets/mysecret",
			Name: "mysecret_user",
			Data: "root",
		},
		&SecretData{
			Path: "secrets/mysecret",
			Name: "mysecret_password",
			Data: "s3cr3t",
		},
	}

	_ = w.Write(s...)

	content, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	expected := `
export mysecret_user="root"
export mysecret_password="s3cr3t"`

	assert.Equal(t, expected, string(content))
}
