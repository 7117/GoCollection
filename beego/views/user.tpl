<div>
    {{if .isplay}}
     {{.email}}
    {{else}}
     {{.website}}
    {{end}}
</div>

<div>
    {{range $index, $elem := .a}}
        {{$index}} - {{$elem}} - {{$.email}}
    {{end}}
</div>