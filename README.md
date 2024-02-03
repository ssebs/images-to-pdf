# images-to-pdf

Convert a selection of images files into a single PDF.

## Building
- git clone https://github.com/ssebs/images-to-pdf
  - You'll need to install a C compiler. See https://developer.fyne.io/started/
- `go build -o img2pdf.exe .\cmd\main.go`
- Once the GUI is working...
  - Make sure `fyne` CLI is installed
    - `go install fyne.io/fyne/v2/cmd/fyne@latest`
  - Windows:
    - `PS go-mmp> fyne package -os windows`
  - Mac:
    - `$ fyne package -os darwin`
  - Linux:
    - `$ fyne package -os linux`

## LICENSE
[Apache 2 License](./LICENSE)
