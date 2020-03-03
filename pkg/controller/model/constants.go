package model

import "time"

const (
	GrafanaConfigName          = "grafana-config"
	GrafanaLogVolumes          = "grafana-log"
	GrafanaDataVolumes         = "grafana-data"
	GrafanaDatasourceName      = "gafana-datasource"
	GrafanaHealthEndpoint      = "/api/health"
	DefaultGrafanaImage        = "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom/grafana"
	DefaultGrafanaImageTag     = "v6.5.2-build.1"
	RouterImage                = "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom/icp-management-ingress"
	RouterImageTag             = "2.5.0"
	DashboardImageTag          = "v1.2.0-build.2"
	DashboardImage             = "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom/dashboard-controller"
	InitContainerName          = "init-container"
	DefaultInitImage           = "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom/icp-initcontainer"
	DefaultInitImageTag        = "1.0.0-build.2"
	DefaultRouterPort          = 8080
	GrafanaAdminSecretName     = "grafana-secret"
	GrafanaInitMounts          = "grafana-init-mount"
	GrafanaPlugins             = "grafana-plugins"
	GrafanaSecretsDir          = "/etc/grafana-secrets/"
	GrafanaConfigMapsDir       = "/etc/grafana-configmaps/"
	GrafanaServiceAccountName  = "grafana-serviceaccount"
	GrafanaDeploymentName      = "grafana-deployment"
	GrafanaServiceName         = "grafana-service"
	GrafanaHttpPortName        = "grafana"
	RequeueDelay               = time.Second * 10
	DefaultGrafanaPort         = 3000
	GrafanaRouteName           = "grafana-route"
	GrafanaAdminUserEnvVar     = "username"
	GrafanaAdminPasswordEnvVar = "password"
	ClusterDomain              = "cluster.local"
	PrometheusPort             = 9090
)
