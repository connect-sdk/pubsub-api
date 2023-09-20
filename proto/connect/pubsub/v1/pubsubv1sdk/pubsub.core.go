package pubsubv1sdk

import (
	context "context"
	slog "log/slog"
	http "net/http"

	connect "connectrpc.com/connect"
	idtoken "google.golang.org/api/idtoken"
)

// PubsubServiceClientConfig represents the config for google.pubsub.v1.PubsubServiceClient client.
type PubsubServiceClientConfig struct {
	Client        *http.Client
	ClientURL     string
	ClientOptions []connect.ClientOption
}

// PubsubServiceClientOption represents an option for google.pubsub.v1.PubsubServiceClient client.
type PubsubServiceClientOption interface {
	// Apply applies the option.
	Apply(config *PubsubServiceClientConfig)
}

var _ PubsubServiceClientOption = PubsubServiceClientOptionFunc(nil)

// PubsubServiceClientOptionFunc represent a function that implementes google.pubsub.v1.PubsubServiceClientOption option.
type PubsubServiceClientOptionFunc func(*PubsubServiceClientConfig)

// Apply applies the option.
func (fn PubsubServiceClientOptionFunc) Apply(config *PubsubServiceClientConfig) {
	fn(config)
}

// PubsubServiceClientWithAuthorization enables an oidc authorization.
func PubsubServiceClientWithAuthorization() PubsubServiceClientOption {
	fn := func(config *PubsubServiceClientConfig) {
		// client uri
		uri := config.ClientURL
		// prepare the options
		options := []idtoken.ClientOption{}
		options = append(options, idtoken.WithHTTPClient(config.Client))
		// prepare the client
		client, err := idtoken.NewClient(context.Background(), uri, options...)
		if err != nil {
			slog.Error("unable to create an id token", err)
		}
		// set the client
		config.Client = client
	}

	return PubsubServiceClientOptionFunc(fn)
}

// PubsubServiceClientWithProtocol enables a given protocol.
func PubsubServiceClientWithProtocol(name string) PubsubServiceClientOption {
	fn := func(config *PubsubServiceClientConfig) {
		// prepare the protocol
		switch name {
		case "grpc":
			config.ClientOptions = append(config.ClientOptions, connect.WithGRPC())
		case "grpc+web":
			config.ClientOptions = append(config.ClientOptions, connect.WithGRPCWeb())
		}
	}

	return PubsubServiceClientOptionFunc(fn)
}

// PubsubServiceClientWithCodec enables a given codec.
func PubsubServiceClientWithCodec(name string) PubsubServiceClientOption {
	fn := func(config *PubsubServiceClientConfig) {
		// prepare the protocol
		switch name {
		case "json":
			config.ClientOptions = append(config.ClientOptions, connect.WithProtoJSON())
		}
	}

	return PubsubServiceClientOptionFunc(fn)
}
