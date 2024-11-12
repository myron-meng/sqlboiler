{{$specialCases := dict "os" "OS" "api" "API" "url" "URL"}}
var TableNames = struct {
    {{range $table := .Tables}}{{if not $table.IsView -}}
    {{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}} string
    {{end}}{{end -}}
}{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    {{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}: "{{$table.Name}}",
    {{end}}{{end -}}
}

var AllTableNames = []string{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    "{{$table.Name}}",
    {{end}}{{end -}}
}

var TableStructNames = struct {
    {{range $table := .Tables}}{{if not $table.IsView -}}
    {{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}} string
    {{end}}{{end -}}
}{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    {{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}: "{{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}",
    {{end}}{{end -}}
}

var AllTableStructNames = []string{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    "{{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}",
    {{end}}{{end -}}
}

var StructNameToSingularMap = map[string]string{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    "{{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}": "{{($.Aliases.Table $table.Name).UpSingular}}",
    {{end}}{{end -}}
}

var StructNameToPluralMap = map[string]string{
    {{range $table := .Tables}}{{if not $table.IsView -}}
    "{{if index $specialCases $table.Name}}{{index $specialCases $table.Name}}{{else}}{{titleCase $table.Name}}{{end}}": "{{($.Aliases.Table $table.Name).UpPlural}}",
    {{end}}{{end -}}
}
