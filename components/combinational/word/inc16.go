package word

type Inc16 struct {
	add16 *Add16
}

func NewInc16() *Inc16 {
	return &Inc16{NewAdd16()}
}

func (i *Inc16) Update(in uint16) uint16 {
	return i.add16.Update(in, 1)
}
