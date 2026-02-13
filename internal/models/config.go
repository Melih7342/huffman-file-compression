package models

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	CompressMode   bool
	DecompressMode bool
	Verbosity      bool
	Recursive      bool
	Performance    bool
	Mode           string
	OutputPath     string
	InputPaths     []string
}

func ParseConfig() *Config {
	c := flag.Bool("c", false, "Compress file(s)")
	d := flag.Bool("d", false, "Decompress file(s)")
	v := flag.Bool("v", false, "Verbose output")
	r := flag.Bool("r", false, "Recursive directory content compression")
	h := flag.Bool("h", false, "Help")
	p := flag.Bool("p", false, "Performance metrics")
	o := flag.String("o", "", "Choose custom output location")

	flag.Parse()

	if *h {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) == 0 {
		printExamples()
		os.Exit(0)
	}

	mode := "c"

	if *d {
		mode = "d"
	}

	return &Config{
		CompressMode:   *c,
		DecompressMode: *d,
		Verbosity:      *v,
		Recursive:      *r,
		Performance:    *p,
		Mode:           mode,
		OutputPath:     *o,
		InputPaths:     args,
	}
}

func printExamples() {
	fmt.Println("\nExamples:")
	fmt.Println(" - Compression: ./huff -c test.txt")
	fmt.Println(" - Decompression: ./huff -d test.txt.huff")
	fmt.Println(" - Define output location: ./huff -c -o C:/Users/myUser test.txt")
	fmt.Println(" - Compress files in a directory: ./huff -c -r C:/Users/myUser")
}
