package adapter

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCounter(T *testing.T) {
	c := &Counter{os.Stdout}
	c.Count(3)

	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	c = &Counter{
		Writer: pipeWriter,
	}

	file, err := ioutil.TempFile("/tmp", "pipe")
	if err != nil {
		panic(err)
	}

	tee := io.TeeReader(pipeReader, file)

	go func() {
		io.Copy(os.Stdout, tee)
	}()

	c.Count(5)
}
