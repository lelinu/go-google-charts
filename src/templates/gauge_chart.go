package templates

var GaugeTpl = `
{{- define "gauge-chart" }}
<head>
  <meta charset="utf-8">
  <script src="https://unpkg.com/chart.js@2.8.0/dist/Chart.bundle.js"></script>
  <script src="https://unpkg.com/chartjs-gauge@0.2.0/dist/chartjs-gauge.js"></script>
  <script src="https://unpkg.com/chartjs-plugin-datalabels@0.7.0/dist/chartjs-plugin-datalabels.js"></script>
  <script>
    var riskAssessmentValue = {{.riskAssessmentValue}}
    var config = {
      type: 'gauge',
      data: {
        labels: ['Low', 'Medium', 'High', 'Very High'],
        datasets: [{
        data: [0.5, 1.5, 2.5, 5],
        value: riskAssessmentValue,
         backgroundColor: ['green', 'yellow', 'orange', 'red'],
          borderWidth: 0
        }]
      },
      options: {
        responsive: true,
        title: {
          display: false,
          text: ''
        },
        layout: {
          padding: {
            bottom: 0
          }
        },
        needle: {
          radiusPercentage: 2,
          widthPercentage: 3.2,
          lengthPercentage: 60,
          color: 'rgba(0, 0, 0, 1)'
        },
        valueLabel: {
          display: false
        },
        plugins: {
          datalabels: {
            display: true,
            formatter:  function (value, context) {
              return context.chart.data.labels[context.dataIndex];
            },
            color: 'rgba(0, 0, 0, 1.0)',
            backgroundColor: null,
            font: {
              size: 11,
              weight: 'bold'
            }
          }
        }
      }
    };

    window.onload = function() {
        var ctx = document.getElementById('canvas-chart').getContext('2d');
        ctx.scale(5,5)
        window.myGauge = new Chart(ctx, config);
    };

  </script>
</head>

<body>
 <div id="outer-div" style="width:250px;height:125px;">
     <div id="canvas-holder" style="width:100%; object-fit: contain">
         <canvas style="display: block;" id="canvas-chart"></canvas>
     </div>
 </div>
</body>

{{ end }}
`

