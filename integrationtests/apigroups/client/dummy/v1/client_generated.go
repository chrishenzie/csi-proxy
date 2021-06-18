// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	v1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	"google.golang.org/grpc"
)

// GroupName is the group name of this API.
const GroupName = "dummy"

// Version is the api version.
var Version = apiversion.NewVersionOrPanic("v1")

type Client struct {
	client     v1.DummyClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the dummy API group version v1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(GroupName, Version)
	return NewClientWithPipePath(pipePath)
}

// NewClientWithPipePath returns a client to make calls to the named pipe located at "pipePath".
// It's the caller's responsibility to Close the client when done.
func NewClientWithPipePath(pipePath string) (*Client, error) {

	// verify that the pipe exists
	_, err := winio.DialPipe(pipePath, nil)
	if err != nil {
		return nil, err
	}

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1.NewDummyClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1.DummyClient = &Client{}

func (w *Client) ComputeDouble(context context.Context, request *v1.ComputeDoubleRequest, opts ...grpc.CallOption) (*v1.ComputeDoubleResponse, error) {
	return w.client.ComputeDouble(context, request, opts...)
}

func (w *Client) TellMeAPoem(context context.Context, request *v1.TellMeAPoemRequest, opts ...grpc.CallOption) (*v1.TellMeAPoemResponse, error) {
	return w.client.TellMeAPoem(context, request, opts...)
}
