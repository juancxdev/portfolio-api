package api

func Start(port int) {
	r := routes()
	server := newServer(port, r)
	server.Start()
}
