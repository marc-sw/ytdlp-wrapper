package ytdlp

import (
	"testing"
)

func TestVersion(t *testing.T) {
	got, err := Version()

	want := "2025.12.08"
	if err != nil {
		t.Errorf("Error: %v", err)
	} else if got != want {
		t.Errorf("Version() = %v, want %v", got, want)
	}
}

func TestCommandBuilding(t *testing.T) {
	var arguments Arguments = Arguments{}

	arguments.SetFlag(FlagVerbose)
	arguments.SetParam(ParamAudio(AudioAAC))
	arguments.SetFlag(FlagVerbose)
	arguments.Remove(FlagVerbose)
	arguments.SetFlag(FlagConsoleTitle)
	arguments.SetParam(ParamAudio(AudioAAC))

	var args []string = arguments.Build()

	if len(args) != 3 {
		t.Errorf("wrong args length, expectet: %v, got: %v", 3, len(args))
	}
}

/*
func TestOpenBrowser(t *testing.T) {
	err := OpenHomepage()

	if err != nil {
		t.Errorf("open browser error: %v",)
	}
}*/
