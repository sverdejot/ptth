package ptth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRequest(t *testing.T) {
	tests := []struct {
		Name       string
		RawRequest string
		Want       Request
	}{
		{
			Name: "parse GET request",
			RawRequest: "GET /status HTTP/1.1\r\nHello: world\r\n\r\n",
			Want: Request{
				Method: "GET",
				URI: "/status",
				Protocol: "HTTP/1.1",
				Headers: map[string]string{
					"Hello": "world",
				},
			},
		},
		{
			Name: "parse POST request",
			RawRequest: "POST /user/1 HTTP/1.1\r\nHello: world\r\n\r\n{\"name\":\"samuel\"}\r\n",
			Want: Request{
				Method: POST,
				URI: "/user/1",
				Protocol: "HTTP/1.1",
				Headers: map[string]string{
					"Hello": "world",
				},
				Body: `{"name":"samuel"}`,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			stubReq := strings.NewReader(tt.RawRequest)

			got, err := parseRequest(stubReq)

			assert.NoError(t, err)
			assert.Equal(t, tt.Want, got)
		})
	}
}
