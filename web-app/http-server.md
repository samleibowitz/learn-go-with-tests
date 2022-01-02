# HTTP Server

Here we're using a few things worth calling out.

- [`http.HandlerFunc`](https://pkg.go.dev/net/http#HandlerFunc) lets you use an ordinary old function as an HTTP request handler. Specifically, in this case, we're using the `PlayerServer` function to handle creating responses to web requests.
- [`http.ListenAndServe`](https://pkg.go.dev/net/http#ListenAndServe)  listens on a TCP port and passes incoming requests to that port to the designated handler. Neat.