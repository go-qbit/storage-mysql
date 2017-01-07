package mysql

import (
	"github.com/go-qbit/model/expr"
)

var exprProcessor = &ExprProcessor{}

type ExprProcessor struct{}

type WriteFunc func(*SqlBuffer)

func (p *ExprProcessor) Eq(op1, op2 expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op1.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteByte('=')
		op2.GetProcessor(p).(WriteFunc)(buf)
	})
}

func (p *ExprProcessor) Lt(op1, op2 expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op1.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteByte('<')
		op2.GetProcessor(p).(WriteFunc)(buf)
	})
}

func (p *ExprProcessor) Le(op1, op2 expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op1.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteString("<=")
		op2.GetProcessor(p).(WriteFunc)(buf)
	})
}

func (p *ExprProcessor) Gt(op1, op2 expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op1.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteByte('>')
		op2.GetProcessor(p).(WriteFunc)(buf)
	})
}

func (p *ExprProcessor) Ge(op1, op2 expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op1.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteString(">=")
		op2.GetProcessor(p).(WriteFunc)(buf)
	})
}

func (p *ExprProcessor) In(op expr.IExpression, values []expr.IExpression) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		op.GetProcessor(p).(WriteFunc)(buf)
		buf.WriteString(" IN (")
		for i, value := range values {
			if i > 0 {
				buf.WriteByte(',')
			}
			value.GetProcessor(p).(WriteFunc)(buf)
		}
		buf.WriteByte(')')
	})
}

func (p *ExprProcessor) ModelField(fieldName string) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		buf.WriteIdentifier(fieldName)
	})
}

func (p *ExprProcessor) Value(value interface{}) interface{} {
	return WriteFunc(func(buf *SqlBuffer) {
		buf.WriteValue(value)
	})
}
