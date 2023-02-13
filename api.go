package main

import (
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
	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		log.Fatal(err)
	}
	return containers
}

func startContainer(client *docker.Client, id string) {
	if err := client.StartContainer(id, nil); err != nil {
		log.Fatal(err)
	}
}

func containerExist(client *docker.Client, name string) bool {
	containers := getContainers(client)
	containerNames := []string{}
	for _, container := range containers {
		containerNames = append(containerNames, container.Names[0][1:])
	}
	return contains(containerNames, name)
}

func imageExist(client *docker.Client, name string) bool {
	images := getImages(client)
	imgNames := []string{}
	for _, image := range images {
		if len(image.RepoTags) > 0 {
			imgNames = append(imgNames, image.RepoTags[0])
		}
	}
	return contains(imgNames, name)
}
