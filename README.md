# PNM
PNM is a collection of open image formats. They are also refered as **portable anymap format(PNM)**. This library supports ppm(portable pixmap), pgm(portable graymap) and pbm(portable bitmap) formats. There is also an encoder and decoder implementations.

* Description \
Each file starts with a two byte magic number (in ascii) that identifies the type of the file it is (PPM, PGM, PBM) and its encoding (**ascii/plain or binary/raw**). The magic number is a capital P followed by a single-digit number. Below is a explanation table with **ascii/binary** magic number formats.

| **Type**        | **Magic Number** | **Extension** | **Color**                                                                    |
|-----------------|------------------|---------------|------------------------------------------------------------------------------|
| Portable Bitmap | P1/P4            | .pbm          | 0-1(White & Black)                                                           |
| Portable GrayMap| P2/P5            | .pgm          | 0-255(gray scale), variable, black to white range                            |
| Portable PixMap | P3/P6            | .ppm          | 16 777 216 (0-255 for each RGB channel), some support for 0-65535 per channel|

<p></p>

The ascii/plain format allow for human readability and easy transfer to other platforms. The binary/raw formats are more efficient in size but will be dependent of platforms.

* Encoder usage
```go
"math/rand"
"os"
"github.com/jcbritobr/pnm"
    
func main() {
    image := pnm.NewPGMImage(800, 800, 255, pnm.PGMBinary)
    file, err := os.Create("testdata/synimage.pgm")
    if err != nil {
        panic(err)
    }
    encoder := pnm.NewEncoder(file)

    imgbuf := image.Buffer()

    for i := range image.Buffer() {
        data := rand.Intn(255-0) + 0
        imgbuf[i] = byte(data)
    }

    err = encoder.Encode(image)
    if err != nil {
        panic(err)
    }
}
```

* Decoder usage
```go
"fmt"
"os"

"github.com/jcbritobr/pnm"

func main() {
    var image pnm.PGMImage
    file, err := os.Open("testdata/synimage.pgm")
    if err != nil {
        panic(err)
    }
    decoder := pnm.NewDecoder(file, pnm.PGMBinary)
    err = decoder.Decode(&image)

    if err != nil {
        panic(err)
    }
}
```
