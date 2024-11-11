package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/karlpokus/ago"
)

var in = flag.String("in", "iso", "input timestamp format: s, ms, ns or iso-8601")
var out = flag.String("out", "all", "output format: ago, all")

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	ts := flag.Arg(0)
	if ts == "" {
		log.Fatalf("Missing timestamp input")
	}
	t, err := parseTimestamp(ts, *in)
	if err != nil {
		log.Fatal(err)
	}
	if *out == "ago" {
		log.Println(ago.ParseWithContext(t))
		return
	}
	log.Println(t.Format(time.RFC3339))
	log.Println(ago.ParseWithContext(t))
	log.Println(t.Unix())
	log.Println(t.UnixMilli())
	log.Println(t.UnixNano())
}

// parseTimestamp parses the timestamp in ts formatted as f to time.Time
func parseTimestamp(ts string, f string) (time.Time, error) {
	if f == "iso" {
		return time.Parse(time.RFC3339, ts)
	}
	var t time.Time
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return t, err
	}
	switch f {
	case "s":
		t = time.Unix(i, 0)
	case "ms":
		t = time.Unix(i/1000, 0)
	case "ns":
		t = time.Unix(0, i)
	}
	return t, nil
}
