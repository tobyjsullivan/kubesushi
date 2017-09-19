package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/urfave/negroni"
	"os"
	"fmt"
	"encoding/json"
	"crypto/rand"
	"math/big"
	"strings"
)



func main() {
	r := mux.NewRouter()

	r.HandleFunc("/deployment-requests", handleDeploymentRequest).Methods(http.MethodPost)

	n := negroni.New()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Now server is running on port %s\n", port)
	n.Run(":"+port)
}

func handleDeploymentRequest(w http.ResponseWriter, r *http.Request) {
	var parsedReq struct{
		Image string `json:"image"`
	}

	if err := json.NewDecoder(r.Body).Decode(&parsedReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endpointUrl := "http://127.0.0.1:8001/apis/apps/v1beta1/namespaces/default/deployments"

	id, _ := rand.Int(rand.Reader, big.NewInt(10000))
	deploymentName := fmt.Sprintf("kubesushi-dep-%d", id.Int64())

	reqBody := fmt.Sprintf(`
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s
        image: %s
`, deploymentName, deploymentName, deploymentName, parsedReq.Image)

	req, err := http.NewRequest(http.MethodPost, endpointUrl, strings.NewReader(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-type", "application/yaml")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	

	w.WriteHeader(res.StatusCode)
	w.Write([]byte("success"))
}
