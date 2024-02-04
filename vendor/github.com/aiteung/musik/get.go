package musik

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CloneRepo(url string) {
	filename := DownloadPackage(url)
	Unzip(filename)
}

func DownloadPackage(url string) (filename string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	urlarr := strings.Split(url, "/")
	out, _ := os.Create(urlarr[len(urlarr)-1])
	defer out.Close()
	io.Copy(out, resp.Body)
	return urlarr[len(urlarr)-1]
}

func Unzip(src string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := f.Name
		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			fmt.Println(fpath)
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()

		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}

	}
	return filenames, nil

}
