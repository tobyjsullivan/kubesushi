# Kubesushi

Launch Docker containers in a single, simple request. Still very much a work in progress.

## Running

1. Launch a Kubernetes cluster somewhere (like GKE)
2. Configure kubectl on your local machine.
3. `kubectl proxy`
4. `go run ./*.go`

```sh
curl --request POST \
  --url http://localhost:3000/deployment-requests \
  --header 'content-type: application/json' \
  --data '{
  "image": "hello-world"
}'
```

