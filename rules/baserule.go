package rules

type BaseRule struct {
	next Rule
}

func (b *BaseRule) SetNext(r Rule) {
	b.next = r
}

func (b *BaseRule) callNext(args *CallArgs) {
	if b.next != nil {
		b.next.Check(args)
	}
}
