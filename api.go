package main

import (
	"bytes"
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func conn() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func getImages(client *docker.Client) []docker.APIImages {
	images, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		log.Fatal(err)
	}
	return images
}

func getContainers(client *docker.Client) []docker.APIContainers {
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}
	return containers
}

func createContainer(client *docker.Client) {
	config := docker.Config{
		Image: IMAGE_NAME,
		Cmd:   []string{"tail", "-f", "/dev/null"},
	}
	opts := docker.CreateContainerOptions{
		Name:   CONTAINER_NAME,
		Config: &config,
	}
	_, err := client.CreateContainer(opts)
	if err != nil {
		log.Fatal("🍥: ", err)
	}
}

func startContainer(client *docker.Client) {
	err := client.StartContainer(CONTAINER_NAME, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func containerRunning(client *docker.Client, name string) bool {
	containers := getContainers(client)
	for _, container := range containers {
		if container.Names[0][1:] == name {
			status := container.State
			return status == "running"
		}
	}
	return false
}

func containerExist(client *docker.Client, name string) bool {
	containers := getContainers(client)
	containerNames := []string{}
	for _, container := range containers {
		containerNames = append(containerNames, container.Names[0][1:])
	}
	return contains(containerNames, name)
}

func imageExist(client *docker.Client) bool {
	images := getImages(client)
	imgNames := []string{}
	for _, image := range images {
		if len(image.RepoTags) > 0 {
			imgNames = append(imgNames, image.RepoTags[0])
		}
	}
	return contains(imgNames, IMAGE_NAME)
}

func pullImage(client *docker.Client) {
	var buf bytes.Buffer
	err := client.PullImage(docker.PullImageOptions{Repository: IMAGE_NAME, OutputStream: &buf}, docker.AuthConfiguration{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
