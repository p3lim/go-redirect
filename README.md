# go-redirecet

Simple redirect program written in Go that uses environment variables to dictate redirect target.

Built as a scratch container with minimal fluff.

### Usage:

Example redirecting `0.0.0.0:9000` -> `https://example.org:8443`

	docker run -d \
	    -e REDIRECT_TARGET=https://example.org:8443 \
	    -p 0.0.0.0:9000:8080/tcp \
	    ghcr.io/p3lim/go-redirect:latest

Variables:

- `REDIRECT_SOURCE` - address to listen on (default: "0.0.0.0:80", "0.0.0.0:8080" in the container image)
- `REDIRECT_TARGET` - where to redirect to, full URI
