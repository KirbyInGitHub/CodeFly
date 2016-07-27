package template

// EnumTemplate Swift Enum 模板
const EnumTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public enum {{ .Name }}: Int {
    {{ range $i, $f := .Cases }}
    case {{ $f.Name }} = {{ $f.Value }} {{ end }}
    
    public init?(code: Int?) {
        
        guard let code = code else { return nil }
        
        self.init(rawValue: code)
    }
}`