package basic

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestMakeUI(t *testing.T) {
	out, in := MakeUI()
	if out.Text != "Hello Fyne!" {
		t.Errorf("Incorrect text")
	}
	// mock input, 输入后out的内容变成为Hello Hank!
	test.Type(in, "Hank")
	if out.Text != "Hello Hank!" {
		t.Errorf("Incorrect text")
	}
}
