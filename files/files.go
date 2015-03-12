package golang_files

import (
	"archive/zip"
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadZipLines(zip_path string) (fChannel chan string) {
	fChannel = make(chan string, 200)
	go func() {
		zip_path, _ := filepath.Abs(zip_path)
		h_zip, err := zip.OpenReader(zip_path)
		if err != nil {
			log.Fatal(err)
		}
		defer h_zip.Close()
		defer close(fChannel)
		for _, tmp := range h_zip.File {
			func() {
				f, err := tmp.Open()
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					fChannel <- scanner.Text()
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}()
		}
	}()
	return
}

func ReadLines(f_path string) (fChannel chan string) {
	fChannel = make(chan string, 200)
	go func() {
		f_path, _ := filepath.Abs(f_path)
		f, err := os.Open(f_path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		defer close(fChannel)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fChannel <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return
}
