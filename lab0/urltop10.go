package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// URLTop10 generates RoundsArgs for getting the 10 most frequent URLs.
func URLTop10(nWorkers int) RoundsArgs {
	var args RoundsArgs
	// map: count per map file
	// reduce: add counts, sort and get the 10 most frequent URLs
	args = append(args, RoundArgs{
		MapFunc:    URLCountMap,
		ReduceFunc: URLCountReduce,
		NReduce:    1,
	})
	return args
}

// URLCountMap is the map function in the first round
func URLCountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	kvMap := make(map[string]int)
	var kvs []KeyValue
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		kvMap[l] += 1
	}
	for k, v := range kvMap {
		s := fmt.Sprintf("%s %s\n", k, strconv.Itoa(v))
		kvs = append(kvs, KeyValue{"", s})
	}
	return kvs
}

// URLCountReduce is the reduce function in the first round
func URLCountReduce(key string, values []string) string {
	kvMap := make(map[string]int)
	for _, v := range values {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		kvMap[tmp[0]] += n
	}
	// get top 10
	us, cs := TopN(kvMap, 10)
	buf := new(bytes.Buffer)
	for i := range us {
		_, err := fmt.Fprintf(buf, "%s: %d\n", us[i], cs[i])
		if err != nil {
			panic(err)
		}
	}
	return buf.String()
}
