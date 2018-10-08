package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nathj07/extractor/targz"
)

// TODO: will this recurse and unpack directories?
// TODO: Have the extract function return errors not log
// TODO: unit test coverage of the extract method

var (
	source     = flag.String("source", "", "path to archive file to extract")
	outputPath = flag.String("dest", "", "where to write the extracted files")
	exts       = flag.String("exts", "", "optional CSV list of file extensions, if supplied only files with these extensions will be extracted")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
			%s is a tool to unpack a tar.gz archive and list out the files within. Flags:
			`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *source == "" {
		fmt.Fprintf(os.Stderr, "You must supply a valid path to a tar.gz file")
		os.Exit(1)
	}

	dest := *outputPath
	if dest == "" {
		dest, _ = os.Getwd()
		log.Printf("No output path supplied extracting to %s", dest)
	}
	formats := map[string]struct{}{}
	if *exts != "" {
		for _, ext := range strings.Split(*exts, ",") {
			// prepend the ext with "."
			formats["."+ext] = struct{}{}
		}
	}

	files, err := targz.Extract(*source, dest, formats)
	if err != nil {
		log.Fatalf("failed to extract data: %v", err)
	}
	log.Printf("Unpacked %d file(s) from %s into %s", len(files), *source, dest)
}
