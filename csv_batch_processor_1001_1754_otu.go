// 代码生成时间: 2025-10-01 17:54:35
package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Command line flags
var inputDir = flag.String("inputDir", "./", "Directory containing CSV files to process")
var outputDir = flag.String("outputDir", "./output", "Directory to store processed CSV files")

// ProcessCSVFile processes a single CSV file
func ProcessCSVFile(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("unable to open file: %w", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return fmt.Errorf("unable to read CSV records: %w", err)
    }

    // Process each record (this is where you would add your custom logic)
    for _, record := range records {
        fmt.Println(strings.Join(record, ", "))
    }

    // Write processed records back to a new file
    outputFileName := strings.TrimSuffix(filePath, ".csv") + "_processed.csv"
    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return fmt.Errorf("unable to create output file: %w", err)
    }
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    defer writer.Flush()
    if err := writer.WriteAll(records); err != nil {
        return fmt.Errorf("unable to write back processed records: %w", err)
    }

    return nil
}

// ProcessCSVFilesInDir processes all CSV files in the given directory
func ProcessCSVFilesInDir(directory string) error {
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        return fmt.Errorf("unable to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
            filePath := filepath.Join(directory, file.Name())
            if err := ProcessCSVFile(filePath); err != nil {
                return fmt.Errorf("error processing file %s: %w", file.Name(), err)
            }
        }
    }

    return nil
}

func main() {
    flag.Parse()

    if err := ProcessCSVFilesInDir(*inputDir); err != nil {
        log.Fatalf("failed to process CSV files: %s", err)
    }
}
