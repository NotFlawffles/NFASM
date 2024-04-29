package main

type Span struct {
    stream string
    index, row, column, length uint64
}

func NewSpan(stream string, index, row, column, length uint64) Span {
    return Span{stream, index, row, column, length}
}

func (this *Span) WithLength(length uint64) *Span {
    this.length = length
    return this
}
