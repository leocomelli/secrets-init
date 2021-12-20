package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// ContentParser defines secret content parser behaviors
type ContentParser interface {
	Parse(*SecretData) []*SecretData
}

// JSONContenParser represents a JSON parser
type JSONContentParser struct {
	Tmpl string
}

// Parse lists each json entry as a secret
//
// Consider a secret called myscret with the content below
//
//     {
//        "user": "root",
//        "password": "s3cr3t",
//        "host" : "127.0.0.1:5432",
//     }
//
// The following secrets will be returned
//
// myscret_user: root
// mysecret_password: s3cr3t
// mysecret_host: 127.0.0.1:5432
//
func (j *JSONContentParser) Parse(s *SecretData) []*SecretData {
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(s.Data), &m); err != nil {
		log.Println("WARN: invalid json")
	}

	var secrets []*SecretData
	for k, v := range m {
		secrets = append(secrets, &SecretData{Name: s.Name, Path: s.Path, Data: s.Data, ContentKey: k, ContentValue: fmt.Sprintf("%v", v)})
	}

	return secrets
}

// NoParser represents no parser
type NoParser struct {
	Tmpl string
}

// Parse puts the secret data in a slice
func (n *NoParser) Parse(s *SecretData) []*SecretData {
	return []*SecretData{s}
}
