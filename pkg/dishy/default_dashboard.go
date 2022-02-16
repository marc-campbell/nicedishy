package dishy

const defaultDashboard = `{
	  "uid": "%s",
	  "title": "%s",
	  "tags": [ "templated" ],
	  "timezone": "browser",
	  "panels": [
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "palette-classic"
			  },
			  "custom": {
				"axisLabel": "",
				"axisPlacement": "auto",
				"barAlignment": 0,
				"drawStyle": "line",
				"fillOpacity": 0,
				"gradientMode": "none",
				"hideFrom": {
				  "legend": false,
				  "tooltip": false,
				  "viz": false
				},
				"lineInterpolation": "linear",
				"lineWidth": 1,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "auto",
				"spanNulls": false,
				"stacking": {
				  "group": "A",
				  "mode": "none"
				},
				"thresholdsStyle": {
				  "mode": "off"
				}
			  },
			  "mappings": [],
			  "thresholds": {
				"mode": "absolute",
				"steps": [
				  {
					"color": "green",
					"value": null
				  },
				  {
					"color": "red",
					"value": 80
				  }
				]
			  }
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 8,
			"w": 12,
			"x": 0,
			"y": 0
		  },
		  "id": 2,
		  "options": {
			"legend": {
			  "calcs": [],
			  "displayMode": "list",
			  "placement": "bottom"
			},
			"tooltip": {
			  "mode": "single"
			}
		  },
		  "targets": [
			{
			  "datasource": {
				"type": "postgres",
				"uid": "%s"
			  },
			  "format": "time_series",
			  "group": [],
			  "hide": false,
			  "metricColumn": "none",
			  "rawQuery": false,
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  download_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\")\nORDER BY 1",
			  "refId": "B",
			  "select": [
				[
				  {
					"params": [
					  "download_speed"
					],
					"type": "column"
				  }
				]
			  ],
			  "table": "dishy_speed",
			  "timeColumn": "\"time\"",
			  "timeColumnType": "timestamptz",
			  "where": [
				{
				  "name": "$__timeFilter",
				  "params": [],
				  "type": "macro"
				}
			  ]
			}
		  ],
		  "title": "Download Speed",
		  "type": "timeseries"
		},
		{
		  "description": "",
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "palette-classic"
			  },
			  "custom": {
				"axisLabel": "",
				"axisPlacement": "auto",
				"barAlignment": 0,
				"drawStyle": "line",
				"fillOpacity": 0,
				"gradientMode": "none",
				"hideFrom": {
				  "legend": false,
				  "tooltip": false,
				  "viz": false
				},
				"lineInterpolation": "linear",
				"lineWidth": 1,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "auto",
				"spanNulls": false,
				"stacking": {
				  "group": "A",
				  "mode": "none"
				},
				"thresholdsStyle": {
				  "mode": "off"
				}
			  },
			  "mappings": [],
			  "thresholds": {
				"mode": "absolute",
				"steps": [
				  {
					"color": "green",
					"value": null
				  },
				  {
					"color": "red",
					"value": 80
				  }
				]
			  }
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 8,
			"w": 12,
			"x": 12,
			"y": 0
		  },
		  "id": 4,
		  "options": {
			"legend": {
			  "calcs": [],
			  "displayMode": "list",
			  "placement": "bottom"
			},
			"tooltip": {
			  "mode": "single"
			}
		  },
		  "targets": [
			{
			  "datasource": {
				"type": "postgres",
				"uid": "%s"
			  },
			  "format": "time_series",
			  "group": [],
			  "metricColumn": "none",
			  "rawQuery": false,
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  upload_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\")\nORDER BY 1",
			  "refId": "A",
			  "select": [
				[
				  {
					"params": [
					  "upload_speed"
					],
					"type": "column"
				  }
				]
			  ],
			  "table": "dishy_speed",
			  "timeColumn": "\"time\"",
			  "timeColumnType": "timestamptz",
			  "where": [
				{
				  "name": "$__timeFilter",
				  "params": [],
				  "type": "macro"
				}
			  ]
			}
		  ],
		  "title": "Upload Speed",
		  "type": "timeseries"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "palette-classic"
			  },
			  "custom": {
				"axisLabel": "",
				"axisPlacement": "auto",
				"barAlignment": 0,
				"drawStyle": "line",
				"fillOpacity": 0,
				"gradientMode": "none",
				"hideFrom": {
				  "legend": false,
				  "tooltip": false,
				  "viz": false
				},
				"lineInterpolation": "linear",
				"lineWidth": 1,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "auto",
				"spanNulls": false,
				"stacking": {
				  "group": "A",
				  "mode": "none"
				},
				"thresholdsStyle": {
				  "mode": "off"
				}
			  },
			  "mappings": [],
			  "thresholds": {
				"mode": "absolute",
				"steps": [
				  {
					"color": "green",
					"value": null
				  },
				  {
					"color": "red",
					"value": 80
				  }
				]
			  }
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 17,
			"w": 24,
			"x": 0,
			"y": 8
		  },
		  "id": 6,
		  "options": {
			"legend": {
			  "calcs": [],
			  "displayMode": "list",
			  "placement": "bottom"
			},
			"tooltip": {
			  "mode": "single"
			}
		  },
		  "targets": [
			{
			  "datasource": {
				"type": "postgres",
				"uid": "%s"
			  },
			  "format": "time_series",
			  "group": [],
			  "metricColumn": "none",
			  "rawQuery": false,
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  pop_ping_latency_ms\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\")\nORDER BY 1",
			  "refId": "A",
			  "select": [
				[
				  {
					"params": [
					  "pop_ping_latency_ms"
					],
					"type": "column"
				  }
				]
			  ],
			  "table": "dishy_data",
			  "timeColumn": "\"time\"",
			  "timeColumnType": "timestamp",
			  "where": [
				{
				  "name": "$__timeFilter",
				  "params": [],
				  "type": "macro"
				}
			  ]
			}
		  ],
		  "title": "Ping / Latency",
		  "type": "timeseries"
		}
	  ],
	  "refresh": "5s",
	  "schemaVersion": 6,
	  "version": 0
	}`
