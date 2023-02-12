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

	// blank up --count={count}
	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Start containers",
		Long:  "Start 1 or more empty containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("place holder...")
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
	rootCmd.AddCommand(upCmd, listContCmd, listImgCmd)
	rootCmd.Execute()
}
