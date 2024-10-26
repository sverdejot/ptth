package ptth

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	Method string
	URI    string

	Protocol string

	Headers map[string]string

	Body string
}

func parseRequest(r io.Reader) (req Request, err error) {
	sc := bufio.NewScanner(r)
	sc.Split(splitHTTPRequest)

	sc.Scan()
	method, uri, protocol, err := parseRequestHeading(sc.Text())
	if err != nil {
		return Request{}, fmt.Errorf("cannot parse HTTP request: %w", err)
	}
	req.Method = method
	req.URI = uri
	req.Protocol = protocol
	sc.Scan()

	headers := make(map[string]string)
	for header := sc.Text(); header != ""; header = sc.Text() {
		key, value, err := parseHeader(header)
		if err != nil {
			return req, fmt.Errorf("cannot parse header: %w", err)
		}
		headers[key] = value
		sc.Scan()
	}
	req.Headers = headers
	sc.Scan()

	req.Body = sc.Text()

	return req, nil
}

func parseHeader(raw string) (string, string, error) {
	parts := strings.Split(raw, ": ")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("malformed header: %s", raw)
	}

	return parts[0], parts[1], nil
}

func parseRequestHeading(raw string) (method, uri, protocol string, err error) {
	parts := strings.Split(raw, " ")

	if len(parts) != 3 {
		err = fmt.Errorf("malfored request heading, only got %v", parts)
		return
	}

	method, uri, protocol = parts[0], parts[1], parts[2]

	return
}

func splitHTTPRequest(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.Index(data, []byte("\r\n")); i >= 0 {
		return i + 2, data[:i], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}
