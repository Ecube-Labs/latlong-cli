package main

import (
    "os"
    "github.com/bradfitz/latlong"
    "fmt"
    "strconv"
)

func main() {
    if (len(os.Args) < 3) {
        panic("Not enough arguments");
    }

    lat, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        panic("Cannot parse latitude")
    }

    long, err := strconv.ParseFloat(os.Args[2], 64)
    if err != nil {
        panic("Cannot parse longitude")
    }

    tz := latlong.LookupZoneName(lat, long);

    fmt.Print(tz);
    return;
}
