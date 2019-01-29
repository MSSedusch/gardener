{{- define "azure-azuredns.variables" -}}
variable "CLIENT_ID" {
  description = "Azure client id of technical user"
  type        = "string"
}

variable "CLIENT_SECRET" {
  description = "Azure client secret of technical user"
  type        = "string"
}

variable "TENANT_ID" {
  description = "Azure tenant of technical user"
  type        = "string"
}
{{- end -}}
