package localmetrics

import (
	"fmt"
	"net/http"
	"time"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("localmetrics")

const (
	apiEndpoint = "https://apis/app.siliconpod.com/valpha1"
)

var (
	MetricSiliconPodCreateFailure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "silicon_pod_create_failure",
		Help:        "Metric for the failure of creating a cluster deployment.",
		ConstLabels: prometheus.Labels{"name": "silicon-pod-operator"},
	}, []string{"clusterdeployment_name"})

	MetricSiliconPodDeleteFailure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "silicon_pod_delete_failure",
		Help:        "Metric for the failure of deleting a cluster deployment.",
		ConstLabels: prometheus.Labels{"name": "silicon-pod-operator"},
	}, []string{"clusterdeployment_name"})

	MetricSiliconPodHeartBeat = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:        "silicon_pod_heartbeat",
		Help:        "Metric for heartbeating of the silicon pod api.",
		ConstLabels: prometheus.Labels{"name": "silicon-pod-operator"},
	})

	MetricsList = []prometheus.Collector{
		MetricSiliconPodCreateFailure,
		MetricSiliconPodDeleteFailure,
		MetricSiliconPodHeartBeat,
	}
)

// UpdateAPIMetrics updates all API endpoint metrics ever 5 minutes
func UpdateAPIMetrics(timer *prometheus.Timer) {
	d := time.Tick(5 * time.Minute)
	for range d {
		UpdateMetricPagerDutyHeartbeat(timer)
	}

}

// UpdateMetricSiliconPodCreateFailure updates guage to 1 when creation fails
func UpdateMetricSiliconPodCreateFailure(x int, cd string) {
	MetricPagerDutyCreateFailure.With(prometheus.Labels{
		"clusterdeployment_name": cd}).Set(
		float64(x))
}

// UpdateMetricSiliconPodDeleteFailure updates guage to 1 when deletion fails
func UpdateMetricSiliconPodDeleteFailure(x int, cd string) {
	MetricPagerDutyDeleteFailure.With(prometheus.Labels{
		"clusterdeployment_name": cd}).Set(
		float64(x))
}

// UpdateMetricSiliconPodHeartBeat curls the  API, updates the gauge to 1
// when successful.
func UpdateMetricSiliconPodHeartBeat(timer *prometheus.Timer) {
	metricLogger := log.WithValues("Namespace", "silicon-pod-operator")
	metricLogger.Info("Metrics for silicon-pod-operator API")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		metricLogger.Error(err, "Failed to reach api when authenticated")
		MetricPagerDutyHeartbeat.Observe(
			float64(timer.ObserveDuration().Seconds()))

		return
	}
	defer resp.Body.Close()

	// if there is an api key make an authenticated called
	if APIKey != "" {
		req, _ := http.NewRequest("GET", apiEndpoint, nil)
		req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Token token=%s", APIKey))
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			metricLogger.Error(err, "Failed to reach api when authenticated")
			MetricPagerDutyHeartbeat.Observe(
				float64(timer.ObserveDuration().Seconds()))

			return
		}
		defer resp.Body.Close()

	}
	MetricSiliconPodHeartBeat.Observe(float64(0))
}
