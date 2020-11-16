package templates

var OrgChartTpl = `
{{- define "org-chart" }}
<head>
	<style>
		#chart_div {
			width: 900px;
		}
	</style>
    <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
      google.charts.load('current', {packages:["orgchart"]});
      google.charts.setOnLoadCallback(drawChart);

      function drawChart() {

		var json = {{.jsonData}}
		var jsonData = JSON.parse(json)

		var dataTable = new google.visualization.DataTable();
        dataTable.addColumn('string', 'value');
        dataTable.addColumn('string', 'parent_value');
        dataTable.addColumn('string', 'row_type');

		for(i=0;i<jsonData.length;i++){
			var obj = jsonData[i];
			var rIndex = dataTable.addRow([obj.value, obj.parent_value, obj.row_type]);

			switch (obj.row_type) {
              case 'Company':
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid #7f7fbf; background: #7f7fbf');
                break;

              case 'Ubo Company':
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid #7f7fbf; background: #7f7fbf');
                break;

              case 'Director':
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid #c00; background: #c00');
                break;

              case 'Ubo':
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid navy; background: navy');
                break;
            }
		}

        var options = {
           'allowHtml': true
        };
        
        // Create the chart.
 		var chart_div = document.getElementById('chart_div');
        var chart = new google.visualization.OrgChart(document.getElementById('chart_div'));
        chart.draw(dataTable, options);
      }
    </script>
    </head>
  <body>
    <div id="chart_div"></div>
  </body>
{{ end }}
`

