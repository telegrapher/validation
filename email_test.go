package validation

import (
	"fmt"
	"testing"
)

var emailTests = map[string]struct {
	input  string
	output bool
}{
	"email": {
		"yolo@yolo.com", true,
	},
	"not email": {
		"yolo#email", false,
	},
}

func TestEmail(t *testing.T) {
	for name, ct := range emailTests {
		t.Run(fmt.Sprintf("Test %s", name), func(t *testing.T) {
			out := Email(ct.input)
			if out != ct.output {
				t.Fatalf("Input %s returned %t, expected %t.\n", ct.input, out, ct.output)
			}
		})
	}
}

var EmailDomainTests = map[string]struct {
	input  [2]string
	output bool
}{
	"domain match": {
		[2]string{"yolo@email.com", "email.com"}, true,
	},
	"domain mismatch": {
		[2]string{"yolo@email.com", "other.net"}, false,
	},
	"broken email": {
		[2]string{"yolo@@email", "email"}, false,
	},
}

func TestEmailDomainTests(t *testing.T) {
	for name, ct := range EmailDomainTests {
		t.Run(fmt.Sprintf("Test %s", name), func(t *testing.T) {
			out := EmailDomain(ct.input[0], ct.input[1])
			if out != ct.output {
				t.Fatalf("Input %s returned %t, expected %t.\n", ct.input, out, ct.output)
			}
		})
	}
}
