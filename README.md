# images-to-pdf
Convert a selection of images files into a single PDF.

## Usage Example
![gif](./image-to-pdf.gif)

## Building
- From source:
  - `git clone https://github.com/ssebs/images-to-pdf`
    - You'll need to install a C compiler. See https://developer.fyne.io/started/
  - `go build -o img2pdf.exe .\cmd\main.go`
- From go pkg
  - `go install github.com/ssebs/images-to-pdf@latest`

## CLI Usage
```powershell
PS F:\LocalProgramming\images-to-pdf> .\img2pdf.exe -h
Images to PDF
Usage of F:\LocalProgramming\images-to-pdf\img2pdf.exe:
  -a    Whether or not to archive images (use a=false if needed) (default true)
  -d string
        Folder where images are stored (default ".")
  -o string
        Filename of PDF file (default "out.pdf")
PS F:\LocalProgramming\images-to-pdf> 
```

## How to run CLI
- Download [img2pdf.exe from the releases page](https://github.com/ssebs/images-to-pdf/releases/)
- Move the `img2pdf.exe` file to the folder you want to convert images
- EZPZ:
  - Double click the `img2pdf.exe` file
- CLI:
  - SHIFT + RIGHT CLICK the folder > Open in Powershell window here
    - Or command window
  - Type this command:
    - `.\img2pdf.exe -a`

## LICENSE
[Apache 2 License](./LICENSE)
