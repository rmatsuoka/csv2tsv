package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var progname = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [file ...]\n", progname)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	log.SetPrefix(progname + ": ")
	log.SetFlags(0)

	flag.Usage = usage
	flag.Parse()

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if flag.NArg() == 0 {
		csv2tsv(w, os.Stdin)
	} else {
		for _, fname := range flag.Args() {
			f, err := os.Open(fname)
			if err != nil {
				log.Print(err)
				continue
			}
			csv2tsv(w, f)
			f.Close()
		}
	}
}

func csv2tsv(w io.Writer, r io.Reader) {
	cr := csv.NewReader(r)
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Print(err)
			return
		}
		output(w, rec)
	}
}

var replacer = strings.NewReplacer("\r", "\\r", "\n", "\\n", "\t", "\\t", "\\", "\\\\")

func output(w io.Writer, rec []string) {
	for i, f := range rec {
		replacer.WriteString(w, f)
		if i == len(rec)-1 {
			fmt.Fprint(w, "\n")
		} else {
			fmt.Fprint(w, "\t")
		}
	}
}
