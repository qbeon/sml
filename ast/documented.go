package ast

// BaseDocumented represents the basic documented expression info
type BaseDocumented struct {
	Doc str
}

// Documentation implements the DocumentedExpression interface
func (doc *BaseDocumented) Documentation() str {
	return doc.Doc
}
