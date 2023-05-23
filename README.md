# Kubernetes Prometheus Exporter

This is a simple Prometheus Exporter which scrapes info from Kubernetes API and convert them into Prometheus metrics

To run the Program :- 

1) Set Path to your `KUBECONFIG` in the code
2) Run 
```
go run main.go
```
3) Then to see the Metrics , make a curl Request 
```
curl http://localhost:8000/metrics
```
