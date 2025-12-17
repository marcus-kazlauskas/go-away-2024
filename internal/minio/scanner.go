package minio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// MinioScanner handles multipart borders in txt from MinIO
type MinioScanner struct {
	s   *bufio.Scanner
	eof *string
}

func NewScanner(file *os.File) (MinioScanner, error) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	delimiter := "----------------------------"
	scanner.Scan()
	line := scanner.Text()
	if !strings.Contains(line, delimiter) {
		return MinioScanner{}, DataError()
	}
	endOfFileId := strings.Trim(line, "-")

	for range 3 {
		scanner.Scan()
	}

	return MinioScanner{
		s:   scanner,
		eof: &endOfFileId,
	}, nil
}

func DataError() error {
	return fmt.Errorf("incorrect input data")
}

func (m *MinioScanner) Line() (string, error) {
	scan := m.s.Scan()
	line := m.s.Text()
	if !scan && m.eof == nil || m.eof != nil && strings.Contains(line, *m.eof) {
		return "", io.EOF
	}
	return line, nil
}

func (m *MinioScanner) LineLink() (*string, *error) {
	line, err := m.Line()
	lineLink := &line
	errLink := &err
	return lineLink, errLink
}
