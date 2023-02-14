package main

import (
	"testing"
)

func TestContainerExist(t *testing.T) {
	client := conn()
	name := CONTAINER_NAME
	if !containerExist(client, name) {
		t.Errorf("Container %s does not exist", name)
	}
}

func TestCreateContainer(t *testing.T) {
	client := conn()
	createContainer(client)
}

func TestImageExist(t *testing.T) {
	client := conn()
	name := IMAGE_NAME
	if !imageExist(client) {
		t.Errorf("Image %s does not exist", name)
	}
}

func TestImagePull(t *testing.T) {
	client := conn()
	pullImage(client)
}
