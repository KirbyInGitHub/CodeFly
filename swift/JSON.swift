//
// JSON.swift
//
// 模板生成的结构必须继承 JSON 类
//

public protocol JSONParser: NSObjectProtocol {

    init()

    func from(json: Any) -> Bool

    var json: Any { get }
}

extension JSONParser {

    public init?(json: Any?) {

        guard let json = json else { return nil }

        self.init()

        guard self.from(json: json) else { return nil }
    }
}

extension Array where Element: JSONParser {

    public init?(json: Any?) {

        guard let elements = json as? [Any] else { return nil }

        self = elements.flatMap { Element(json: $0) }
    }

    public mutating func from(json: Any) -> Bool {

        if let elements = Array<Element>(json: json) {

            self = elements

            return true

        } else {

            return false
        }
    }

    public var json: Any { return self.flatMap { $0.json } }
}

public class JSON: NSObject, JSONParser {

    override public required init() { super.init() }

    @discardableResult
    public func from(json: Any) -> Bool { return false }

    public var json: Any { fatalError("Not implemented.") }
}

extension JSON: NSCopying {

    public func copy(with zone: NSZone? = nil) -> Any {

        let obj = type(of: self).init()

        obj.from(json: self.json)

        return obj
    }
}