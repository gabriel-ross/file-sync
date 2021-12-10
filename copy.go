package main

import (
	"io"
	"log"
	"os"
)

func CopyFileWindows(dstDir, src string, ignore *[]string) {
	/*
		copies src into directory dstDir, ignoring any files passed to ignore
		input:
		* dstDir: directory
		* src: file or directory
	*/
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		log.Fatal(err)
	}
	srcName := srcFileInfo.Name()

	// check if src is in ignore
	for _, ignoreFile := range *ignore {
		if srcName == ignoreFile {
			return
		}
	}

	// TODO: check if file already exists in destination directory and decide whether to overwrite
	// it based off which file was most recently updated?
	// or just create a copy with the addition of "-copy" to filename

	switch srcFileInfo.IsDir() {
	case false:
		// src is not directory
		srcFile, err := os.Open(src)
		if err != nil {
			log.Fatal(err)
		}
		defer srcFile.Close()

		copyFile, err := os.Create(dstDir + "\\" + srcName)
		if err != nil {
			log.Fatal(err)
		}
		defer copyFile.Close()

		_, err = io.Copy(copyFile, srcFile)
		if err != nil {
			log.Fatal(err)
		}
		return
	case true:
		// src is directory. copy directory with permissions and copy contents of directory
		err = os.Mkdir(dstDir+"\\"+srcName, srcFileInfo.Mode().Perm())
		if err != nil {
			log.Fatal(err)
		}
		contents, err := os.ReadDir(src)
		if err != nil {
			log.Fatal(err)
		}

		// copy contents of src directory to newly copied directory
		// TODO: clean up string concatenation (use string builder or array and join)
		for _, file := range contents {
			CopyFileWindows(dstDir+"\\"+srcName, src+"\\"+file.Name(), ignore)
		}
		return
	}
}
