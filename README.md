# Elements Core Web UI
Simple server-side rendered web page to connect to your Elements Core instance


![Elements Core Web UI](https://raw.githubusercontent.com/tiero/elements-web/master/elements-ui-mockup.png)


# Usage

Run web server on port 8080

```bash
docker run -it --rm --name web -p 8080:8080 -e RPC_USER=elements -e RPC_PASS=elements -e RPC_HOST=localhost -e RPC_PORT=18884 ghcr.io/tiero/elements-web
```
