package main

import "github.com/influxdb/influxdb/client"

func main( ) {
    spam( )
}

func spam( ) {
    c, err := client.NewClient( &client.ClientConfig{
	Database: "enway",
    } )
    if err != nil {
	panic(err)
    }

    series := &client.Series {
	Name: "load",
	Columns: []string{ "cpu", "ram" },
	Points: [][]interface{} { {20,100}, },
    }

    err = c.WriteSeries([]*client.Series{ series } )
    if err != nil {
	panic(err)
    }
}
