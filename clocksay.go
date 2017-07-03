package main

import (
	"github.com/d2r2/go-i2c"
	"log"
	"os"
)

var font = map[byte]uint{
	' ':  0x0000,
	'!':  0x0006,
	'"':  0x0220,
	'#':  0x12ce,
	'$':  0x12ed,
	'%':  0x0c24,
	'&':  0x235d,
	'\'': 0x0400,
	'(':  0x2400,
	')':  0x0900,
	'*':  0x3fc0,
	'+':  0x12c0,
	',':  0x0800,
	'-':  0x00c0,
	'.':  0x0000,
	'/':  0x0c00,
	'0':  0x0c3f,
	'1':  0x0006,
	'2':  0x00db,
	'3':  0x008f,
	'4':  0x00e6,
	'5':  0x2069,
	'6':  0x00fd,
	'7':  0x0007,
	'8':  0x00ff,
	'9':  0x00ef,
	':':  0x1200,
	';':  0x0a00,
	'<':  0x2400,
	'=':  0x00c8,
	'>':  0x0900,
	'?':  0x1083,
	'@':  0x02bb,
	'A':  0x00f7,
	'B':  0x128f,
	'C':  0x0039,
	'D':  0x120f,
	'E':  0x00f9,
	'F':  0x0071,
	'G':  0x00bd,
	'H':  0x00f6,
	'I':  0x1200,
	'J':  0x001e,
	'K':  0x2470,
	'L':  0x0038,
	'M':  0x0536,
	'N':  0x2136,
	'O':  0x003f,
	'P':  0x00f3,
	'Q':  0x203f,
	'R':  0x20f3,
	'S':  0x00ed,
	'T':  0x1201,
	'U':  0x003e,
	'V':  0x0c30,
	'W':  0x2836,
	'X':  0x2d00,
	'Y':  0x1500,
	'Z':  0x0c09,
	'[':  0x0039,
	'\\': 0x2100,
	']':  0x000f,
	'^':  0x0c03,
	'_':  0x0008,
	'`':  0x0100,
	'a':  0x1058,
	'b':  0x2078,
	'c':  0x00d8,
	'd':  0x088e,
	'e':  0x0858,
	'f':  0x0071,
	'g':  0x048e,
	'h':  0x1070,
	'i':  0x1000,
	'j':  0x000e,
	'k':  0x3600,
	'l':  0x0030,
	'm':  0x10d4,
	'n':  0x1050,
	'o':  0x00dc,
	'p':  0x0170,
	'q':  0x0486,
	'r':  0x0050,
	's':  0x2088,
	't':  0x0078,
	'u':  0x001c,
	'v':  0x2004,
	'w':  0x2814,
	'x':  0x28c0,
	'y':  0x200c,
	'z':  0x0848,
	'{':  0x0949,
	'|':  0x1200,
	'}':  0x2489,
	'~':  0x0520,
}

func setup(i2c *i2c.I2C) {
	/* Enable oscillator  */
	_, err := i2c.Write([]byte{0x21})
	if err != nil {
		log.Fatal(err)
	}

	/* Max brightness */
	brightness := byte(15)

	/* Set brightness (0-15) */
	_, err = i2c.Write([]byte{0xe0 | brightness})
	if err != nil {
		log.Fatal(err)
	}

	/* Display ON, blinking OFF */
	_, err = i2c.Write([]byte{0x81})
	if err != nil {
		log.Fatal(err)
	}
}

func writeString(i2c *i2c.I2C, s string) {

	payload := make([]byte, 17)

	/* Address 0x00 */
	payload[0] = 0

	for i := 0; i < 8; i++ {
		var char byte
		if i >= len(s) {
			char = ' '
		} else {
			char = s[i]
		}
		bitmap := font[char];

		payload[i*2+1] = byte(bitmap & 0xff)        /* LSB */
		payload[i*2+2] = byte((bitmap >> 8) & 0xff) /* MSB */
	}

	_, err := i2c.Write(payload)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	/* RaspberryPi's I2C defaults to bus 1 */
	/* HT16K33 defaults to address 0x70 */
	i2c, err := i2c.NewI2C(0x70, 1)
	if err != nil {
		log.Fatal(err)
	}

	defer i2c.Close()

	setup(i2c)

	writeString(i2c, os.Args[1])
}
