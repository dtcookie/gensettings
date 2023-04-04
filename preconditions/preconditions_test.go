package preconditions_test

import (
	"testing"

	pc "github.com/dtcookie/dynatrace/gensettings/preconditions"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
)

type TestStruct struct {
	Name    *string
	Enabled *bool
}

func TestEquals(t *testing.T) {
	testStruct := TestStruct{Name: nil, Enabled: opt.NewBool(true)}
	attName := pc.Attribute{Value: testStruct.Name, Name: "name"}
	attEnabled := pc.Attribute{Value: testStruct.Enabled, Name: "enabled"}

	if err := pc.Equals(attEnabled, true).Evaluate(attName); err != nil {
		if err.Error() != "'name' must be specified if 'enabled' is set to 'true'" {
			t.Error("unexpected error", err)
		}
	} else {
		t.Error("Error expected")
		return
	}

	testStruct = TestStruct{Name: opt.NewString("Foo"), Enabled: nil}
	attName = pc.Attribute{Value: testStruct.Name, Name: "name"}
	attEnabled = pc.Attribute{Value: testStruct.Enabled, Name: "enabled"}

	if err := pc.Equals(attEnabled, true).Evaluate(attName); err != nil {
		if err.Error() != "'name' requires 'enabled' to be set to 'true'" {
			t.Error("unexpected error", err)
		}
	} else {
		t.Error("Error expected")
		return
	}

	testStruct = TestStruct{Name: opt.NewString("Foo"), Enabled: opt.NewBool(false)}
	attName = pc.Attribute{Value: testStruct.Name, Name: "name"}
	attEnabled = pc.Attribute{Value: testStruct.Enabled, Name: "enabled"}

	if err := pc.Equals(attEnabled, true).Evaluate(attName); err != nil {
		if err.Error() != "'name' requires 'enabled' to be set to 'true'" {
			t.Error("unexpected error", err)
		}
	} else {
		t.Error("Error expected")
		return
	}
}
