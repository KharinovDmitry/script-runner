package executor

import (
	"TestTask-PGPro/lib/adapter/executor/utils"
	"context"
	"os"
	"os/exec"
)

type LinuxAdapter struct {
}

func NewLinuxAdapter() LinuxAdapter {
	return LinuxAdapter{}
}

func (l LinuxAdapter) Run(ctx context.Context, text string, outputChan chan<- []byte) error {
	defer close(outputChan)

	file, err := utils.CreateTempFileWithText(text)
	if err != nil {
		return err
	}
	file.Close()
	defer os.Remove(file.Name())

	if err := utils.AddFileExecutablePermission(file.Name()); err != nil {

		return err
	}

	cmd := exec.CommandContext(ctx, "sh", file.Name())

	output, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	for {
		outBytes := make([]byte, 1024)
		n, err := output.Read(outBytes)
		if err != nil {
			break
		}
		outputChan <- outBytes[:n]
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
