package main

import "testing"

func TestSayHello(t *testing.T) {
	want := []byte("hello jenkins")

	got := sayHello()

	if string(got) != string(want) {
		t.Error("Invalid string")
	}
}

func TestSayByeBye(t *testing.T) {
	want := []byte("bye bye")

	got := sayByeBye()

	if string(got) != string(want) {
		t.Error("Invalid string")
	}
}
