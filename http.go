package duckcache

const defaultBasePath = "/_duck-cache/"

// HTTPPool implement PeerPicker for a pool of HTTP peers.
type HTTPPool struct {
	// this peer's base URL, e.g. "https://example.net:8080"
	self     string
	basePath string
}
