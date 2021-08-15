package client

type Client interface {
	FromFile(filePath string) (string, error)
	FromUrl(url string) (string, error)
}
