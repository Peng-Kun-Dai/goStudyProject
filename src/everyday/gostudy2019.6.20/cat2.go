package main

import (
	"fmt"
	"os"
)

//read and write files with slice

func cat2(f *os.File) {
	const NBUF = 1024
	var buf [NBUF]byte //cache slice
	for {
		//read
		switch nr, err := f.Read(buf[:]); true {
		//nr is number of bytes read
		case nr < 0:
			fmt.Fprint(os.Stderr, "error reading %s \n", err.Error())
			os.Exit(1)
		case nr == 0: //EOF
			return
		case nr > 0:
			//write
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprint(os.Stderr, "error writing %s \n", err.Error())
			}

		}
	}

}
