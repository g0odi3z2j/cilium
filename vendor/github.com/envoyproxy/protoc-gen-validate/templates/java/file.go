package java

const fileTpl = `// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: {{ .File.InputPath }}

package {{ javaPackage .File }};
{{ if isOfFileType . }}
public class {{ classNameFile . }}Validator {
	public static io.envoyproxy.pgv.ValidatorImpl validatorFor(Class clazz) {
		{{ range .AllMessages }}
		if (clazz.equals({{ qualifiedName . }}.class)) return new {{ simpleName .}}Validator();
		{{- end }}
		return null;
	}

{{ range .AllMessages -}}
	{{- template "msg" . -}}
{{- end }}
}
{{ else }}
/**
* Validates {@code {{ simpleName . }}} protobuf objects.
*/
public class {{ classNameMessage .}}Validator implements io.envoyproxy.pgv.ValidatorImpl<{{ qualifiedName . }}>{
	public static io.envoyproxy.pgv.ValidatorImpl validatorFor(Class clazz) {
		if (clazz.equals({{ qualifiedName . }}.class)) return new {{ simpleName .}}Validator();
		{{ range .AllMessages }}
		if (clazz.equals({{ qualifiedName . }}.class)) return new {{ simpleName .}}Validator();
		{{- end }}
		return null;
	}
	{{- template "msgInner" . -}}
	{{ range .AllMessages -}}
	{{- template "msg" . -}}
	{{- end }}
}
{{ end }}
`
