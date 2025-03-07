package monitoring

import (
	"github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
)

// This dashboard json is dynamically configured based on installation type (rhmi or rhoam)
// The installation name taken from the v1alpha1.RHMI.ObjectMeta.Name
func GetMonitoringGrafanaDBEndpointsDetailedJSON(installationName string) string {
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
	"gnetId": 5345,
	"graphTooltip": 0,
	"iteration": 1561558894526,
	"links": [],
	"panels": [{
			"collapsed": false,
			"gridPos": {
				"h": 1,
				"w": 24,
				"x": 0,
				"y": 0
			},
			"id": 15,
			"panels": [],
			"repeat": "services",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
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
				"w": 6,
				"x": 0,
				"y": 1
			},
			"id": 2,
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
			"minSpan": 3,
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
			"repeatDirection": "h",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
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
		},
		{
			"cacheTimeout": null,
			"colorBackground": false,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 3,
				"x": 6,
				"y": 1
			},
			"id": 23,
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
					"text": "x",
					"value": "x"
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
				"expr": "avg(probe_duration_seconds{service=~\"$services\"})",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Average Probe Duration",
			"type": "singlestat",
			"valueFontSize": "50%",
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
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"format": "s",
			"gauge": {
				"maxValue": 100,
				"minValue": 0,
				"show": false,
				"thresholdLabels": false,
				"thresholdMarkers": true
			},
			"gridPos": {
				"h": 2,
				"w": 3,
				"x": 15,
				"y": 1
			},
			"id": 24,
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
					"text": "x",
					"value": "x"
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
				"expr": "avg(probe_dns_lookup_time_seconds{service=~\"$services\"})",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "",
			"title": "Average DNS Lookup",
			"type": "singlestat",
			"valueFontSize": "50%",
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueName": "current"
		},
		{
			"aliasColors": {},
			"bars": false,
			"dashLength": 10,
			"dashes": false,
			"datasource": "Prometheus",
			"fill": 1,
			"gridPos": {
				"h": 6,
				"w": 9,
				"x": 6,
				"y": 3
			},
			"id": 17,
			"legend": {
				"avg": false,
				"current": false,
				"max": false,
				"min": false,
				"rightSide": false,
				"show": true,
				"total": false,
				"values": false
			},
			"lines": true,
			"linewidth": 1,
			"links": [],
			"nullPointMode": "null",
			"percentage": false,
			"pointradius": 5,
			"points": false,
			"renderer": "flot",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
				}
			},
			"seriesOverrides": [{
				"alias": "UP/DOWN",
				"yaxis": 2
			}],
			"spaceLength": 10,
			"stack": false,
			"steppedLine": false,
			"targets": [{
					"expr": "probe_duration_seconds{service=~\"$services\"}",
					"format": "time_series",
					"interval": "$interval",
					"intervalFactor": 1,
					"legendFormat": "seconds",
					"refId": "A"
				},
				{
					"expr": "probe_success{service=~\"$services\"}",
					"format": "time_series",
					"instant": false,
					"interval": "$interval",
					"intervalFactor": 1,
					"legendFormat": "UP/DOWN",
					"refId": "B"
				}
			],
			"thresholds": [],
			"timeFrom": null,
			"timeRegions": [],
			"timeShift": null,
			"title": "Probe Duration",
			"tooltip": {
				"shared": true,
				"sort": 0,
				"value_type": "individual"
			},
			"type": "graph",
			"xaxis": {
				"buckets": null,
				"mode": "time",
				"name": null,
				"show": true,
				"values": []
			},
			"yaxes": [{
					"format": "s",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				},
				{
					"decimals": 0,
					"format": "short",
					"label": "",
					"logBase": 1,
					"max": "1",
					"min": "0",
					"show": true
				}
			],
			"yaxis": {
				"align": false,
				"alignLevel": null
			}
		},
		{
			"aliasColors": {},
			"bars": false,
			"dashLength": 10,
			"dashes": false,
			"datasource": "Prometheus",
			"fill": 1,
			"gridPos": {
				"h": 6,
				"w": 9,
				"x": 15,
				"y": 3
			},
			"id": 21,
			"legend": {
				"avg": false,
				"current": false,
				"max": false,
				"min": false,
				"show": true,
				"total": false,
				"values": false
			},
			"lines": true,
			"linewidth": 1,
			"links": [],
			"nullPointMode": "null",
			"percentage": false,
			"pointradius": 5,
			"points": false,
			"renderer": "flot",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
				}
			},
			"seriesOverrides": [],
			"spaceLength": 10,
			"stack": false,
			"steppedLine": false,
			"targets": [{
				"expr": "probe_dns_lookup_time_seconds{service=~\"$services\"}",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"legendFormat": "seconds",
				"refId": "A"
			}],
			"thresholds": [],
			"timeFrom": null,
			"timeRegions": [],
			"timeShift": null,
			"title": "DNS Lookup",
			"tooltip": {
				"shared": true,
				"sort": 0,
				"value_type": "individual"
			},
			"type": "graph",
			"xaxis": {
				"buckets": null,
				"mode": "time",
				"name": null,
				"show": true,
				"values": []
			},
			"yaxes": [{
					"format": "s",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				},
				{
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				}
			],
			"yaxis": {
				"align": false,
				"alignLevel": null
			}
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
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 4
			},
			"id": 18,
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
			"minSpan": 3,
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
			"repeatDirection": "h",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
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
				"expr": "probe_http_ssl{service=~\"$services\"}",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0,1",
			"title": "SSL",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
					"op": "=",
					"text": "N/A",
					"value": "null"
				},
				{
					"op": "=",
					"text": "YES",
					"value": "1"
				},
				{
					"op": "=",
					"text": "NO",
					"value": "0"
				}
			],
			"valueName": "current"
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
			"decimals": 2,
			"format": "dtdurations",
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
				"x": 0,
				"y": 6
			},
			"id": 19,
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
			"minSpan": 3,
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
			"repeatDirection": "h",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
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
				"expr": "probe_ssl_earliest_cert_expiry{service=~\"$services\"}-time()",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "0,1209600",
			"title": "SSL Cert Expiry",
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
					"op": "=",
					"text": "N/A",
					"value": "null"
				},
				{
					"op": "=",
					"text": "YES",
					"value": "1"
				},
				{
					"op": "=",
					"text": "NO",
					"value": "0"
				}
			],
			"valueName": "current"
		},
		{
			"cacheTimeout": null,
			"colorBackground": true,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
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
				"x": 0,
				"y": 8
			},
			"id": 20,
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
			"minSpan": 3,
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
			"repeatDirection": "h",
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "sso",
					"value": "sso"
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
				"expr": "probe_http_status_code{service=~\"$services\"}",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "300,400",
			"title": "HTTP Status Code",
			"transparent": false,
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
					"op": "=",
					"text": "N/A",
					"value": "null"
				},
				{
					"op": "=",
					"text": "YES",
					"value": "1"
				},
				{
					"op": "=",
					"text": "NO",
					"value": "0"
				}
			],
			"valueName": "current"
		},
		{
			"backgroundColor": "rgba(128,128,128,0.1)",
			"colorMaps": [{
				"color": "#CCC",
				"text": "N/A"
			}],
			"crosshairColor": "#8F070C",
			"datasource": "Prometheus",
			"display": "timeline",
			"expandFromQueryS": 0,
			"extendLastValue": true,
			"gridPos": {
				"h": 3,
				"w": 9,
				"x": 6,
				"y": 9
			},
			"highlightOnMouseover": true,
			"id": 55,
			"legendSortBy": "-ms",
			"lineColor": "rgba(0,0,0,0.1)",
			"links": [],
			"metricNameColor": "#000000",
			"rangeMaps": [{
				"from": "null",
				"text": "N/A",
				"to": "null"
			}],
			"rowHeight": 25,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
				}
			},
			"showDistinctCount": true,
			"showLegend": true,
			"showLegendCounts": true,
			"showLegendNames": false,
			"showLegendPercent": true,
			"showLegendValues": true,
			"showTimeAxis": true,
			"showTransitionCount": true,
			"targets": [{
				"expr": "probe_http_status_code{service=~\"$services\"}",
				"format": "time_series",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"textSize": 14,
			"textSizeTime": 12,
			"timeOptions": [{
					"name": "Years",
					"value": "years"
				},
				{
					"name": "Months",
					"value": "months"
				},
				{
					"name": "Weeks",
					"value": "weeks"
				},
				{
					"name": "Days",
					"value": "days"
				},
				{
					"name": "Hours",
					"value": "hours"
				},
				{
					"name": "Minutes",
					"value": "minutes"
				},
				{
					"name": "Seconds",
					"value": "seconds"
				},
				{
					"name": "Milliseconds",
					"value": "milliseconds"
				}
			],
			"timePrecision": {
				"name": "Minutes",
				"value": "minutes"
			},
			"timeTextColor": "#d8d9da",
			"title": "Probe Status",
			"type": "natel-discrete-panel",
			"units": "short",
			"useTimePrecision": false,
			"valueMaps": [{
				"op": "=",
				"text": "N/A",
				"value": "null"
			}],
			"valueTextColor": "#000000",
			"writeAllValues": true,
			"writeLastValue": false,
			"writeMetricNames": false
		},
		{
			"cacheTimeout": null,
			"colorBackground": true,
			"colorPostfix": false,
			"colorPrefix": false,
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
				"h": 2,
				"w": 6,
				"x": 0,
				"y": 10
			},
			"id": 53,
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
					"text": "x",
					"value": "x"
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
			"colorBackground": true,
			"colorValue": false,
			"colors": [
				"#299c46",
				"rgba(237, 129, 40, 0.89)",
				"#d44a3a"
			],
			"datasource": "Prometheus",
			"decimals": 0,
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
				"x": 0,
				"y": 32
			},
			"id": 75,
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
			"minSpan": 3,
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
			"repeatDirection": "h",
			"repeatIteration": 1561558894526,
			"repeatPanelId": 20,
			"repeatedByRow": true,
			"scopedVars": {
				"services": {
					"selected": false,
					"text": "x",
					"value": "x"
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
				"expr": "probe_http_status_code{service=~\"$services\"}",
				"format": "time_series",
				"interval": "$interval",
				"intervalFactor": 1,
				"refId": "A"
			}],
			"thresholds": "300,400",
			"title": "HTTP Status Code",
			"transparent": false,
			"type": "singlestat",
			"valueFontSize": "80%",
			"valueMaps": [{
					"op": "=",
					"text": "N/A",
					"value": "null"
				},
				{
					"op": "=",
					"text": "YES",
					"value": "1"
				},
				{
					"op": "=",
					"text": "NO",
					"value": "0"
				}
			],
			"valueName": "current"
		}
	],
	"refresh": "30s",
	"schemaVersion": 16,
	"style": "dark",
	"tags": [
		"blackbox",
		"prometheus"
	],
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
					"selected": false,
					"tags": [],
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
				"sort": 0,
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
	"title": "Endpoints Detailed",
	"uid": "xtkCtBkiz2",
	"version": 14
}`
}
