package keys

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Test_updateKeyCommand(t *testing.T) {
	cmd := UpdateKeyCommand(sdk.NewDefaultConfig())
	assert.NotNil(t, cmd)
	// No flags  or defaults to validate
}

func Test_runUpdateCmd(t *testing.T) {
	fakeKeyName1 := "runUpdateCmd_Key1"
	fakeKeyName2 := "runUpdateCmd_Key2"

	config := sdk.NewDefaultConfig()
	cmd := UpdateKeyCommand(config)

	// fails because it requests a password
	assert.EqualError(t, runUpdateCmd(config)(cmd, []string{fakeKeyName1}), "EOF")

	// try again
	mockIn, _, _ := tests.ApplyMockIO(cmd)
	mockIn.Reset("pass1234\n")
	assert.EqualError(t, runUpdateCmd(config)(cmd, []string{fakeKeyName1}), "Key runUpdateCmd_Key1 not found")

	// Prepare a key base
	// Now add a temporary keybase
	kbHome, cleanUp1 := tests.NewTestCaseDir(t)
	defer cleanUp1()
	viper.Set(flags.FlagHome, kbHome)

	kb, err := NewKeyBaseFromHomeFlag(config)
	assert.NoError(t, err)
	_, err = kb.CreateAccount(fakeKeyName1, tests.TestMnemonic, "", "", 0, 0)
	assert.NoError(t, err)
	_, err = kb.CreateAccount(fakeKeyName2, tests.TestMnemonic, "", "", 0, 1)
	assert.NoError(t, err)

	// Try again now that we have keys
	// Incorrect key type
	mockIn.Reset("pass1234\nNew1234\nNew1234")
	err = runUpdateCmd(config)(cmd, []string{fakeKeyName1})
	assert.EqualError(t, err, "locally stored key required. Received: keys.offlineInfo")

	// TODO: Check for other type types?
}
