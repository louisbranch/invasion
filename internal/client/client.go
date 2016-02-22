package client

type Client interface {
	Write([]byte) error
}
