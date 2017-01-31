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
    id, first_name, last_name, email, phone, gender, ip_address string
}

func RecordToUsers (record []string, col_index []int) (u Users) {
    u.id = record[col_index[0]]
    u.first_name = record[col_index[1]]
    u.last_name = record[col_index[2]]
    u.email = record[col_index[3]]
    u.phone = record[col_index[4]]
    u.gender = record[col_index[5]]
    u.ip_address = record[col_index[6]]
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
    rFile, err := os.Open("data/dummy.csv")
    if err != nil {
        fmt.Println("Error:",err)
        return
    }

    rand.Seed(int64(time.Now().Nanosecond()))
    col_index := rand.Perm(7)
    results := LoadCSVDataToChannel(rFile,col_index)

    //creating writer
    wFile, err := os.Create("data/result_channels.csv")
    if err != nil {
        fmt.Println("Error:",err)
        return
    }
    defer wFile.Close()
    writer := csv.NewWriter(wFile)

    for line := range results {
        writer.Write([]string{<-results})
        writer.Flush()
    }

    // LoadCSVDataToChannel(records[:len(records)/2],col_index,c)
    // LoadCSVDataToChannel(records[len(records)/2:],col_index,c)
    // x, y = <-c, <-c

    fmt.Println("Time taken: ", time.Since(start_time))

}
