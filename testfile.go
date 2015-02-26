package main

import "errors"

type TestStruct struct {
	Field1     int    `json:"field_1" sql:"test"`
	TestField2 string `json:"field_2"`
	NewField   string
}

type TestStruct2 struct {
	Field1     int    `json:"field_1"`
	TestField2 string `json:"field_2"`
	Embed
}

type Embed struct {
	EmbedField int
}

type Item struct {
	A string
	B int
}

var ErrNoSuchItem = errors.New("no such item")
