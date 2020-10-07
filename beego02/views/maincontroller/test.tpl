{{.aa}}

<br/>
{{range $index, $elem := .ss}}
    {{$index}}=={{$elem}}  {{$.aa}}
    <br>
{{end}}