package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/yamatt/dwnldr/pkg/config"   // Example of importing your config package
	"github.com/yamatt/dwnldr/pkg/handlers" // Example of importing the downloader logic
)

var configPath string = "dwnldr.toml"
var outputDir string = "."

// Define the root command
var rootCmd = &cobra.Command{
    Use:   "dwnldr",
    Short: "A tool to download various sources",
    
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use --help.")
    },
}

// Define a subcommand for downloading from a URL
var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download a file from a specified URL",
    Run: func(cmd *cobra.Command, args []string) {
         // Load the config
        config, err := config.Load(configPath)
        if err != nil {
            fmt.Printf("Error loading config: %v\n", err)
            os.Exit(1)
        }

        // Print the config for debugging
        fmt.Printf("Using config: %s\n", configPath)

        // Example: Call your download logic
        if len(args) < 1 {
            fmt.Println("You must provide a URL to download.")
            os.Exit(1)
        }
        
        url := args[0]
        fmt.Printf("Downloading from URL: %s\n", url)

        // Call the downloader to handle the actual download
        err = handlers.Download(url, outputDir, *config)
        if err != nil {
            fmt.Printf("Error downloading file: %v\n", err)
            os.Exit(1)
        }
    },
}

// Function to set up the CLI and add subcommands
func init() {
    // Add the subcommands to the root command
    rootCmd.AddCommand(downloadCmd)

    downloadCmd.Flags().StringVarP(&configPath, "config", "c", "config/dwnldr.toml", "Path to config file")
    downloadCmd.Flags().StringVarP(&outputDir, "output-dir", "o", ".", "Output directory for the downloaded file")
}

// Execute the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
