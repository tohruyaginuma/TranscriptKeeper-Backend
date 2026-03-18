package external

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type CloudflareWorkersAIClient struct {
	APIKey     string
	AccountID  string
	HTTPClient *http.Client
}

func NewCloudflareWorkersAIClient(apiKey string, accountID string) *CloudflareWorkersAIClient {
	return &CloudflareWorkersAIClient{
		APIKey:    apiKey,
		AccountID: accountID,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *CloudflareWorkersAIClient) Transcribe(ctx context.Context, audio io.Reader, language string) (result WhisperOutput, err error) {
	b, err := io.ReadAll(audio)
	if err != nil {
		return WhisperOutput{}, err
	}

	ctx = context.WithoutCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	audioBase64 := base64.StdEncoding.EncodeToString(b)

	requestBody := WhisperInput{
		Audio:     audioBase64,
		Task:      "transcribe",
		Language:  language,
		VadFilter: true,
		BeamSize:  5,
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return WhisperOutput{}, err
	}

	modelName := "@cf/openai/whisper-large-v3-turbo"
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", c.AccountID, modelName)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return WhisperOutput{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return WhisperOutput{}, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return WhisperOutput{}, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return WhisperOutput{}, fmt.Errorf("failed to transcribe: status=%d message=%s", resp.StatusCode, parseCloudflareError(respBytes))
	}

	type cloudflareResponse struct {
		Result  WhisperOutput `json:"result"`
		Success bool          `json:"success"`
		Errors  []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	var wrapped cloudflareResponse
	if err := json.Unmarshal(respBytes, &wrapped); err == nil && (wrapped.Success || len(wrapped.Errors) > 0) {
		if !wrapped.Success {
			return WhisperOutput{}, fmt.Errorf("cloudflare response unsuccessful: %s", parseCloudflareError(respBytes))
		}
		return wrapped.Result, nil
	}

	var response WhisperOutput
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return WhisperOutput{}, err
	}

	return response, nil
}

func parseCloudflareError(respBytes []byte) string {
	type cloudflareErrorResponse struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	var response cloudflareErrorResponse
	if err := json.Unmarshal(respBytes, &response); err == nil && len(response.Errors) > 0 {
		return response.Errors[0].Message
	}

	body := strings.TrimSpace(string(respBytes))
	if body == "" {
		return "empty response body"
	}

	if len(body) > 512 {
		return body[:512] + "..."
	}

	return body
}
