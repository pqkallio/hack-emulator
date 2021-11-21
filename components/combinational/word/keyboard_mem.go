package word

type KeyboardMem struct {
	val uint16
}

func NewKeyboardMem() *KeyboardMem {
	return &KeyboardMem{}
}

func (k *KeyboardMem) Get() uint16 {
	return k.val
}

func (k *KeyboardMem) Update(val uint16) {
	k.val = val
}
