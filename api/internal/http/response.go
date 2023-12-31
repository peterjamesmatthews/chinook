package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// WriteJSONToResponse writes a JSON response to the http.ResponseWriter.
//
// # Errors:
//   - failed to marshal payload
//   - failed to write payload
func WriteJSONToResponse(w http.ResponseWriter, payload any) error {
	buffer := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	// marshal payload
	err := encoder.Encode(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload\n%w", err)
	}

	// set headers
	w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
	w.Header().Set("Content-Type", "application/json")

	// write payload to response
	if _, err = w.Write(buffer.Bytes()); err != nil {
		return fmt.Errorf("failed to write payload\n%w", err)
	}

	return nil
}
