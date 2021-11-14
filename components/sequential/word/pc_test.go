package word

import "testing"

func TestNOP(t *testing.T) {
	t.Parallel()
	pc := NewPC()
	expected := uint16(0)
	actual := pc.Update(0x1234, false, false, false)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	pc.Tick()
	actual = pc.Update(0x1234, false, false, false)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}
}

func TestIncAndReset(t *testing.T) {
	t.Parallel()
	pc := NewPC()
	expected := uint16(0)
	actual := pc.Update(0x1234, false, true, false)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	pc.Tick()
	actual = pc.Update(0x1234, false, false, true)
	expected = uint16(1)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	pc.Tick()
	actual = pc.Update(0x1234, false, false, false)
	expected = uint16(0)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}
}

func TestLoadAndOverflow(t *testing.T) {
	t.Parallel()
	pc := NewPC()
	expected := uint16(0)
	actual := pc.Update(0xffff, true, false, false)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	pc.Tick()
	actual = pc.Update(0x1234, false, true, false)
	expected = uint16(0xffff)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	pc.Tick()
	actual = pc.Update(0x1234, false, true, false)
	expected = uint16(0)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}
}
