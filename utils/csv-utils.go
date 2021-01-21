package utils

import (
    "os"
    "encoding/csv"
)

type CsvUtils struct {
    
}

func NewCsvUtils() *CsvUtils {
    var n CsvUtils
    return &n
}

func (instance *CsvUtils) WriteFile(fileName string, line map[string]string) {
    file, err := os.Create("result.csv")
    
    if err != nil {
        panic("Failed to open new csv file for writing")
    }

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for k,v := range(line) {
        err := writer.Write([]string{k, v})

        if err != nil {
            panic("Failed to write line to CSV ")
        }
    }

    writer.Flush()
}
