# go-redirect

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

### Synology DSM

I made this due to limitations of port bindings in Synology DSM, where port 80 and 443 are reserved for internal purposes. Thankfully, one of those purposes is a reverse proxy built right into the login portal. This is a quick guide on how to make Synology DSM respond to port 80 and 443.

1. Install Docker through the Package Center
2. In the Control Panel, under Terminal & SNMP, enable SSH
	- If you've enabled the firewall you will also need to allow this port
3. SSH into the Synology box, then run `sudo docker pull ghcr.io/p3lim/go-request:latest`
	- We have to do this because the Docker GUI doesn't support GitHub's Container Registry for some reason
	- You should now see the image in the Docker GUI
4. Create a new container through the "Create" wizard
	1. In the Image section, select the image we pulled
	2. In the Network section, use the default bridge network
	3. In the General section, give it a name and enable auto-restart
	4. In the General section, open up the Advanced Settings
		- Add a new environment variable `REDIRECT_TARGET` and set the value as you please, e.g. `https://synology.example.org`
	5. In the Port section, set the local port to `8080`
	6. Skip the Volume section
	7. Run the container after the wizard is done and click complete
5. In the Control Panel, under Login Portal, go to the Advanced section, and click "Reverse Proxy", where we'll create the following:
	- http -> container
		- Source Protocol: HTTP
		- Source Hostname: anything you want, e.g. `synology.example.org`
		- Source Port: 80
		- Destination Protocol: HTTP
		- Destination Hostname: localhost
		- Destination Port: 8080
	- https -> default https port
		- Source Protocol: HTTPS
		- Source Hostname: same as the previous one, e.g. `synology.example.org`
		- Source Port: 443
		- Destination Protocol: HTTPS
		- Destination Hostname: localhost
		- Destination Port: 5001

At this point, if you have your DNS set up to point `synology.example.org` to your Synology box, you'll be able to access it using the normal ports 80 and 443. I highly suggest you replace the default certificate.
