package utils
import (
	"text/template"
	"bytes"
	b64 "encoding/base64"

	corev1 "k8s.io/api/core/v1"
	"github.com/IBM/ibm-grafana-operator/pkg/apis/operator/v1alpha1"
	_ "github.com/IBM/ibm-grafana-operator/pkg/artifacts"
)

const (
	// configmap name and file key
	FileKeys = []map[string][]interface{
		"grafana-lua-script-config": ["grafana.lua", GrafanaLuaScript],
		"util-lua-script-config": ["monitoring-util.lua", UtilLuaScript],
		"router-config": ["nginx.conf", RouterConfig],
		"grafana-crd-entry": ["run.sh", RouterEntry],
	}
	
	DashboardData := ["grafana-default-dashboards", "helm-release-dashboard.json",HelmReleaseDashboard], "kubenertes-pod-dashboard.json", KubernetesPodDashboard, "mcm-monitoring-dashboard.json", MCMMonitoringDashboard]

	Namespace = "openshift-cs-monitoring"
	clusterPort = ""
	Environment = "openshift"
	clusterDomain = ""
	// These should come from ingress setting.
	prometheusFullName = ""
	prometheusPort = 0
	grafanaFullName = ""
	grafanaPort = 0

)

func createConfigmap(name, fileKey, data string) corev1.ConfigMap {
	
	var data map[string]string
	data[fileKey] = data
	configmap := &corev1.ConfigMap{
		ObejctMeta: core.ObejctMeta{
			Name: name
		}
		Data: data
	}
	configmap.ObejctMeta.Labels["app"] = "grafana"
	return configmap
}

// createConfigmaps will create all the confimap for the grafana.
func CreateConfigmaps(FileKeys map[string]string) []corev1.ConfigMap{
	configmaps := []corev1.ConfigMaps{}

	type Data struct{
		Namespace string
		Environment string
		ClusterDomain string
		GrafanaFullName string
		PrometheusFullName string
		ClusterPort int32
		PrometheusPort int32
		GrafanaPort int32
	}

	tplData := Data {
		Namespace = "openshift-cs-monitoring"
		ClusterPort = ""
		Environment = "openshift"
		ClusterDomain = ""
		// These should come from ingress setting.
		PrometheusFullName = ""
		PrometheusPort = 0
		GrafanaFullName = ""
		GrafanaPort = 0
	}

	var buff bytes.Buffer
	for name, value := range FileKeys {
			FileKeys := value[0]
			tpl := value[1]
			data, err := tpl.Execute(&buff, tplData)
			if err != nil {
				panic(err)
			}
			configmaps = append(configmaps, createConfigmap(name, value, string(data)))
	}

	name = DashboardData[0]
	count := 1
	dashboardConfigData := map[string]string
	for count < len(DashboardData[1:]) {
		fileKey = DashboardData[count]
		tpl = DashboardData[count + 1]
		err = tpl.Execute(&buff, tplData)
		if err != nil {
			panic(err)
		}
		dashboardConfigData[fileKey] = string(buff)
		count = count + 2
	}
	dashboardConfigMap := corev1.ConfigMap{
		corev1.ObejctMeta{
			Name: name,
		},
		Data: dashboardConfigData
	}

	configmaps = append(configmaps, dashboardConfigMap)
	return configmaps
}

func CreateGrafanaSecret (cr *v1alpha1.Grafana) corev1.Secret {
	user := cr.Spec.Config.Security.AdminUser
	password := cr.Spec.Config.Security.AdminPassword

	encUser := b64.StdEncoding.EncodeToString([]byte(user))
	encPass := b64.StdEncoding.EncodeToString([]byte(password))
	return corev1.Secret{
		corev1.ObejctMeta {
			Name: "grafana-secret",
			Labels: map[string]string{"app":"grafana"}
		},
		corev1.SecretType: corev1.SecretTypeOpauqe,
		Data: map[string]string{"usernam":encUser, "password": encPass},
	}
}