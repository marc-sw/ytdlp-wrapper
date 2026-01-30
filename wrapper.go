package ytdlp

import (
	"os/exec"
	"strings"
)

type Arguments map[string]string

func NewArguments() Arguments {
	return make(map[string]string)
}

func (command Arguments) SetFlag(flag string) {
	command.SetParam(flag, "")
}

func (command Arguments) SetParam(flag string, value string) {
	command[flag] = value
}

func (command Arguments) Remove(flag string) {
	delete(command, flag)
}

func (command Arguments) Contains(flag string) bool {
	_, exists := command[flag]
	return exists
}

func (command Arguments) Build() []string {
	var args []string
	for k, v := range command {
		args = append(args, string(k))
		if len(v) > 0 {
			args = append(args, v)
		}
	}
	return args
}

func Version() (string, error) {
	cmd := exec.Command("yt-dlp", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func OpenHomepage() error {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#usage-and-options")
	err := cmd.Run()

	return err
}
