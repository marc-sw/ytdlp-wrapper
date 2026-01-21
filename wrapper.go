package ytdlp

import (
	"os/exec"
	"strings"
)

type Command map[Flag]string

func (command Command) SetFlag(flag Flag) {
	command.SetParam(flag, "")
}

func (command Command) SetParam(flag Flag, value string) {
	command[flag] = value
}

func (command Command) Remove(flag Flag) {
	delete(command, flag)
}

func (command Command) Contains(flag Flag) bool {
	_, exists := command[flag]
	return exists
}

func (command Command) Build() []string {
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
