package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// TODO: refactor to cmd package and extract package
// TODO: allow cli to accept formats
// TODO: add readme
// TODO: add .gitignore
// TODO: add license
var (
	gzipPath   = flag.String("gzip", "", "path to tar.gzip file file")
	outputPath = flag.String("dest", "", "where to write the gzip contents")
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
	}
	formats := map[string]struct{}{}
	formats[".nxml"] = struct{}{}
	formats[".pdf"] = struct{}{}
	formats[".html"] = struct{}{}
	ExtractTarGz(*gzipPath, dest, nil)
}

func ExtractTarGz(gzipPath, dest string, formats map[string]struct{}) {
	f, err := os.Open(gzipPath)
	if err != nil {
		log.Fatalf("Error opening gzip file: %v", err)
	}
	defer f.Close()
	gzipStream := bufio.NewReader(f)
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(dest, header.Name), 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			if len(formats) > 0 {
				if _, ok := formats[filepath.Ext(header.Name)]; !ok {
					continue
				}
			}
			outFile, err := os.Create(filepath.Join(dest, header.Name))
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			fmt.Println("Written:", outFile.Name())
		default:
			log.Fatalf(
				"ExtractTarGz: unknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}
	}
}
