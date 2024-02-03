package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ssebs/images-to-pdf/convert"
)

func main() {
	fmt.Println("Images to PDF")

	dir := flag.String("d", ".", "Folder where images are stored")
	dest := flag.String("o", "out.pdf", "Filename of PDF file")
	shouldArchive := flag.Bool("a", false, "Whether or not to archive images")
	flag.Parse()

	// Get files in dir
	files, err := convert.ListFiles(*dir)
	CheckErr(err)
	if len(files) == 0 {
		log.Fatal("No files found in ", *dir)
	}

	// Convert to PDF
	pdf, err := convert.ImagesToPDF(files)
	CheckErr(err)

	// Save PDF
	err = os.WriteFile(*dest, pdf.Contents, 0644)
	CheckErr(err)

	// Archive old files
	if *shouldArchive {
		err = convert.ArchiveImages(*dir, files)
		CheckErr(err)
	}

	fmt.Println("dir:", *dir)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
