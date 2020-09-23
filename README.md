# go-google-charts
A Golang wrapper that use google charts in the backend as templates. Currently, only OrgCharts are utilized. Very useful for pdf generation. The idea is to build a JSON object ( can be from an API call ) and pass it to the template with 3 pre-defined columns: value, parent_value &amp; row_type.

Run the project:
```Golang 
go run main.go 
```

The following example will be displayed in localhost:8080

![Company Hierarchy Example](https://raw.githubusercontent.com/lelinu/go-google-charts/master/src/example/org-chart-example.jpg)

