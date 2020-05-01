package opcua

import (
	"errors"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func chooseEndpoint(endpoint, policy, mode string) (*ua.EndpointDescription, error) {
	endpoints, err := opcua.GetEndpoints(endpoint)
	if err != nil {
		return nil, err
	}

	ep := opcua.SelectEndpoint(endpoints, policy, ua.MessageSecurityModeFromString(mode))
	if ep == nil {
		return nil, errors.New("Failed to find suitable endpoint")
	}

	return ep, nil
}

func InitOpcUaClient(ep *ua.EndpointDescription) {

}
