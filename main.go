// fb2zip project main.go
package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var outputdir string

func walkpath(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && filepath.Ext(path) == ".fb2" {
		//имена для папки и архивного файла
		zipFileName := filepath.Join(outputdir, path) + ".zip"
		zipDirName := filepath.Dir(zipFileName)
		//созд. папку для архива
		if err := os.MkdirAll(zipDirName, 0666); err != nil {
			log.Fatal(err)
		}
		//созд. файл архива
		zipFile, err := os.Create(zipFileName)
		if err != nil {
			log.Fatal(err)
		}
		// созд. архиватор
		zipper := zip.NewWriter(zipFile)
		//созд. файл в архиве
		zipped, err := zipper.Create(f.Name())
		// читаем исх. файл
		buffer, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		// пишем в архив
		_, err = zipped.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}

		if err := zipper.Close(); err != nil {
			log.Fatal(err)
		}
		log.Println(zipFileName)
	}
	return nil
}

func main() {
	//архивируем начиная с текущей папки
	outputdir, _ = filepath.Abs(filepath.Dir("."))
	//результат в папку с исходным именем плюс _fb2zip
	outputdir += "_fb2zip"
	filepath.Walk(filepath.Base(""), walkpath)
}
