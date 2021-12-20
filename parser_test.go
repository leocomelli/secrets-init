package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONParser(t *testing.T) {

	expected := []*SecretData{
		&SecretData{
			Path:         "/project/123/secrets/mysecret",
			Name:         "mysecret",
			Data:         `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
			ContentKey:   "user",
			ContentValue: "myuser",
		},
		&SecretData{
			Path:         "/project/123/secrets/mysecret",
			Name:         "mysecret",
			Data:         `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
			ContentKey:   "password",
			ContentValue: "s3cr3t",
		},
		&SecretData{
			Path:         "/project/123/secrets/mysecret",
			Name:         "mysecret",
			Data:         `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
			ContentKey:   "host",
			ContentValue: "localhost:5432",
		},
	}

	s := &SecretData{
		Path: "/project/123/secrets/mysecret",
		Name: "mysecret",
		Data: `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
	}

	p := &JSONContentParser{}
	ss := p.Parse(s)

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], ss[i])
	}
}

func TestNoParser(t *testing.T) {

	expected := []*SecretData{
		&SecretData{
			Path:         "/project/123/secrets/mysecret",
			Name:         "mysecret",
			Data:         `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
			ContentKey:   "",
			ContentValue: "",
		},
	}

	s := &SecretData{
		Path: "/project/123/secrets/mysecret",
		Name: "mysecret",
		Data: `{"user": "myuser", "password": "s3cr3t", "host": "localhost:5432"}`,
	}

	p := &NoParser{}
	ss := p.Parse(s)

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], ss[i])
	}
}
