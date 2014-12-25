package main

import (
    "strconv"
    "time"
    "github.com/influxdb/influxdb/client"
)

const boxen = 1000
const datums = 10

func main( ) {

    done := make( chan bool )
    finished := 0

    for i := 0; i < boxen; i++ {
	go spam( i, done )
    }

    for
    {
	// Have to block this goroutine until the network transaction has finished,
	// as this one ending seems to kill them all.
	<- done
	if finished++; finished == boxen { break }
    }
}

func spam( box int, done chan bool ) {

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

    done <- true
}
