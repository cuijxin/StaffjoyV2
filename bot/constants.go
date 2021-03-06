package bot

const (
	// ServerPort tells the gRPC server what port to listen on
	ServerPort = ":1000"
	// Endpoint defines the DNS of the account server for clients
	// to access the server in Kubernetes.
	Endpoint = "botserver-service" + ServerPort
)
