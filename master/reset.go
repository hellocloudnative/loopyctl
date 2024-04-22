package master

import (
	"loopyctl/pkg/logger"
)

func Reset() error {
	if err := runCommand("sealos", "reset", "--force"); err != nil {
		return err
	}
	logger.Info("Cluster reset successfully")
	return nil
}
