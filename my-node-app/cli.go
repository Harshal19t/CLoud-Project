package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/sqweek/dialog"
)

func main() {
	// Open file dialog
	filePath, err := dialog.File().Filter("Image files", "jpg", "jpeg", "png", "gif").Load()
	if err != nil {
		fmt.Println("Error opening file dialog:", err)
		return
	}

	// Check if a file was selected
	if filePath == "" {
		fmt.Println("No file selected. Exiting.")
		return
	}

	// Print selected file path
	fmt.Println("Selected file:", filePath)

	// Send the selected image to the server
	err = uploadImage(filePath, "10.1.124.10:8080") //add the ip address where you want to send the image
	if err != nil {
		fmt.Println("Error uploading image:", err)
	}
}

func uploadImage(filePath, targetURL string) error {
	// Connect to the server
	conn, err := net.Dial("tcp", targetURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffer to read the file
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return err
	}

	// Copy the image from the buffer to the connection
	_, err = io.Copy(conn, &buf)
	if err != nil {
		return err
	}

	fmt.Println("Image sent successfully.")
	return nil
}
