package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cybrota/go-a2a/a2a"
)

// CardResolver resolves Agent Card from a given URL
type CardResolver struct {
	BaseURL       string `json:"base_url"`
	AgentCardPath string `json:"agent_card_path" yaml:"agent_card_path" mapstructure:"agent_card_path"`
}

// GetAgentCard fetches an AgentCard from a well-known agent URL
func (c *CardResolver) GetAgentCard() (*a2a.AgentCard, error) {
	resp, err := http.Get(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http read body: %w", err)
	}

	var ac a2a.AgentCard
	err = json.Unmarshal(body, &ac)
	if err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}

	return &ac, nil
}

func NewCardResolver(BaseURL string, AgentCardPath string) *CardResolver {
	if AgentCardPath == "" {
		AgentCardPath = "./well-known/agent.json"
	}
	return &CardResolver{
		BaseURL:       BaseURL,
		AgentCardPath: AgentCardPath,
	}
}
