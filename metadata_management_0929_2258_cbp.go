// 代码生成时间: 2025-09-29 22:58:45
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

// Metadata represents the structure for metadata
type Metadata struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Description string `json:"description"`
}

// metadataStorage is a mock storage for metadata
var metadataStorage map[string]Metadata

func init() {
    metadataStorage = make(map[string]Metadata)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add [name] [version] [description]",
    Short: "Add a new metadata entry",
    Long:  `Add a new metadata entry to the system`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 3 {
            cmd.Help()
            return
        }
        name := args[0]
        version := args[1]
        description := args[2]
        
        if _, exists := metadataStorage[name]; exists {
            fmt.Println("Metadata with the same name already exists")
            return
        }

        meta := Metadata{Name: name, Version: version, Description: description}
        metadataStorage[name] = meta
        fmt.Printf("Metadata '%s' added successfully
", name)
    },
}

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all metadata entries",
    Long:  `List all metadata entries in the system`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(metadataStorage) == 0 {
            fmt.Println("No metadata entries found")
            return
        }
        for key, value := range metadataStorage {
            fmt.Printf("Name: %s, Version: %s, Description: %s
", value.Name, value.Version, value.Description)
        }
    },
}

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
    Use:   "remove [name]",
    Short: "Remove an existing metadata entry",
    Long:  `Remove an existing metadata entry from the system`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 1 {
            cmd.Help()
            return
        }
        name := args[0]
        if _, exists := metadataStorage[name]; !exists {
            fmt.Println("Metadata with the name not found")
            return
        }
        delete(metadataStorage, name)
        fmt.Printf("Metadata '%s' removed successfully
", name)
    },
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "metadata",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
metadata [command] [flags] [args]`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    // Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Add commands to the rootCmd
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(removeCmd)

    // Execute adds all child commands to the root command and sets up
    // the Cobra middleware
    Execute()
}
