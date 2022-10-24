package monobanksdk

import (
	"testing"
)

func TestGetCurrency(t *testing.T) {
	c := NewPublicClient()

	_, err := c.GetCurrency()

	if err != nil {
		t.Errorf("expected error nil, but actual error %v", err)
		return
	}
}

func TestClient_Statement(t *testing.T) {
	c := NewClient("Your TOKEN")

	_, err := c.Statement("YOUR ACCOUNT", 0, 0)

	if err != nil {
		t.Errorf("expected error nil, but actual error %v", err)
		return
	}
}
func TestClient_Webhook(t *testing.T) {
	c := NewClient("YOUR TOKEN")

	err := c.Webhook("webhook_URL")

	if err != nil {
		t.Errorf("expected error nil, but actual error %v", err)
		return
	}
}

func TestClient_ClientInfo(t *testing.T) {
	c := NewClient("YOUR TOKEN")

	_, err := c.ClientInfo()

	if err != nil {
		t.Errorf("expected error nil, but actual error %v", err)
		return
	}
}
