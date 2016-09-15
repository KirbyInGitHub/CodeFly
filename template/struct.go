package template

// StructTemplate Swift Struct 模板
const StructTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public class {{ .Name }}: JSONItem {
    {{ range $i, $f := .Fields }}
    var {{ $f.Name }}{{ $ss.ValueTypeFormat $f }}{{ end }}
    
    public var allKeys: Set<String> {
        return [{{ range $i, $f := .Fields }}{{ if ne $i 0 }}, {{ end }}"{{ $f.Name }}"{{ end }}]
    } 

    public required init?(json: Any?) {

        super.init()

        guard let json = json as? [String: Any] else { return nil }
        {{ range $i, $f := .Fields }}
        self.{{ $f.Name }} = {{ $ss.FromDict $f }}{{ end }}
    }

    public func toJSON() -> Any? {

        var json = [String: Any]()
        {{ range $i, $f := .Fields }}
        json["{{ $f.Name }}"] = {{ $ss.ToDict $f }}{{ end }}

        return json
    }
}`
