package dishy

const defaultDashboard = `{
	  "uid": "%s",
	  "title": "%s",
	  "tags": [ "templated" ],
	  "timezone": "browser",
	  "refresh": "5s",
	  "schemaVersion": 6,
	  "version": 0,
	  "panels": [
		{
		  "collapsed": false,
		  "gridPos": {
			"h": 1,
			"w": 24,
			"x": 0,
			"y": 0
		  },
		  "id": 22,
		  "panels": [],
		  "title": "Speed",
		  "type": "row"
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
				"lineInterpolation": "smooth",
				"lineStyle": {
				  "fill": "solid"
				},
				"lineWidth": 3,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "never",
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 8,
			"w": 10,
			"x": 0,
			"y": 1
		  },
		  "id": 12,
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  upload_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Measured Upload Speed",
		  "type": "timeseries"
		},
		{
		  "description": "",
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "thresholds"
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 8,
			"w": 6,
			"x": 10,
			"y": 1
		  },
		  "id": 28,
		  "options": {
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"mean"
			  ],
			  "fields": "",
			  "values": false
			},
			"showThresholdLabels": false,
			"showThresholdMarkers": true
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  upload_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Average (Mean) Upload Speed",
		  "type": "gauge"
		},
		{
		  "description": "",
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "continuous-GrYlRd"
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 8,
			"w": 5,
			"x": 16,
			"y": 1
		  },
		  "id": 26,
		  "options": {
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"max"
			  ],
			  "fields": "/^upload_speed$/",
			  "values": false
			},
			"showThresholdLabels": false,
			"showThresholdMarkers": true
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  upload_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Fastest Upload Speed",
		  "type": "gauge"
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
				"lineInterpolation": "smooth",
				"lineWidth": 3,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "never",
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 7,
			"w": 10,
			"x": 0,
			"y": 9
		  },
		  "id": 10,
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  download_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
			  "refId": "A",
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
		  "title": "Measured Download Speed",
		  "type": "timeseries"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "thresholds"
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 7,
			"w": 6,
			"x": 10,
			"y": 9
		  },
		  "id": 30,
		  "options": {
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"mean"
			  ],
			  "fields": "",
			  "values": false
			},
			"showThresholdLabels": false,
			"showThresholdMarkers": true
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  download_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
			  "refId": "A",
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
		  "title": "Average (Mean) Download Speed",
		  "type": "gauge"
		},
		{
		  "description": "",
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "thresholds"
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
			  },
			  "unit": "bps"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 7,
			"w": 5,
			"x": 16,
			"y": 9
		  },
		  "id": 32,
		  "options": {
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"max"
			  ],
			  "fields": "",
			  "values": false
			},
			"showThresholdLabels": false,
			"showThresholdMarkers": true
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  download_speed\nFROM dishy_speed\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
			  "refId": "A",
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
		  "title": "Fastest Download Speed",
		  "type": "gauge"
		},
		{
		  "collapsed": false,
		  "gridPos": {
			"h": 1,
			"w": 24,
			"x": 0,
			"y": 16
		  },
		  "id": 16,
		  "panels": [],
		  "title": "Ping",
		  "type": "row"
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
			"h": 8,
			"w": 10,
			"x": 0,
			"y": 17
		  },
		  "id": 14,
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  pop_ping_latency_ms\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Average Latency",
		  "type": "timeseries"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "thresholds"
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
			"w": 6,
			"x": 10,
			"y": 17
		  },
		  "id": 18,
		  "options": {
			"colorMode": "value",
			"graphMode": "area",
			"justifyMode": "auto",
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"min"
			  ],
			  "fields": "",
			  "values": false
			},
			"textMode": "auto"
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  pop_ping_latency_ms\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Min Latency",
		  "type": "stat"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "thresholds"
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
			"w": 5,
			"x": 16,
			"y": 17
		  },
		  "id": 20,
		  "options": {
			"colorMode": "value",
			"graphMode": "area",
			"justifyMode": "auto",
			"orientation": "auto",
			"reduceOptions": {
			  "calcs": [
				"mean"
			  ],
			  "fields": "",
			  "values": false
			},
			"textMode": "auto"
		  },
		  "pluginVersion": "8.3.4",
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  pop_ping_latency_ms\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
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
		  "title": "Average (Mean) Latency",
		  "type": "stat"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "palette-classic"
			  },
			  "custom": {
				"hideFrom": {
				  "legend": false,
				  "tooltip": false,
				  "viz": false
				}
			  },
			  "mappings": []
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 9,
			"w": 5,
			"x": 0,
			"y": 25
		  },
		  "id": 6,
		  "options": {
			"legend": {
			  "displayMode": "list",
			  "placement": "bottom"
			},
			"pieType": "pie",
			"reduceOptions": {
			  "calcs": [
				"lastNotNull"
			  ],
			  "fields": "",
			  "values": false
			},
			"tooltip": {
			  "mode": "single"
			}
		  },
		  "pluginVersion": "8.3.4",
		  "targets": [
			{
			  "datasource": {
				"type": "postgres",
				"uid": "%s"
			  },
			  "format": "time_series",
			  "group": [],
			  "metricColumn": "software_version",
			  "rawQuery": false,
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  software_version AS metric,\n  downlink_throughput_bps\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1,2",
			  "refId": "A",
			  "select": [
				[
				  {
					"params": [
					  "downlink_throughput_bps"
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
		  "title": "Firmware Versions",
		  "type": "piechart"
		},
		{
		  "fieldConfig": {
			"defaults": {
			  "color": {
				"mode": "palette-classic"
			  },
			  "custom": {
				"hideFrom": {
				  "legend": false,
				  "tooltip": false,
				  "viz": false
				}
			  },
			  "mappings": []
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 9,
			"w": 5,
			"x": 5,
			"y": 25
		  },
		  "id": 4,
		  "options": {
			"legend": {
			  "displayMode": "list",
			  "placement": "bottom"
			},
			"pieType": "pie",
			"reduceOptions": {
			  "calcs": [
				"lastNotNull"
			  ],
			  "fields": "",
			  "values": false
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
			  "metricColumn": "hardware_version",
			  "rawQuery": false,
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  hardware_version AS metric,\n  downlink_throughput_bps\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1,2",
			  "refId": "A",
			  "select": [
				[
				  {
					"params": [
					  "downlink_throughput_bps"
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
		  "title": "Hardware Versions",
		  "type": "piechart"
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
				"lineInterpolation": "smooth",
				"lineWidth": 3,
				"pointSize": 5,
				"scaleDistribution": {
				  "type": "linear"
				},
				"showPoints": "never",
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
			  },
			  "unit": "percent"
			},
			"overrides": []
		  },
		  "gridPos": {
			"h": 9,
			"w": 11,
			"x": 10,
			"y": 25
		  },
		  "id": 36,
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
			  "rawSql": "SELECT\n  \"time\" AS \"time\",\n  pop_ping_drop_rate\nFROM dishy_data\nWHERE\n  $__timeFilter(\"time\") and dishy_id=\"%s\" \nORDER BY 1",
			  "refId": "A",
			  "select": [
				[
				  {
					"params": [
					  "pop_ping_drop_rate"
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
		  "title": "Ping Drop Rate",
		  "type": "timeseries"
		}
	  ]
	}`
