package monitoring

import (
	"github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
)

// This dashboard json is dynamically configured based on installation type (rhmi or rhoam)
// The installation name taken from the v1alpha1.RHMI.ObjectMeta.Name
func GetMonitoringGrafanaDBEndpointsSummaryJSON(installationName string) string {
	quota := ``
	if installationName == resources.InstallationNames[string(v1alpha1.InstallationTypeManagedApi)] {
		quota = `,
			{
				"datasource": "Prometheus",
				"enable": true,
				"expr": "count by (quota,toQuota)(rhoam_quota{toQuota!=\"\"})",
				"hide": false,
				"iconColor": "#FADE2A",
				"limit": 100,
				"name": "Quota",
				"showIn": 0,
				"step": "",
				"tagKeys": "stage,quota,toQuota",
				"tags": "",
				"titleFormat": "Quota Change (million per day)",
				"type": "tags",
				"useValueForTime": false
			}`
	}
	return `{
	"annotations": {
		"list": [{
				"builtIn": 1,
				"datasource": "-- Grafana --",
				"enable": true,
				"hide": true,
				"iconColor": "rgba(0, 211, 255, 1)",
				"name": "Annotations & Alerts",
				"type": "dashboard"
			},
			{
				"datasource": "Prometheus",
				"enable": true,
				"expr": "count by (stage,version,to_version)(` + installationName + `_version{to_version!=\"\"})",
				"hide": false,
				"iconColor": "#FADE2A",
				"limit": 100,
				"name": "Upgrade",
				"showIn": 0,
				"step": "",
				"tagKeys": "stage,version,to_version",
				"tags": "",
				"titleFormat": "Upgrade",
				"type": "tags",
				"useValueForTime": false
			}` + quota + `
		]
	},
	"editable": true,
	"gnetId": null,
	"graphTooltip": 0,
	"iteration": 1561549685989,
	"links": [],
	"panels": [{
			"cacheTimeout": null,
			"colorBackground": true,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"format": "none",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 6,
				"w": 24,
				"x": 0,
				"y": 0
			},
			"id": 12,
			"interval": null,
			"links": [{
				"dashboard": "Endpoints Detailed",
				"title": "Drill Down",
				"type": "dashboard",
				"url": "/d/xtkCtBkiz2/endpoints-detailed"
			}],
			"mappingType": 1,
			"mappingTypes": [{
					"name": "value to text",
					"value": 1
				},
				{
					"name": "range to text",
					"value": 2
				}
			],
			"maxDataPoints": 100,
			"nullPointMode": "connected",
			"nullText": null,
			"postfix": "",
			"postfixFontSize": "50%",
			"prefix": "",
			"prefixFontSize": "50%",
			"rangeMaps": [{
				"from": "null",
				"text": "N/A",
				"to": "null"
			}],
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": false,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "count(probe_success) - sum(probe_success)",
				"format": "time_series",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Total Endpoints DOWN",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "current"
		},
		{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 6
			},
			"id": 4,
			"panels": [],
			"repeat": "services",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "fuse",
					"value": "fuse"
				}
			},
			"title": "$services UP/DOWN Status",
			"type": "row"
		},
		{
			"cacheTimeout": null,
			"colorBackground": true,
			"colorValue": false,
			"colors": [
				"#d44a3a",
				"rgba(237, 129, 40, 0.89)",
				"#299c46"
			],
			"datasource": "Prometheus",
			"format": "none",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 3,
				"w": 24,
				"x": 0,
				"y": 7
			},
			"id": 2,
			"interval": null,
			"links": [{
				"dashboard": "Endpoints Detailed",
				"includeVars": true,
				"title": "Drill Down",
				"type": "dashboard",
				"url": "/d/xtkCtBkiz2/endpoints-detailed"
			}],
			"mappingType": 1,
			"mappingTypes": [{
					"name": "value to text",
					"value": 1
				},
				{
					"name": "range to text",
					"value": 2
				}
			],
			"maxDataPoints": 100,
			"nullPointMode": "connected",
			"nullText": null,
			"postfix": "",
			"postfixFontSize": "50%",
			"prefix": "",
			"prefixFontSize": "50%",
			"rangeMaps": [{
				"from": "null",
				"text": "N/A",
				"to": "null"
			}],
			"repeat": null,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "fuse",
					"value": "fuse"
				}
			},
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": false,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "probe_success{service=~\"$services\"}",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "$services",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
					"op": "=",
					"text": "N/A",
					"value": "null"
				},
				{
					"op": "=",
					"text": "UP",
					"value": "1"
				},
				{
					"op": "=",
					"text": "DOWN",
					"value": "0"
				}
			],
			"valueName": "current"
		}
	],
	"refresh": "10s",
	"schemaVersion": 16,
	"style": "dark",
	"tags": [],
	"templating": {
		"list": [{
				"auto": true,
				"auto_count": 10,
				"auto_min": "10s",
				"current": {
					"text": "30s",
					"value": "30s"
				},
				"hide": 0,
				"label": "Interval",
				"name": "interval",
				"options": [{
						"selected": false,
						"text": "auto",
						"value": "$__auto_interval_interval"
					},
					{
						"selected": false,
						"text": "5s",
						"value": "5s"
					},
					{
						"selected": false,
						"text": "10s",
						"value": "10s"
					},
					{
						"selected": true,
						"text": "30s",
						"value": "30s"
					},
					{
						"selected": false,
						"text": "1m",
						"value": "1m"
					},
					{
						"selected": false,
						"text": "10m",
						"value": "10m"
					},
					{
						"selected": false,
						"text": "30m",
						"value": "30m"
					},
					{
						"selected": false,
						"text": "1h",
						"value": "1h"
					},
					{
						"selected": false,
						"text": "6h",
						"value": "6h"
					},
					{
						"selected": false,
						"text": "12h",
						"value": "12h"
					},
					{
						"selected": false,
						"text": "1d",
						"value": "1d"
					},
					{
						"selected": false,
						"text": "7d",
						"value": "7d"
					},
					{
						"selected": false,
						"text": "14d",
						"value": "14d"
					},
					{
						"selected": false,
						"text": "30d",
						"value": "30d"
					}
				],
				"query": "5s,10s,30s,1m,10m,30m,1h,6h,12h,1d,7d,14d,30d",
				"refresh": 2,
				"skipUrlSync": false,
				"type": "interval"
			},
			{
				"allValue": null,
				"current": {
					"text": "All",
					"value": [
						"$__all"
					]
				},
				"datasource": "Prometheus",
				"definition": "label_values(probe_success, service)",
				"hide": 0,
				"includeAll": true,
				"label": null,
				"multi": true,
				"name": "services",
				"options": [],
				"query": "label_values(probe_success, service)",
				"refresh": 1,
				"regex": "",
				"skipUrlSync": false,
				"sort": 1,
				"tagValuesQuery": "",
				"tags": [],
				"tagsQuery": "",
				"type": "query",
				"useTags": false
			}
		]
	},
	"time": {
		"from": "now-1h",
		"to": "now"
	},
	"timepicker": {
		"refresh_intervals": [
			"5s",
			"10s",
			"30s",
			"1m",
			"5m",
			"15m",
			"30m",
			"1h",
			"2h",
			"1d"
		],
		"time_options": [
			"5m",
			"15m",
			"1h",
			"6h",
			"12h",
			"24h",
			"2d",
			"7d",
			"30d"
		]
	},
	"timezone": "",
	"title": "Endpoints Summary",
	"uid": "hZJ_054Zk",
	"version": 5
}`
}
