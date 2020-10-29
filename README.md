# PNM
PNM is a collection of open image formats. They are also referred as **portable anymap format(PNM)**. This library supports ppm(portable pixmap), pgm(portable graymap) and pbm(portable bitmap) formats. There is also an encoder and decoder implementations.

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
import (
    "math/rand"
    "os"
    "github.com/jcbritobr/pnm"
)

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
import (
    "fmt"
    "os"

    "github.com/jcbritobr/pnm"
)

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

* Using ImageBuffer and color.Model (RGBA)
```go
import (
    "os"
    "github.com/jcbritobr/pnm"
    "github.com/jcbritobr/pnm/buffer"
    "github.com/jcbritobr/pnm/color"
)

func main() {
    var image pnm.PPMImage
    file, err := os.Open("testdata/tree_1.ppm")
    if err != nil {
        t.Errorf("fail with %v", err)
    }
    defer file.Close()

    decoder := pnm.NewDecoder(file, pnm.PPMBinary)
    err = decoder.Decode(&image)
    if err != nil {
        t.Errorf("fail to decode image %v", err)
    }

    imbuf := buffer.NewImageBuffer(image)
    rgba := color.RGBA{}
    newbuf := []byte{}
    for {
        n, _ := imbuf.Read(rgba[:])
        if n <= 0 {
            break
        }
        r, g, b, _ := rgba.RGBA()
        r = 255 - r
        g = 255 - g
        b = 255 - b
        newbuf = append(newbuf, r)
        newbuf = append(newbuf, g)
        newbuf = append(newbuf, b)
    }

    image.SetBuffer(newbuf)

    fe, err := os.Create("testdata/tree_3.ppm")
    if err != nil {
        t.Errorf("fail to create file %v", err)
    }
    defer fe.Close()

    encoder := pnm.NewEncoder(fe)
    err = encoder.Encode(&image)
    if err != nil {
        t.Errorf("Encode() = %v want %v", err, nil)
    }
}
```
