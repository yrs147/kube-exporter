# Kubernetes Prometheus Exporter

This is a simple Prometheus Exporter which scrapes info from Kubernetes API and convert them into Prometheus metrics

To run the Program :- 

1) Before Running the Exporter Make sure your cluster is up and running
2) Set Path to your `KUBECONFIG` in the code
3) Run 
```
go run main.go
```
3) Then to see the Metrics , make a curl Request 
```
curl http://localhost:8000/metrics
```
