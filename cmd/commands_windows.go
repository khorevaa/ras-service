//go:build windows
// +build windows

package cmd

func init() {
	Commands = append(Commands,
		&installCmd{}, &removeCmd{},
	)
}
