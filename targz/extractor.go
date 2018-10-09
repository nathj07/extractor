package targz

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Extract opens the supplied tar.gz archive file and reads the contents to the given destination.
// If a map of file extension is supplied only those that match will be extracted.
// The returned slice holds the paths to all the extracted files. If there is an error returned the
// slice will be nil.
func Extract(archive, dest string, formats map[string]struct{}) ([]string, error) {
	f, err := os.Open(archive)
	if err != nil {
		return nil, fmt.Errorf("Error opening gzip file: %v", err)
	}
	defer f.Close()
	gzipStream := bufio.NewReader(f)
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return nil, fmt.Errorf("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)
	extracted := []string{}
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filepath.Join(dest, header.Name), 0755); err != nil {
				return nil, fmt.Errorf("ExtractTarGz: MkdirAll() failed: %s", err.Error())
			}
		case tar.TypeReg:
			if len(formats) > 0 {
				if _, ok := formats[filepath.Ext(header.Name)]; !ok {
					continue
				}
			}
			outFile, err := os.Create(filepath.Join(dest, header.Name))
			if err != nil {
				return nil, fmt.Errorf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return nil, fmt.Errorf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			extracted = append(extracted, outFile.Name())
		default:
			return nil, fmt.Errorf("ExtractTarGz: unknown type: %s in %s", string(header.Typeflag), header.Name)
		}
	}
	return extracted, nil
}
