package main

import (
    "fmt"
    "encoding/csv"
    "io"
    "os"
    "math/rand"
    "time"
)

type Users struct {
    id, first_name, last_name string
}

func RecordToUsers (record []string, col_index []int) (u Users) {
    u.id = record[col_index[0]]
    u.first_name = record[col_index[1]]
    u.last_name = record[col_index[2]]
    return
}

func LoadCSVDataToChannel (in io.Reader, col_index []int) <-chan Users{
    out := make(chan Users)
    go func (){
        defer close(out)
        reader := csv.NewReader(in)
        for {
            record, err := reader.Read()
            if err == io.EOF {
                return
            }
            user := RecordToUsers(record,col_index)
            out <- user
        }
    }()
    return out
}

func main(){
    start_time := time.Now()
    rFile, err := os.Open("data/small.csv")
    if err != nil {
        fmt.Println("Error:",err)
        return
    }
    // reader := csv.NewReader(rFile)
    // records, err := reader.ReadAll()
    rand.Seed(int64(time.Now().Nanosecond()))
    col_index := rand.Perm(3)
    results := LoadCSVDataToChannel(rFile,col_index)
    for line := range results {
        fmt.Println(line)
    }


    // LoadCSVDataToChannel(records[:len(records)/2],col_index,c)
    // LoadCSVDataToChannel(records[len(records)/2:],col_index,c)

    // x, y := <-c, <-c
    // fmt.Println(x)
    // fmt.Println(y)

    // Loading csv file
    // rFile, err := os.Open("data/small.csv") //3 columns
    //rFile, err := os.Open("data/dummy.csv") //7 columns
    // if err != nil {
    //     fmt.Println("Error:", err)
    //     return
    // }
    // defer rFile.Close()

    // Creating csv reader
    // reader := csv.NewReader(rFile)
    // reader.Comma = ',' //comma is the default delimiter, so I'm just adding this line for future reference
    // lines, err := reader.ReadAll()
    // if err == io.EOF {
    //     fmt.Println("Error:", err)
    //     return
    // }

    // Creating csv writer
    // wFile, err := os.Create("data/result_channels.csv")
    // if err != nil {
    //     fmt.Println("Error:",err)
    //     return
    // }
    // defer wFile.Close()
    // writer := csv.NewWriter(wFile)
    //
    // // Read data, randomize columns and write new lines to results.csv
    // rand.Seed(int64(time.Now().Nanosecond()))
    // var col_index []int
    // for i,line :=range lines{
    //     if i == 0 {
    //         //randomize column index based on the number of columns recorded in the 1st line
    //         col_index = rand.Perm(len(line))
    //     }
    //     //writer.Write([]string{line[col_index[0]], line[col_index[1]], line[col_index[2]]}) //3 columns
    //     writer.Write([]string{line[col_index[0]], line[col_index[1]], line[col_index[2]], line[col_index[3]], line[col_index[4]], line[col_index[5]], line[col_index[6]]})
    //     writer.Flush()
    // }

    //print report
    // fmt.Println("No. of lines: ",len(lines))
    fmt.Println("Time taken: ", time.Since(start_time))

}
