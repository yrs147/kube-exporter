package main

import (
	"fmt"
	"net/http"
	"time"
	"context"

	corev1 "k8s.io/api/core/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	podsRunning = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "kubernetes_pods_running",
		Help: "Number of running pods",
	})
)

func main() {
	// Create a Prometheus registry and register the metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(podsRunning)

	// Create a Kubernetes clientset
	config, err := clientcmd.BuildConfigFromFlags("", "/home/yrs/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Start the metric scraping goroutine
	go scrapeKubernetesMetrics(clientset)

	// Serve the metrics on /metrics endpoint
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8000", nil)
}

func scrapeKubernetesMetrics(clientset *kubernetes.Clientset) {
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("Error retrieving Kubernetes metrics:", err)
			continue
		}

		runningPods := 0
		for _, pod := range pods.Items {
			if pod.Status.Phase == corev1.PodRunning {
				runningPods++
			}
		}

		podsRunning.Set(float64(runningPods))

		time.Sleep(5 * time.Second) // Scrape metrics every 5 seconds
	}
}

