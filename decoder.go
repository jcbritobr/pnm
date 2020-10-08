package pnm

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

var (
	errWrongFormat     = errors.New("Wrong image format")
	errMalformedHeader = errors.New("Malformed image header")
)

type header struct {
	magicNumber int
	width       int
	height      int
	maxValue    byte
}

// Decoder decodes a byte or text stream into a pnm image format
type Decoder struct {
	reader *bufio.Reader
	format int64
}

// NewDecoder creates a pnm image decoder
func NewDecoder(reader io.Reader, format int64) *Decoder {
	return &Decoder{reader: bufio.NewReader(reader), format: format}
}

func (d *Decoder) decodeHeader() (*header, error) {
	headerBuf := make([]byte, 0)
	var err error
	var b byte
	var h header
	var numField int

	switch d.format {
	case PPMBinary, PGMBinary, PPMText, PGMText:
		numField = 4
	default:
		numField = 3
	}

	comment := false
	for fields := 0; fields < numField; {
		b, _ = d.reader.ReadByte()
		if b == '#' {
			comment = true
		} else if !comment {
			headerBuf = append(headerBuf, b)
		}
		if comment && b == '\n' {
			comment = false
		} else if !comment && (b == ' ' || b == '\n' || b == '\t') {
			fields++
		}
	}

	hfields := bytes.Fields(headerBuf)
	mn := string(string(hfields[0]))
	if len(mn) < 2 {
		return nil, errMalformedHeader
	}

	h.magicNumber, err = strconv.Atoi(string(mn[1]))
	if err != nil {
		return nil, err
	}

	if h.magicNumber != int(d.format) {
		return nil, errWrongFormat
	}

	h.width, err = strconv.Atoi(string(hfields[1]))
	if err != nil {
		return nil, err
	}

	h.height, err = strconv.Atoi(string(hfields[2]))
	if err != nil {
		return nil, err
	}

	if numField == 4 {
		mv, err := strconv.Atoi(string(hfields[3]))
		if err != nil {
			return nil, err
		}

		h.maxValue = byte(mv)
	}

	return &h, nil
}

// Decode decodes an open image format file and fills an image struct with these data
func (d *Decoder) Decode(image Image) error {

	header, err := d.decodeHeader()
	if err != nil {
		return err
	}

	updateBuffer := func(i Image) error {

		b := make([]byte, len(i.Buffer()))
		_, err = io.ReadFull(d.reader, b)
		if err != nil {
			return err
		}

		i.SetBuffer(b)

		return nil
	}

	switch image := image.(type) {
	case *PGMImage:
		*image = *NewPGMImage(int64(header.width), int64(header.height), header.maxValue, int64(header.magicNumber))
		updateBuffer(image)
	case *PPMImage:
		*image = *NewPPMImage(int64(header.width), int64(header.height), header.maxValue, int64(header.magicNumber))
		updateBuffer(image)
	case *PBMImage:
		*image = *NewPBMImage(int64(header.width), int64(header.height), int64(header.magicNumber))
		updateBuffer(image)
	default:
		return errWrongFormat
	}

	return nil
}
