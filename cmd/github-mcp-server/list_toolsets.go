package main

import (
	"fmt"
	"strings"

	"github.com/github/github-mcp-server/pkg/github"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/spf13/cobra"
)

var listToolsetsCmd = &cobra.Command{
	Use:   "list-toolsets",
	Short: "List available toolsets",
	Long:  `Display all available toolsets with their IDs, descriptions, and default status.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		return listToolsets()
	},
}

func init() {
	rootCmd.AddCommand(listToolsetsCmd)
}

func listToolsets() error {
	// Create translation helper
	t, _ := translations.TranslationHelper()

	// Build inventory with all toolsets
	r := github.NewInventory(t).Build()

	// Get default toolsets
	defaultIDs := r.DefaultToolsetIDs()
	defaultSet := make(map[string]bool)
	for _, id := range defaultIDs {
		defaultSet[string(id)] = true
	}

	// Print header
	fmt.Println("Available Toolsets:")
	fmt.Println()

	// Print default toolsets section
	fmt.Println("Default Toolsets:")
	fmt.Println(strings.Repeat("-", 80))
	for _, id := range defaultIDs {
		desc := r.ToolsetDescriptions()[id]
		fmt.Printf("  %-25s %s\n", string(id), desc)
	}
	fmt.Println()

	// Print all available toolsets
	fmt.Println("All Available Toolsets:")
	fmt.Println(strings.Repeat("-", 80))

	// Get all available toolsets (excludes special ones like "dynamic")
	allToolsets := r.AvailableToolsets("dynamic")
	for _, ts := range allToolsets {
		isDefault := ""
		if defaultSet[string(ts.ID)] {
			isDefault = " (default)"
		}
		fmt.Printf("  %-25s %s%s\n", string(ts.ID), ts.Description, isDefault)
	}
	fmt.Println()

	// Print usage examples
	fmt.Println("Usage Examples:")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("  # Use default toolsets")
	fmt.Println("  github-mcp-server stdio")
	fmt.Println()
	fmt.Println("  # Use specific toolsets")
	fmt.Println("  github-mcp-server stdio --toolsets=actions,gists,notifications")
	fmt.Println()
	fmt.Println("  # Use default + additional toolsets")
	fmt.Println("  github-mcp-server stdio --toolsets=default,actions,gists")
	fmt.Println()
	fmt.Println("  # Enable all toolsets")
	fmt.Println("  github-mcp-server stdio --toolsets=all")

	return nil
}
