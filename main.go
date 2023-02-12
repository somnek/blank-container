package main

import (
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/spf13/cobra"
)

var (
	count = 1 // container count
)

var (
	rootCmd = &cobra.Command{
		Use:   "blank",
		Short: "Fastest way to spin up empty container 🧩",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// blank up --count={count}
	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Start containers",
		Long:  "Start 1 or more empty containers",
		Run: func(cmd *cobra.Command, args []string) {
			client := conn()

			images := getImages(client)
			for _, image := range images {
				fmt.Println(image.RepoTags)
			}

			containers := getContainers(client)
			for _, container := range containers {
				fmt.Println(container.Names)
			}

		},
	}
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

func main() {
	upCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of containers to start, default to 1")
	rootCmd.AddCommand(upCmd)
	rootCmd.Execute()
}
