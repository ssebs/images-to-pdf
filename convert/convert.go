package convert

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/go-pdf/fpdf"
)

// ImgFile represents an image file with its contents and extension.
type ImgFile struct {
	Contents  []byte
	Filename  string
	extension string
}

// PDFFile represents a PDF file with its contents.
type PDFFile struct {
	Contents []byte
}

// ImagesToPDF takes a list of ImgFile objects and generates a PDF file containing these images.
// It returns a pointer to a PDFFile and an error if the operation encounters any issues.
func ImagesToPDF(imgs []ImgFile) (*PDFFile, error) {
	// Create a new PDF instance with A4 size and millimeter units.
	pdf := fpdf.New("P", "mm", "A4", "")

	// Iterate through the list of image files.
	for idx, img := range imgs {
		// Generate a unique name for the image in the PDF.
		imgName := fmt.Sprintf("img-%d", idx)

		// Skip images with no contents.
		if len(img.Contents) == 0 {
			continue
		}

		// Get image dimensions using ImageInfo.
		imgInfo := pdf.RegisterImageOptionsReader(imgName, fpdf.ImageOptions{
			ImageType: img.extension,
			ReadDpi:   true,
		}, bytes.NewReader(img.Contents))

		// Check if image registration is successful.
		if imgInfo == nil {
			fmt.Println("Failed to register image")
			return nil, fmt.Errorf("failed to register image")
		}

		// Calculate the aspect ratio of the image.
		// aspectRatio := float64(imgInfo.Width()) / float64(imgInfo.Height())

		// Add a new page for each image. P for portrait
		pdf.AddPageFormat("P", fpdf.SizeType{Wd: imgInfo.Width(), Ht: imgInfo.Height()})

		// Place the image on the page.
		pdf.ImageOptions(imgName, 0, 0, -1, -1, false, fpdf.ImageOptions{
			ImageType: img.extension,
			ReadDpi:   true,
		}, 0, "")
	}

	// Generate PDF contents.
	var pdfContents bytes.Buffer
	if err := pdf.Output(&pdfContents); err != nil {
		fmt.Println("Error during PDF output:", err)
		return nil, err
	}

	// Return the PDFFile with its contents.
	return &PDFFile{Contents: pdfContents.Bytes()}, nil
}

// ListFiles retrieves a list of image files from the specified directory.
// It returns a slice of ImgFile objects and an error if the operation encounters any issues.
func ListFiles(dir string) ([]ImgFile, error) {
	imgFormats := []string{"jpg", "jpeg", "png", "gif"}
	// Initialize an empty slice to store ImgFile objects.
	var imgFiles []ImgFile

	// Read the directory entries.
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Sort entries by name.
	sort.Slice(entries, func(i, j int) bool {
		return compareFilenames(entries[i].Name(), entries[j].Name())
	})

	// Iterate through directory entries.
	for _, entry := range entries {

		// Skip directories.
		if !entry.IsDir() {
			fmt.Println("e:", entry.Name())
			fp := filepath.Join(dir, entry.Name())

			// Extract the file extension. (remove the . too)
			ext := filepath.Ext(fp)[1:]
			if !stringInSlice(ext, imgFormats) {
				continue
			}

			contents, err := os.ReadFile(fp)
			if err != nil {
				return nil, err
			}

			imgFiles = append(imgFiles, ImgFile{
				Contents: contents, extension: ext,
				Filename: fp,
			})
		}
	}
	return imgFiles, nil
}

// ArchiveImages moves the specified images to an "archive" folder within the provided directory.
// If the "archive" folder does not exist, it will be created.
// It returns an error if any issues occur during the archiving process.
func ArchiveImages(dir string, imgs []ImgFile) error {
	if len(imgs) == 0 {
		return fmt.Errorf("no images provided for archiving")
	}
	archiveFolderPath := filepath.Join(dir, "archive")

	// Check if the archive folder exists, create it if not.
	if _, err := os.Stat(archiveFolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(archiveFolderPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create archive folder: %v", err)
		}
	}

	// Move each image to the archive folder.
	for _, img := range imgs {
		archiveFilePath := filepath.Join(archiveFolderPath, filepath.Base(img.Filename))
		if err := os.Rename(img.Filename, archiveFilePath); err != nil {
			return fmt.Errorf("failed to move image to archive folder: %v", err)
		}
	}
	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func compareFilenames(a, b string) bool {
	// Special case: "Image.jpg" comes before "Image (X).jpg"
	if a == "Image.jpg" && strings.HasPrefix(b, "Image (") {
		return true
	} else if b == "Image.jpg" && strings.HasPrefix(a, "Image (") {
		return false
	}

	// Extract numeric parts from filenames
	re := regexp.MustCompile(`\d+`)
	numsA := re.FindAllString(a, -1)
	numsB := re.FindAllString(b, -1)

	// If both filenames have numeric parts
	if len(numsA) > 0 && len(numsB) > 0 {
		numA, _ := strconv.Atoi(numsA[len(numsA)-1])
		numB, _ := strconv.Atoi(numsB[len(numsB)-1])
		// Compare numeric parts
		if numA != numB {
			return numA < numB
		}
	}

	// If one filename has a numeric part and the other doesn't, prioritize the one with the numeric part
	if len(numsA) > 0 {
		return true
	} else if len(numsB) > 0 {
		return false
	}

	// Default lexicographic comparison
	return a < b
}
