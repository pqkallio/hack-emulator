package components

type MockOut struct {
	Result Val
}

func (m *MockOut) Update(opts ...UpdateOpts) Val {
	if len(opts) != 1 {
		panic("expected opts to have length of 1")
	}

	m.Result = opts[0].val

	return m.Result
}
