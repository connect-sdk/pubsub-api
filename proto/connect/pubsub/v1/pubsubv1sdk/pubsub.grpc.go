package pubsubv1sdk

import (
	"context"
	http "net/http"

	connect "connectrpc.com/connect"
	interceptor "github.com/connect-sdk/interceptor"
	middleware "github.com/connect-sdk/middleware"
	chi "github.com/go-chi/chi/v5"

	pubsubv1 "github.com/connect-sdk/pubsub-api/proto/connect/pubsub/v1"
	pubsubv1connect "github.com/connect-sdk/pubsub-api/proto/connect/pubsub/v1/pubsubv1connect"
)

var _ pubsubv1.PubsubServiceClient = &PubsubServiceClient{}

// PubsubServiceClient is a client for the google.pubsub.v1.PubsubService service.
type PubsubServiceClient struct {
	client pubsubv1connect.PubsubServiceClient
}

// NewPubsubServiceClient creates a new the google.pubsub.v1.PubsubServiceClient client.
func NewPubsubServiceClient(uri string, options ...PubsubServiceClientOption) pubsubv1.PubsubServiceClient {
	if uri == "" {
		return &pubsubv1.NopPubsubServiceClient{}
	}

	config := &PubsubServiceClientConfig{
		Client:        http.DefaultClient,
		ClientURL:     uri,
		ClientOptions: []connect.ClientOption{},
	}

	// Apply the options
	for _, opt := range options {
		opt.Apply(config)
	}

	var interceptors []connect.Interceptor
	// prepare the interceptors
	interceptors = append(interceptors, interceptor.WithTracer())
	interceptors = append(interceptors, interceptor.WithLogger())
	// prepare the configuration
	config.ClientOptions = append(config.ClientOptions, connect.WithInterceptors(interceptors...))

	client := pubsubv1connect.NewPubsubServiceClient(
		config.Client,
		config.ClientURL,
		config.ClientOptions...)

	return &PubsubServiceClient{client: client}
}

// PushPubsubMessage implements google.pubsub.v1.PubsubServiceClient.
func (x *PubsubServiceClient) PushPubsubMessage(ctx context.Context, m *pubsubv1.PubsubMessage) error {
	r := &pubsubv1.PushPubsubMessageRequest{
		Message: m,
	}

	_, err := x.client.PushPubsubMessage(ctx, connect.NewRequest(r))
	// done
	return err
}

var _ pubsubv1connect.PubsubServiceHandler = &PubsubServiceHandler{}

// PubsubServiceHandler represents an instance of google.pubsub.v1.PubsubServiceHandler handler.
type PubsubServiceHandler struct {
	// PubsubService contains an instance of google.pubsub.v1.PubsubService handler.
	PubsubService pubsubv1.PubsubService
}

// Mount mounts the controller to a given router.
func (x *PubsubServiceHandler) Mount(r chi.Router) {
	var interceptors []connect.Interceptor
	// prepare the interceptors
	interceptors = append(interceptors, interceptor.WithTracer())
	interceptors = append(interceptors, interceptor.WithLogger())
	interceptors = append(interceptors, interceptor.WithValidator())

	var options []connect.HandlerOption
	// prepare the options for interceptor collection
	options = append(options, connect.WithInterceptors(interceptors...))
	// prepare the options
	options = append(options, interceptor.WithRecover())

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithLogger())
		// create the handler
		path, handler := pubsubv1connect.NewPubsubServiceHandler(x, options...)
		// mount the handler
		r.Mount(path, handler)
	})
}

// PushMessage implements MessageServiceHandler.
func (x *PubsubServiceHandler) PushPubsubMessage(ctx context.Context, r *connect.Request[pubsubv1.PushPubsubMessageRequest]) (*connect.Response[pubsubv1.PushPubsubMessageResponse], error) {
	response, err := x.PubsubService.PushPubsubMessage(ctx, r.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(response), nil
}
