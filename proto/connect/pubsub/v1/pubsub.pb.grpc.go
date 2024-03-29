package pubsubv1

import "context"

//go:generate counterfeiter -generate

//counterfeiter:generate -o ./pubsubv1fake . PubsubService

// PubsubService is an implementation of the google.pubsub.v1.PubsubService service.
type PubsubService interface {
	// PushPubsubMessage processes a given message.
	PushPubsubMessage(context.Context, *PushPubsubMessageRequest) (*PushPubsubMessageResponse, error)
}

//counterfeiter:generate -o ./pubsubv1fake . PubsubServiceClient

// PubsusbServiceClient is an implementation of the google.pubsub.v1.PubsubServiceClient client.
type PubsubServiceClient interface {
	// PushPubsubMessage pushes a given message to google.pubsub.v1.PubsubService service.
	PushPubsubMessage(context.Context, *PubsubMessage) error
}

var _ PubsubServiceClient = &NopPubsubServiceClient{}

// NopPubsubServiceClient is a no-op client for the google.pubsub.v1.PubsubService service.
type NopPubsubServiceClient struct{}

// PushPubsubMessage implements pubsubv1.PubsubServiceClient.
func (*NopPubsubServiceClient) PushPubsubMessage(context.Context, *PubsubMessage) error {
	return nil
}
