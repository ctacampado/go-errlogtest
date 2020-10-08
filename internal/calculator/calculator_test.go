package calculator

import (
	"errors"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tc := struct {
		in  []float64
		exp float64
	}{
		in:  []float64{1, 2, 3},
		exp: 6,
	}

	if want, got := tc.exp, Add(tc.in...); want != got {
		t.Errorf("expected return value is %f, got %f | fail", want, got)
		return
	}
}

func TestSub(t *testing.T) {
	tc := struct {
		in  []float64
		exp float64
	}{
		in:  []float64{1, 2, 3},
		exp: -6,
	}

	if want, got := tc.exp, Sub(tc.in...); want != got {
		t.Errorf("expected return value is %f, got %f | fail", want, got)
		return
	}
}

func TestMultiply(t *testing.T) {
	tc := struct {
		in  []float64
		exp float64
	}{
		in:  []float64{1, 2, 3},
		exp: 6,
	}

	if want, got := tc.exp, Multiply(tc.in...); want != got {
		t.Errorf("expected return value is %f, got %f | fail", want, got)
		return
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		in  []float64
		exp DivAns
	}{
		{
			in:  []float64{10, 0, 0},
			exp: DivAns{Q: 0, E: errors.New("cannot divide by zero")},
		},
		{
			in:  []float64{1, 1, 0},
			exp: DivAns{Q: 0, E: errors.New("cannot divide by zero")},
		},
		{
			in:  []float64{0, 0, 1},
			exp: DivAns{Q: 0, E: errors.New("cannot divide by zero")},
		},
		{
			in:  []float64{4, 2, 1},
			exp: DivAns{Q: 2, E: nil},
		},
		{
			in:  []float64{1, 2, 3},
			exp: DivAns{Q: 0.16666666666666666, E: nil},
		},
	}

	for i, test := range testCases {
		if want, got := test.exp, Divide(test.in...); fmt.Sprintf("%+v", want) != fmt.Sprintf("%+v", got) {
			t.Errorf("expected return value is %+v, got %+v | fail", want, got)
		} else {
			t.Logf("ts %d: | args: %+v | pass", i, want)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Errorf("no test")
}

func TestSqrd(t *testing.T) {
	t.Errorf("no test")
}

func TestAreaOfSquare(t *testing.T) {
	// 1 Given a square
	type Square struct {
		a float64
		b float64
		c float64
		d float64
	}
	square := Square{2, 2, 2, 2}

	// 2 When I calculate the Area of the square
	test := struct {
		in  Square
		exp float64
	}{
		in:  square,
		exp: 4.0,
	}
	calcArea := func(s Square) float64 {
		return Sqrd(s.a)
	}
	got := calcArea(test.in)

	// 3 Then I should get the Area of the square
	if want := test.exp; want != got {
		t.Errorf("expected return value is %+v, got %+v | fail", want, got)
	} else {
		t.Logf("args: %+v,  got %+v | pass", want, got)
	}
}
