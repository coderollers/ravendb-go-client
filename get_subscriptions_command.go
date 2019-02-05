package ravendb

import (
	"net/http"
	"strconv"
)

var (
	_ RavenCommand = &GetSubscriptionsCommand{}
)

// GetSubscriptionsCommand describes "delete subscription" command
type GetSubscriptionsCommand struct {
	RavenCommandBase

	start    int
	pageSize int

	Result []*SubscriptionState
}

func newGetSubscriptionsCommand(start int, pageSize int) *GetSubscriptionsCommand {
	cmd := &GetSubscriptionsCommand{
		RavenCommandBase: NewRavenCommandBase(),

		start:    start,
		pageSize: pageSize,
	}
	cmd.IsReadRequest = true
	return cmd
}

func (c *GetSubscriptionsCommand) createRequest(node *ServerNode) (*http.Request, error) {
	url := node.URL + "/databases/" + node.Database + "/subscriptions?start=" + strconv.Itoa(c.start) + "&pageSize=" + strconv.Itoa(c.pageSize)

	return NewHttpGet(url)
}

func (c *GetSubscriptionsCommand) setResponse(response []byte, fromCache bool) error {
	if len(response) == 0 {
		return nil
	}
	var res *GetSubscriptionsResult
	if err := jsonUnmarshal(response, &res); err != nil {
		return err
	}
	c.Result = res.Results
	return nil
}
