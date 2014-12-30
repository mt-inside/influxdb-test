package main

import (
    "strconv"
    "sync"
    "time"
    "github.com/influxdb/influxdb/client"
)

const boxen = 1000
const datums = 10

func main( ) {

    var waiter sync.WaitGroup

    for i := 0; i < boxen; i++ {
	waiter.Add( 1 )
	go spam( i, &waiter )
    }

    waiter.Wait()
}

func spam( box int, waiter *sync.WaitGroup ) {

    defer waiter.Done()

    c, err := client.NewClient( &client.ClientConfig{
	Database: "load",
    } )
    if err != nil {
	panic(err)
    }

    series := &client.Series {
	Name: "box" + strconv.Itoa( box ),
	Columns: []string{ "cpu", "ram" },
    }

    for i := 0; i < datums; i++ {
	series.Points = [][]interface{} { { 20+i , 100 }, }

	err = c.WriteSeries([]*client.Series{ series } )
	if err != nil {
	    panic(err)
	}

	time.Sleep( 1 * time.Second )
    }
}
