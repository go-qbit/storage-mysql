package mysql

import (
	"bytes"
	"strings"
)

type SqlBuffer struct {
	*bytes.Buffer
	args []interface{}
}

func NewSqlBuffer() *SqlBuffer {
	return &SqlBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

func (b *SqlBuffer) Reset() {
	b.Buffer.Reset()
	b.args = b.args[:0]
}

func (b *SqlBuffer) GetSQL() string {
	return b.Buffer.String()
}

func (b *SqlBuffer) GetArgs() []interface{} {
	return b.args
}

func (b *SqlBuffer) String() string {
	buf := &bytes.Buffer{}
	argPos := 0

	query := b.GetSQL()
	args := b.GetArgs()

	for i := 0; i < len(query); i++ {
		q := strings.IndexByte(query[i:], '?')
		if q == -1 {
			buf.WriteString(query[i:])
			break
		}
		buf.WriteString(query[i : i+q])
		i += q

		buf.WriteString(Quote(args[argPos]))
		argPos++
	}

	return buf.String()
}

func (b *SqlBuffer) WriteIdentifier(identifier string) {
	b.WriteString(QuoteIdentifier(identifier))
}

func (b *SqlBuffer) WriteValue(value interface{}) {
	b.WriteByte('?')
	b.args = append(b.args, value)
}

func (b *SqlBuffer) WriteIdentifiersList(identifiers []string) {
	for i, identifier := range identifiers {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteIdentifier(identifier)
	}
}

func (b *SqlBuffer) WriteValuesList(values []interface{}) {
	for i, value := range values {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteValue(value)
	}
}
