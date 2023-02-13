package main

import "testing"

func TestContainerExist(t *testing.T) {
	client := conn()
	name := "empty-container"
	if !containerExist(client, name) {
		t.Errorf("Container %s does not exist", name)
	}
}

func TestImageExist(t *testing.T) {
	client := conn()
	name := "empty-container:latest"
	if !imageExist(client, name) {
		t.Errorf("Image %s does not exist", name)
	}
}
