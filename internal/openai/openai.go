// Package openai provides a client wrapper for the OpenAI API.
package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

// Client represents an OpenAI API client.
type Client struct {
	// ai is the OpenAI API client.
	ai *openai.Client
}

// NewClient returns a new OpenAI API client.
func NewClient(key string) *Client {
	return &Client{
		ai: openai.NewClient(key),
	}
}

// Request creates a chat completion request with streaming support for the
// given prompt.
func (c *Client) Request(ctx context.Context, model, prompt string) (*openai.ChatCompletionStream, error) {
	systemPrompt := "You're Brane, an expert personal assistant. Review the user's markdown document and answer any " +
		"questions the user may have. Reply in an encouraging tone, but be concise and never ask follow up " +
		"questions. Today's date: " + time.Now().Format(time.DateOnly)

	req := openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		},
		Stream: true,
	}

	stream, err := c.ai.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return stream, nil
}

// Response streams the response from the OpenAI API to stdout.
func (*Client) Response(ctx *cli.Context, stream *openai.ChatCompletionStream) error {
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Fprintf(ctx.App.Writer, "\n")

			return nil
		}

		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Fprint(ctx.App.Writer, resp.Choices[0].Delta.Content)
	}
}
