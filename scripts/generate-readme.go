package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type ProjectStructure struct {
	Path        string
	Description string
	Depth       int
	Children    []*ProjectStructure
}

func generateReadme(root *ProjectStructure) string {
	var sb strings.Builder

	// Project title and description
	sb.WriteString("# Baby Tracker\n\n")
	sb.WriteString("## Project Overview\n\n")
	sb.WriteString("Baby Tracker is a comprehensive baby monitoring application integrating Home Assistant, Go backend, HTMX frontend, and local LLM technologies.\n\n")

	// Table of Contents
	sb.WriteString("## Project Structure\n\n")

	// Recursive function to generate project structure
	var writeStructure func(node *ProjectStructure)
	writeStructure = func(node *ProjectStructure) {
		// Skip root if it's empty
		if node.Path != "" {
			description := node.Description
			if description == "" {
				description = "No description"
			}

			// Use depth for indentation
			indent := strings.Repeat("  ", node.Depth)

			sb.WriteString(fmt.Sprintf("%s- `%s/`: %s\n",
				indent,
				filepath.Base(node.Path),
				description,
			))
		}

		// Sort children to ensure consistent output
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].Path < node.Children[j].Path
		})

		// Recursively write children
		for _, child := range node.Children {
			writeStructure(child)
		}
	}

	writeStructure(root)

	// Add sections for key components
	sb.WriteString("\n## Key Components\n\n")

	componentDescriptions := map[string]string{
		"backend":       "Go-based web application backend with database and LLM integration",
		"frontend":      "HTMX-powered web interface for user interactions",
		"homeassistant": "Custom Home Assistant integration for system-wide monitoring",
		"docs":          "Comprehensive project documentation",
		"scripts":       "Utility scripts for setup and deployment",
	}

	for component, description := range componentDescriptions {
		sb.WriteString(fmt.Sprintf("### %s\n%s\n\n",
			strings.Title(component),
			description,
		))
	}

	// Development and Setup
	sb.WriteString("## Development\n\n")
	sb.WriteString("### Prerequisites\n")
	sb.WriteString("- Go 1.24+\n")
	sb.WriteString("- Home Assistant\n")
	sb.WriteString("- PostgreSQL\n")
	sb.WriteString("- Docker (optional)\n\n")

	sb.WriteString("### Setup\n")
	sb.WriteString("1. Clone the repository\n")
	sb.WriteString("2. Copy `.env.example` to `.env` and configure\n")
	sb.WriteString("3. Run `./scripts/setup.sh`\n\n")

	// License and Contribution
	sb.WriteString("## License\n")
	sb.WriteString("MIT License\n\n")
	sb.WriteString("## Contributing\n")
	sb.WriteString("Please read `docs/DEVELOPMENT.md` for details on our code of conduct and the process for submitting pull requests.\n")

	return sb.String()
}

func crawlDirectory(rootPath string) *ProjectStructure {
	root := &ProjectStructure{
		Path:  rootPath,
		Depth: 0,
	}

	// Custom descriptions for specific directories
	descriptions := map[string]string{
		"backend":           "Go Web Application Backend",
		"frontend":          "HTMX Web Frontend",
		"homeassistant":     "Home Assistant Integration",
		"docs":              "Project Documentation",
		"scripts":           "Utility Scripts",
		".github":           "GitHub CI/CD Configurations",
		"cmd":               "Application entry points",
		"internal":          "Internal packages and core logic",
		"custom_components": "Custom Home Assistant components",
		"static":            "Static web assets",
		"templates":         "HTML templates",
		"workflows":         "GitHub Actions workflow definitions",
		"database":          "Database access layer",
		"llm":               "LLM integration and functionalities",
		"models":            "Data models and definitions",
		"services":          "Business logic and services",
		"migrations":        "Database migration scripts",
		"baby_tracker":      "Custom Home Assistant component for baby tracking",
		"babyTracker":       "Root directory of the Baby Tracker project", // Add this one
	}

	// Keep track of processed directories to avoid duplicates
	processedDirs := make(map[string]bool)

	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip some directories and files
		if path == rootPath ||
			strings.Contains(path, ".git") ||
			strings.Contains(path, "node_modules") ||
			strings.Contains(path, "vendor") ||
			processedDirs[path] {
			return nil
		}

		// Mark as processed
		processedDirs[path] = true

		// Calculate relative path and depth
		relativePath, _ := filepath.Rel(rootPath, path)
		depth := strings.Count(relativePath, string(filepath.Separator))

		// Only process directories
		if info.IsDir() {
			dirName := filepath.Base(path)
			description := descriptions[dirName]

			// Create node
			node := &ProjectStructure{
				Path:        path,
				Description: description,
				Depth:       depth,
			}

			// Find or create parent
			parent := root
			if depth > 0 {
				// Split the relative path
				parts := strings.Split(relativePath, string(filepath.Separator))

				// Traverse to find or create parent
				for i := 0; i < len(parts)-1; i++ {
					found := false
					for _, existingChild := range parent.Children {
						if filepath.Base(existingChild.Path) == parts[i] {
							parent = existingChild
							found = true
							break
						}
					}

					// If no matching parent found, create one
					if !found {
						newParent := &ProjectStructure{
							Path:     filepath.Join(rootPath, strings.Join(parts[:i+1], string(filepath.Separator))),
							Depth:    i + 1,
							Children: []*ProjectStructure{},
						}
						parent.Children = append(parent.Children, newParent)
						parent = newParent
					}
				}
			}

			// Add to parent's children
			parent.Children = append(parent.Children, node)
		}

		return nil
	})

	return root
}

func main() {
	// Get the current directory
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Generate project structure
	projectStructure := crawlDirectory(rootDir)

	// Generate README content
	readmeContent := generateReadme(projectStructure)

	// Write README
	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error writing README:", err)
		return
	}

	fmt.Println("README.md generated successfully!")
}
