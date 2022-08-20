package global

{{- if .HasGlobal }}

import "github.com/lzsgo/lzs-go-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}