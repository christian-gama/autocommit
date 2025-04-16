package groq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func httpRequest(method, url, apiKey string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	return req, nil
}

type httpErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   any    `json:"param"`
	Type    string `json:"type"`
}

type httpErrorResponse struct {
	Error httpErrorDetail `json:"error"`
}

func handleHTTPResponse(response *http.Response) error {
	if response.StatusCode < 400 {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var parsedResponse httpErrorResponse
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}

	return errors.New(parsedResponse.Error.Message)
}
