package internal

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	path        = "../test"
	filePath    = "../test/NASTY_VAR"
	expEnvSlice = []string{"NASTY_VAR=﴾͡๏̯͡๏﴿ O'RLY?", "SOME_VAR=¯\\_(ツ)_/¯"}
	someProg    = "testProg"
)

func TestNewProgRunner(t *testing.T) {
	pr := NewProgRunnerImpl()
	assert.NotNil(t, pr)
}

func TestProgRunner_ExecutePath(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockProgRunner(ctrl)
	SetProgRunner(m)

	m.EXPECT().Execute().Return(nil)

	pr := GetProgRunner()
	err := pr.Execute()
	assert.NoError(t, err)
}

func TestProgRunner_ExecuteDirPath(t *testing.T) {
	pr := ProgRunnerImpl{
		ClearEnv:  true,
		EnvPath:   path,
		ChildProg: someProg,
		envVars:   make([]string, 0),
	}

	_ = pr.Execute()
	assert.Len(t, pr.envVars, 2)
	assert.Equal(t, expEnvSlice, pr.envVars)
}

func TestProgRunner_ExecuteFilePath(t *testing.T) {
	pr := ProgRunnerImpl{
		ClearEnv:  true,
		EnvPath:   filePath,
		ChildProg: someProg,
		envVars:   make([]string, 0),
	}

	_ = pr.Execute()
	assert.Len(t, pr.envVars, 1)
	assert.Equal(t, []string{expEnvSlice[0]}, pr.envVars)
}
