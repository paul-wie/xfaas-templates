package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go and xfaas: %s", string(req))
}
