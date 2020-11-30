package templates

var OrgChartTpl = `
{{- define "org-chart" }}
<head>
	<style>
		.director {
			  height: 25px;
			  width: 25px;
			  background-color: #c00;
			  border-radius: 50%;
			  display: inline-block;
		  }

		.ubo {
			  height: 25px;
			  width: 25px;
			  background-color: #4545A4;
			  border-radius: 50%;
			  display: inline-block;
		  }

		.company {
			  height: 25px;
			  width: 25px;
			  background-color: #7f7fbf;
			  border-radius: 50%;
			  display: inline-block;
		  }

		#legend {
			margin-left: 20px;
 		}

		hr { background-color: #3388dd; height: 1px; border: 0; }

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
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid #c00; background: #c00;');
                break;

              case 'Ubo':
                dataTable.setRowProperty(rIndex, 'style', 'border: 1px solid #4545A4; background: #4545A4');
                break;
            }
		}

        var options = {
           'allowHtml': true,
           'size': 'small'
        };
        
        // Create the chart.
 		var chart_div = document.getElementById('chart_div');
        var chart = new google.visualization.OrgChart(document.getElementById('chart_div'));
        chart.draw(dataTable, options);
      }
    </script>
	<style type="text/css">
	   .google-visualization-orgchart-table * {
			padding: 1.5px !important
		}

		.chart_div {
			display: block;
			width: 900px;
 			overflow: visible !important;
		}
</style>
    </head>
  <body>
    <div id="chart_div" class="chart_div" style="width: 900px;"></div>
	<hr/>
	<div id="legend">
	 <table cellspacing="0" cellpadding="0" style="width: 450px"> 
		<tr> 
			<td style="width: 30px;"> <span class="director"></span> </td> <td style="width: 70px;"> Director </td> 
			<td style="width: 30px;"> <span class="ubo"></span> </td> <td style="width: 70px;"> UBO </td> 
			<td style="width: 30px;"> <span class="company"></span> </td> <td style="width: 70px;"> Company </td> 
		</tr>
	 </table>
</div>
  </body>
{{ end }}
`

