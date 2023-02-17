package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

func conn() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func removeContainer(client *docker.Client) {
	containerNames := []string{}
	containerIds := []string{}

	for _, container := range getContainers(client) {
		name := container.Names[0][1:]
		id := container.ID
		if strings.Contains(name, CONTAINER_NAME) {
			containerIds = append(containerIds, id)
		}
		containerNames = append(containerNames, name)
	}

	if len(containerNames) == 0 {
		fmt.Println("No container to remove, skipping...")
		return
	}

	for i, id := range containerIds {
		opts := docker.RemoveContainerOptions{ID: id, Force: true}
		if err := client.RemoveContainer(opts); err != nil {
			log.Fatal("Error when trying to remove container", err)
		}
		fmt.Printf("🧹 Container removed... [%s]\n", containerNames[i])
	}
}

func removeImage(client *docker.Client) {
	imageNames := []string{}
	for _, image := range getImages(client) {
		imageNames = append(imageNames, image.RepoTags...)
	}

	if !contains(imageNames, IMAGE_NAME) {
		fmt.Println("No image to remove, skipping...")
		return
	}

	if err := client.RemoveImage(IMAGE_NAME); err != nil {
		log.Fatal("Error when trying to remove image", err)
	}
	fmt.Println("🧹 Image removed...")
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

func createContainer(client *docker.Client, name string) {
	config := docker.Config{
		Image: IMAGE_NAME,
		Cmd:   []string{"tail", "-f", "/dev/null"},
	}
	opts := docker.CreateContainerOptions{
		Name:   name,
		Config: &config,
	}
	_, err := client.CreateContainer(opts)
	if err != nil {
		log.Fatal("🍥: ", err)
	}
}

func startContainer(client *docker.Client, name string) {
	err := client.StartContainer(name, nil)
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
