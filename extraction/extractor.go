package extraction

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// ExtractTarGz opens the supplied tar.gz archive file and reads the contents to the given destination.
// If a map of file extension is supplied only those that match will be extracted.
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
