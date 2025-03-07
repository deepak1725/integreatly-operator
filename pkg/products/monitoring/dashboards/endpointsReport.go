package monitoring

import (
	"github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
)

// This dashboard json is dynamically configured based on installation type (rhmi or rhoam)
// The installation name taken from the v1alpha1.RHMI.ObjectMeta.Name
func GetMonitoringGrafanaDBEndpointsReportJSON(installationName string) string {
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
	"iteration": 1561564976678,
	"links": [],
	"panels": [{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 0
			},
			"id": 6,
			"panels": [],
			"repeat": "services",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"title": "Uptime for $services",
			"type": "row"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": true,
			"colors": [
				"#d44a3a",
				"rgba(237, 129, 40, 0.89)",
				"rgb(255, 255, 255)"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "percentunit",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 1
			},
			"id": 4,
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
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "probe_success{service=\"$services\"}",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0.995,0.995",
			"title": "Uptime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "s",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 6,
				"y": 1
			},
			"id": 52,
			"interval": null,
			"links": [],
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
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "$__range_s - (probe_success{service=\"$services\"} * $__range_s)",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Downtime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
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
				"h": 2,
				"w": 6,
				"x": 12,
				"y": 1
			},
			"id": 39,
			"interval": null,
			"links": [],
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
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "ceil((changes(probe_success{service=~\"$services\"}[$__range]) / 2) + ((1 - (changes(probe_success{service=~\"$services\"}[$__range]) % 2)) * (1 - probe_success{service=~\"$services\"})))",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Outages",
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
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": false
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 18,
				"y": 1
			},
			"id": 25,
			"interval": null,
			"links": [],
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
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": true,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "probe_duration_seconds{service=~\"$services\"}",
				"format": "time_series",
				"instant": false,
				"interval": "",
				"intervalFactor": 1,
				"legendFormat": "",
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Response Time",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 3
			},
			"id": 53,
			"panels": [],
			"repeat": null,
			"repeatIteration": 1561564976678,
			"repeatPanelId": 6,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"title": "Uptime for $services",
			"type": "row"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": true,
			"colors": [
				"#d44a3a",
				"rgba(237, 129, 40, 0.89)",
				"rgb(255, 255, 255)"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "percentunit",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 4
			},
			"id": 54,
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 4,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "probe_success{service=\"$services\"}",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0.995,0.995",
			"title": "Uptime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "s",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 6,
				"y": 4
			},
			"id": 55,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 52,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "$__range_s - (probe_success{service=\"$services\"} * $__range_s)",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Downtime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
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
				"h": 2,
				"w": 6,
				"x": 12,
				"y": 4
			},
			"id": 56,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 39,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "ceil((changes(probe_success{service=~\"$services\"}[$__range]) / 2) + ((1 - (changes(probe_success{service=~\"$services\"}[$__range]) % 2)) * (1 - probe_success{service=~\"$services\"})))",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Outages",
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
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": false
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 18,
				"y": 4
			},
			"id": 57,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 25,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": true,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "probe_duration_seconds{service=~\"$services\"}",
				"format": "time_series",
				"instant": false,
				"interval": "",
				"intervalFactor": 1,
				"legendFormat": "",
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Response Time",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 6
			},
			"id": 58,
			"panels": [],
			"repeat": null,
			"repeatIteration": 1561564976678,
			"repeatPanelId": 6,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"title": "Uptime for $services",
			"type": "row"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": true,
			"colors": [
				"#d44a3a",
				"rgba(237, 129, 40, 0.89)",
				"rgb(255, 255, 255)"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "percentunit",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 7
			},
			"id": 59,
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 4,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "probe_success{service=\"$services\"}",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0.995,0.995",
			"title": "Uptime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "s",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 6,
				"y": 7
			},
			"id": 60,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 52,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "$__range_s - (probe_success{service=\"$services\"} * $__range_s)",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Downtime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
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
				"h": 2,
				"w": 6,
				"x": 12,
				"y": 7
			},
			"id": 61,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 39,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "ceil((changes(probe_success{service=~\"$services\"}[$__range]) / 2) + ((1 - (changes(probe_success{service=~\"$services\"}[$__range]) % 2)) * (1 - probe_success{service=~\"$services\"})))",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Outages",
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
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": false
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 18,
				"y": 7
			},
			"id": 62,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 25,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": true,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "probe_duration_seconds{service=~\"$services\"}",
				"format": "time_series",
				"instant": false,
				"interval": "",
				"intervalFactor": 1,
				"legendFormat": "",
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Response Time",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 9
			},
			"id": 63,
			"panels": [],
			"repeat": null,
			"repeatIteration": 1561564976678,
			"repeatPanelId": 6,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"title": "Uptime for $services",
			"type": "row"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": true,
			"colors": [
				"#d44a3a",
				"rgba(237, 129, 40, 0.89)",
				"rgb(255, 255, 255)"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "percentunit",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 10
			},
			"id": 64,
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 4,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "probe_success{service=\"$services\"}",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0.995,0.995",
			"title": "Uptime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 2,
			"format": "s",
			"gauge": {
				"maxValue": 1,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 6,
				"y": 10
			},
			"id": 65,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 52,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "$__range_s - (probe_success{service=\"$services\"} * $__range_s)",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Downtime",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": true,
			"colors": [
				"rgb(255, 255, 255)",
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
				"h": 2,
				"w": 6,
				"x": 12,
				"y": 10
			},
			"id": 66,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 39,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
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
				"expr": "ceil((changes(probe_success{service=~\"$services\"}[$__range]) / 2) + ((1 - (changes(probe_success{service=~\"$services\"}[$__range]) % 2)) * (1 - probe_success{service=~\"$services\"})))",
				"format": "time_series",
				"interval": "",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "1,1",
			"title": "Outages",
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
			"cacheTimeout": null,
			"colorBackground": false,
			"colorPrefix": false,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": false
			},
			"gridPos": {
				"h": 2,
				"w": 6,
				"x": 18,
				"y": 10
			},
			"id": 67,
			"interval": null,
			"links": [],
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
			"repeatIteration": 1561564976678,
			"repeatPanelId": 25,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "",
					"value": ""
				}
			},
			"sparkline": {
				"fillColor": "rgba(31, 118, 189, 0.18)",
				"full": true,
				"lineColor": "rgb(31, 120, 193)",
				"show": false
			},
			"tableColumn": "",
			"targets": [{
				"expr": "probe_duration_seconds{service=~\"$services\"}",
				"format": "time_series",
				"instant": false,
				"interval": "",
				"intervalFactor": 1,
				"legendFormat": "",
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Response Time",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "avg"
		}
	],
	"schemaVersion": 16,
	"style": "dark",
	"tags": [],
	"templating": {
		"list": [{
			"allValue": null,
			"current": {
				"text": "All",
				"value": "$__all"
			},
			"datasource": "Prometheus",
			"definition": "label_values(probe_success, service)",
			"hide": 0,
			"includeAll": true,
			"label": null,
			"multi": true,
			"name": "services",
			"options": [{
				"selected": true,
				"text": "All",
				"value": "$__all"
			}],
			"query": "label_values(probe_success, service)",
			"refresh": 1,
			"regex": "",
			"skipUrlSync": false,
			"sort": 0,
			"tagValuesQuery": "",
			"tags": [],
			"tagsQuery": "",
			"type": "query",
			"useTags": false
		}]
	},
	"time": {
		"from": "now-3h",
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
	"title": "Endpoints Report",
	"version": 19
}`
}
