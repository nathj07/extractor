package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nathj07/targz_extractor/extractor"
)

// TODO: add license
// TODO: will this recurse and unpack directories?
// TODO: exclude extensions?
// TOTO: named files to extract?

var (
	gzipPath   = flag.String("gzip", "", "path to tar.gzip file file")
	outputPath = flag.String("dest", "", "where to write the gzip contents")
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
	if *gzipPath == "" {
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

	extractor.ExtractTarGz(*gzipPath, dest, formats)
}
