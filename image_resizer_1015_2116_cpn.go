// 代码生成时间: 2025-10-15 21:16:11
Usage:
	image-resizer -d=<directory> -w=<width> -h=<height>
# TODO: 优化性能

Flags:
	-d=<directory>  Directory containing images to resize
	-w=<width>      Width to resize images to (in pixels)
	-h=<height>     Height to resize images to (in pixels)
*/

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Define constants for flags
const (
	flagDir = "d"
	flagWidth = "w"
	flagHeight = "h"
# 增强安全性
)

// Define the main command
# FIXME: 处理边界情况
var rootCmd = &cobra.Command{
	Use:   "image-resizer",
	Short: "Resize multiple images in a directory",
	Long:  `A command line tool to resize multiple images in a directory.`,
	Run:   resizeImages,
}

// Define the flags
var dir, width, height string

func init() {
# TODO: 优化性能
	rootCmd.PersistentFlags().StringVarP(&dir, flagDir, "d", "", "Directory containing images to resize")
	rootCmd.PersistentFlags().StringVarP(&width, flagWidth, "w", "", "Width to resize images to (in pixels)")
	rootCmd.PersistentFlags().StringVarP(&height, flagHeight, "h", "", "Height to resize images to (in pixels)")
}
# FIXME: 处理边界情况

// Main function to execute the CLI application
func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
# 优化算法效率

// Function to resize images in a directory
func resizeImages(cmd *cobra.Command, args []string) {
	// Check if directory path is provided
	if dir == "" {
		cmd.Help()
# NOTE: 重要实现细节
		os.Exit(1)
	}

	// Check if width and height are provided
	if width == "" || height == "" {
		cmd.Help()
		os.Exit(1)
	}

	// Convert width and height to integers
	w, err := strconv.Atoi(width)
	if err != nil {
		fmt.Println("Invalid width value. Please provide an integer.")
		os.Exit(1)
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println("Invalid height value. Please provide an integer.")
		os.Exit(1)
# 扩展功能模块
	}

	// Walk through the directory and process each image file
	error := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a file and has a valid image extension
		if d.Type().IsRegular() && (strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".png")) {
			img, err := loadImage(path)
			if err != nil {
				return err
# NOTE: 重要实现细节
			}

			// Resize the image
			resizedImg := resizeImage(img, w, h)

			// Save the resized image
			err = saveImage(resizedImg, path, w, h)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if error != nil {
		fmt.Printf("An error occurred while processing images: %v\
", error)
		os.Exit(1)
	}
}

// Function to load an image from a file path
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
# 优化算法效率

	return img, nil
# 改进用户体验
}

// Function to resize an image to the specified width and height
func resizeImage(img image.Image, width, height int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	r := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	draw.Draw(m, m.Bounds(), r, image.Point{0, 0}, draw.Src)

	return m
}

// Function to save a resized image to a file
# 优化算法效率
func saveImage(img image.Image, path string, width, height int) error {
	base := filepath.Base(path)
	dir := filepath.Dir(path)
	filename := strings.TrimSuffix(base, filepath.Ext(base))
	newPath := filepath.Join(dir, fmt.Sprintf("%s_resized_%dw_%dh%s", filename, width, height, filepath.Ext(base)))

	file, err := os.Create(newPath)
# FIXME: 处理边界情况
	if err != nil {
# 改进用户体验
		return err
# NOTE: 重要实现细节
	}
	defer file.Close()

	if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		return jpeg.Encode(file, img, nil)
	} else if strings.HasSuffix(path, ".png") {
# NOTE: 重要实现细节
		return png.Encode(file, img)
	}
# 添加错误处理

	return nil
}
