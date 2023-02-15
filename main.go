package main

import (
	"fmt"

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

	xCmd = &cobra.Command{
		Use:   "x",
		Short: "test",
		Run: func(cmd *cobra.Command, args []string) {
			// test
		},
	}

	// blank up --count={count}
	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Start containers",
		Long:  "Start 1 or more empty containers",
		Run: func(cmd *cobra.Command, args []string) {
			client := conn()
			// image
			if !imageExist(client) {
				fmt.Println("Unable to find image \"busybox:latest\", pulling...")
				pullImage(client)
			} else {
				fmt.Println("found existing \"busybox:latest\" image...")
			}

			// container (create)
			if !containerExist(client, CONTAINER_NAME) {
				fmt.Println("Unable to find container \"empty-container\", creating...")
				createContainer(client)
			}

			// container (start)
			if !containerRunning(client, CONTAINER_NAME) {
				fmt.Println("Starting \"empty-container\" container...")
				startContainer(client)
			} else {
				fmt.Println("\"empty-container\" is already running...")
			}
		},
	}

	// clean
	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Remove created images & container",
		Long:  "Remove [empty-container] container & [busybox:latest] images",
		Run: func(cmd *cobra.Command, args []string) {
			// removeContainer(client)
			// removeImage(client)
		},
	}

	// blank images
	listImgCmd = &cobra.Command{
		Use:   "images",
		Short: "List images",
		Long:  "List all the images",
		Run: func(cmd *cobra.Command, args []string) {
			client := conn()
			images := getImages(client)
			for _, image := range images {
				if len(image.RepoTags) != 0 {
					tag := image.RepoTags[0]
					fmt.Println(tag)
				}
			}
		},
	}

	// blank containers
	listContCmd = &cobra.Command{
		Use:   "containers",
		Short: "List containers",
		Long:  "List all the containers",
		Run: func(cmd *cobra.Command, args []string) {
			client := conn()
			containers := getContainers(client)
			for _, container := range containers {
				name := container.Names[0][1:]
				fmt.Println(name)
			}
		},
	}
)

func main() {
	upCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of containers to start, default to 1")
	rootCmd.AddCommand(xCmd, upCmd, listContCmd, listImgCmd)
	rootCmd.Execute()
}
