package bit

// Or8Way is a 8-way OR gate, evaluating to true if any of the inputs are true.
type Or8Way struct {
	or1 *Or
	or2 *Or
	or3 *Or
	or4 *Or
	or5 *Or
	or6 *Or
	or7 *Or
}

func NewOr8Way() *Or8Way {
	return &Or8Way{
		NewOr(), NewOr(),
		NewOr(), NewOr(),
		NewOr(), NewOr(),
		NewOr(),
	}
}

func (or8Way *Or8Way) Update(a, b, c, d, e, f, g, h bool) bool {
	temp1 := or8Way.or1.Update(a, b, nil, 0)
	temp2 := or8Way.or2.Update(temp1, c, nil, 0)
	temp3 := or8Way.or3.Update(temp2, d, nil, 0)
	temp4 := or8Way.or4.Update(temp3, e, nil, 0)
	temp5 := or8Way.or5.Update(temp4, f, nil, 0)
	temp6 := or8Way.or6.Update(temp5, g, nil, 0)
	return or8Way.or7.Update(temp6, h, nil, 0)
}
