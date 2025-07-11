package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func writeMessage(w io.Writer, msg string) {
	fmt.Fprintf(w, "Message: %s\n", msg)
}

func package_() {
	fmt.Println("====== io package =====")

	// Open source file for writing
	file1, err := os.OpenFile("files/source.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer file1.Close()

	// Write a message to file1
	writeMessage(file1, "Hello, World!")

	// Reset file1 pointer to the beginning for reading
	_, err = file1.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking in source file:", err)
		return
	}

	// Read contents from file1
	data, err := io.ReadAll(file1)
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}
	fmt.Println(string(data))

	// Open destination file for writing
	file2, err := os.OpenFile("files/destination.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening destination file:", err)
		return
	}
	defer file2.Close()

	// Reset file1 pointer again for copying
	_, err = file1.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking in source file:", err)
		return
	}

	// Copy contents from file1 to file2
	copied, err := io.Copy(file2, file1)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Println("Copied", copied, "bytes")

	fmt.Println("====== os package =====")

	fmt.Println(os.Stat("readme.txt")) // get file info, useful for checking if a file exists
	os.Create("files/new.txt")         // create a new file
	/*
			  os.Open("files/new.txt") // open a file for reading
			  os.OpenFile("files/new.txt", os.O_RDWR|os.O_CREATE, 0644) // open a file for reading and writing
			  os.Remove("files/new.txt") // remove a file
		      os.Rename("files/new.txt", "files/renamed.txt") // rename a file
		      os.Mkdir("files/newdir", 0755) // create a new directory

	*/
	fmt.Println(os.Getwd()) // get current working directory
	// fmt.Println(os.Environ())
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println(os.Getenv("API_KEY"))

	fmt.Println("====== path package =====")
	//Go programs often run on different operating systems (Windows, Linux, macOS) which use different path separators (\ vs /). The path/filepath package abstracts this away.

	/*
		filepath.Join(elem ...string): Joins path elements using the system-specific separator. This is a must for building paths.

		filepath.Ext(path string): Returns the file extension.

		filepath.Base(path string): Returns the last element of the path.

		filepath.Dir(path string): Returns all but the last element of the path.

		filepath.Abs(path string): Returns an absolute representation of the path.

		filepath.Walk(root string, fn WalkFunc): Recursively walks the file tree rooted at root, calling a user-provided function for each file or directory.
	*/
	dir := "files"
	file := "new.txt"
	path := filepath.Join(dir, file)
	fmt.Println(path)

	fmt.Printf("Base: %s\n", filepath.Base(path)) // users.json
	fmt.Printf("Dir: %s\n", filepath.Dir(path))   // data
	fmt.Printf("Ext: %s\n", filepath.Ext(path))   // .json

	filepath.Walk("files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("  Found: %s (Is Dir? %t)\n", path, info.IsDir())
		return nil
	})

	/*
		 When to Use os vs. path/filepath
		os: For performing actions on the file system (creating, reading, writing, deleting).

		path/filepath: For manipulating path strings in a portable, cross-platform way.
	*/

	fmt.Println("====== time package =====")
	now := time.Now()
	fmt.Println(now)

	// returning individual components of the time
	fmt.Println(now.Year())

	// creating a new time
	newTime := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(newTime)

	d := 24 * time.Hour
	fmt.Println(d)
	fmt.Println(d.Seconds())

	parsedDuration, err := time.ParseDuration("1h30m")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	fmt.Println(parsedDuration)

	layout := "Mon, Jan 2, 2006 at 15:04:05"
	formattedTime := now.Format(layout)
	fmt.Println(formattedTime)

	parsedTime, err := time.Parse(layout, formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(parsedTime)
}
