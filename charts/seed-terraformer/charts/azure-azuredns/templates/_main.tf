{{- define "azure-azuredns.main" -}}

locals {
  azuredns_subscriptionId="${replace("{{ required "record.hostedZoneID is required" .Values.record.hostedZoneID }}", "//subscriptions/([^/]*)/resourceGroups/([^/]*)/providers/([^/]*)/([^/]*)/(.*)/", "$1")}"
  azuredns_resource_group_name="${replace("{{ required "record.hostedZoneID is required" .Values.record.hostedZoneID }}", "//subscriptions/([^/]*)/resourceGroups/([^/]*)/providers/([^/]*)/([^/]*)/(.*)/", "$2")}"
  azuredns_zone_name="${replace("{{ required "record.hostedZoneID is required" .Values.record.hostedZoneID }}", "//subscriptions/([^/]*)/resourceGroups/([^/]*)/providers/([^/]*)/([^/]*)/(.*)/", "$5")}"
  recordsList=[
{{- include "azure-azuredns.records" $.Values | trimSuffix "," | indent 4 }}
  ]
}

provider "azurerm" {
  subscription_id = "${local.azuredns_subscriptionId}"
  tenant_id       = "${var.TENANT_ID}"
  client_id       = "${var.CLIENT_ID}"
  client_secret   = "${var.CLIENT_SECRET}"
}

//=====================================================================
//= AzureDNS Record
//=====================================================================

resource "azurerm_dns_a_record" "recorda" {
  count = {{ if eq (required "record.type is required" .Values.record.type) "ip" }}1{{ else }}0{{ end }}

  name     = "${replace("{{ required "record.name is required" .Values.record.name }}", format(".%s", local.azuredns_zone_name), "")}"
  zone_name           = "${local.azuredns_zone_name}"
  resource_group_name = "${local.azuredns_resource_group_name}"
  ttl      = "120"
  records  = "${local.recordsList}"
}

resource "azurerm_dns_cname_record" "recordcname" {
  count = {{ if eq (required "record.type is required" .Values.record.type) "ip" }}0{{ else }}1{{ end }}

  name     = "${replace("{{ required "record.name is required" .Values.record.name }}", format(".%s", local.azuredns_zone_name), "")}"
  zone_name           = "${local.azuredns_zone_name}"
  resource_group_name = "${local.azuredns_resource_group_name}"
  ttl                 = 120
  record              = "${element(local.recordsList, 0)}"
}

{{- end -}}
