package ravendb

import (
	"net/http"
)

var (
	_ RavenCommand = &GetSubscriptionStateCommand{}
)

// GetSubscriptionStateCommand describes "get subscription state" command
type GetSubscriptionStateCommand struct {
	RavenCommandBase

	subscriptionName string

	Result *SubscriptionState
}

func newGetSubscriptionStateCommand(subscriptionName string) *GetSubscriptionStateCommand {
	cmd := &GetSubscriptionStateCommand{
		RavenCommandBase: NewRavenCommandBase(),

		subscriptionName: subscriptionName,
	}
	cmd.IsReadRequest = true
	return cmd
}

func (c *GetSubscriptionStateCommand) createRequest(node *ServerNode) (*http.Request, error) {
	url := node.URL + "/databases/" + node.Database + "/subscriptions/state?name=" + urlUtilsEscapeDataString(c.subscriptionName)

	return NewHttpGet(url)
}

func (c *GetSubscriptionStateCommand) setResponse(response []byte, fromCache bool) error {
	if len(response) == 0 {
		return throwInvalidResponse()
	}
	return jsonUnmarshal(response, &c.Result)
}
