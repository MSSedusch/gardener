{{- define "azure-azuredns.records" -}}
{{- range $j, $record := .record.values }}
"{{ $record }}",
{{- end -}}
{{- end -}}
