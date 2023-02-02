package camel_test

import (
	"testing"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"

	"github.com/dtcookie/assert"
)

func TestCamel(t *testing.T) {
	assert := assert.New(t)

	assert.Equals("Asdf", camel.Camel("asdf"))
	assert.Equals("ASdf", camel.Camel("a_sdf"))
	assert.Equals("ASdf", camel.Camel("a_Sdf"))
	assert.Equals("ASdf", camel.Camel("A_sdf"))
	assert.Equals("ASDf", camel.Camel("A_s%df"))
	assert.Equals("DtEntityContainerGroup", camel.Camel("dt.entity.container_group"))
	assert.Equals("MarkedForTermination", camel.Camel("MARKED_FOR_TERMINATION"))

}
