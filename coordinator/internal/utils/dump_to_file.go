package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func DumpToFile(path string, m map[string]uint64) error {
	f, err := os.Create(path)
	if err != nil {return err}
	defer f.Close()

	bw := bufio.NewWriterSize(f, 64*1024)
	defer bw.Flush()

	type kv struct {
		w string;
		c uint64
	}
	pairs := make([]kv, 0, len(m))

	for w, c := range m {
		pairs = append(pairs, kv{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].c > pairs[j].c
	})

	for _, p := range pairs {
		fmt.Fprintf(bw, "%s\t%d\n", p.w, p.c)
	}

	return nil
}