package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	count   = 1 // container count
	keepImg = 0 // keep busybox:latest image
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
			// image
			if !imageExist(client) {
				fmt.Println("🎣 Unable to find image \"busybox:latest\", pulling...")
				pullImage(client)
			}

			// container (create & start)
			var startIdx int
			if !containerExist(client) {
				startIdx = 1
			} else {
				// get the max container index name
				var max int
				for _, container := range getContainers(client) {
					name := container.Names[0][1:]
					if strings.Contains(name, CONTAINER_NAME) {
						nameSplit := strings.Split(name, "-")
						containerIdx, err := strconv.Atoi(nameSplit[len(nameSplit)-1])
						if err != nil {
							log.Fatal(err)
						}
						if containerIdx > max {
							max = containerIdx
						}
					}
				}
				startIdx = max + 1
			}
			for i := startIdx; i < count+startIdx; i++ {
				name := fmt.Sprintf("%s-%d", CONTAINER_NAME, i)
				fmt.Println(name)
				createContainer(client, name)
				startContainer(client, name)
			}
		},
	}

	// clean
	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Remove created images & container",
		Long:  "Remove [empty-container] container & [busybox:latest] images",
		Run: func(cmd *cobra.Command, args []string) {
			client := conn()
			removeContainer(client)
			if keepImg == 0 {
				removeImage(client)
			}
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
	cleanCmd.Flags().IntVarP(&keepImg, "keep", "k", 0, "Whether to keep busybox:latest image, default to false")
	rootCmd.AddCommand(upCmd, listContCmd, listImgCmd, cleanCmd)
	rootCmd.Execute()
}
