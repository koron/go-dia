package dia

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"os"
)

// LineProc is a type of callback for each line read.
type LineProc func(line string) error

var (
	// FileNum shows number of current processing file.
	FileNum int

	// Filename shows name of current processing file.
	Filename string

	// LineNum shows line number of current processing file.
	LineNum int
)

// For process stdin or files in arguments.
func For(proc LineProc) error {
	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	if len(args) == 0 {
		FileNum = 0
		Filename = ""
		return procReader(proc, os.Stdin)
	}
	for i, name := range args {
		FileNum = i + 1
		if err := procFile(proc, name); err != nil {
			return err
		}
	}
	return nil
}

func procFile(proc LineProc, name string) error {
	f, err := os.Open(name)
	Filename = name
	if err != nil {
		return err
	}
	defer f.Close()
	return procReader(proc, f)
}

func readLine(r *bufio.Reader) (*bytes.Buffer, error) {
	b, isPrefix, err := r.ReadLine()
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	bb := bytes.NewBuffer(b)
	if isPrefix {
		for {
			b, cont, err := r.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}
			if _, err := bb.Write(b); err != nil {
				return nil, err
			}
			if !cont {
				break
			}
		}
	}
	return bb, nil
}

func procReader(proc LineProc, reader io.Reader) error {
	r := bufio.NewReader(reader)
	for i := 1; ; i++ {
		b, err := readLine(r)
		if err != nil {
			return err
		} else if b == nil {
			break
		}
		LineNum = i
		err = proc(b.String())
		if err != nil {
			return err
		}
	}
	return nil
}
