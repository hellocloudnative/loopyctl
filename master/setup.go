package master

import (
	"loopyctl/pkg/logger"
)

func RunSetup() error {
	if err := runCommand("cp", "kylin-loopy/sealos", "/usr/bin/"); err != nil {
		return err
	}
	logger.Info("Installtools installed successfully")

	if err := runCommand("sealos", "run", "registry.cn-shanghai.aliyuncs.com/labring/kubernetes:v1.27.7", "registry.cn-shanghai.aliyuncs.com/labring/helm:v3.9.4", "labring/kube-ovn:v1.11.5", "--single"); err != nil {
		return err
	}
	logger.Info("loopy kubernetes Setup completed successfully")

	// 执行初始化脚本
	if err := runCommand("sh", "kylin-loopy/init.sh"); err != nil {
		return err
	}
	logger.Info("Init script executed successfully")
	deployFiles := []string{
		"kylin-loopy/deploy/metrics-server.yaml",
		"kylin-loopy/deploy/local-path-provisioner",
	}
	for _, file := range deployFiles {
		if err := kubectlApply(file); err != nil {
			return err
		}
	}
	logger.Info("Deploy files applied successfully")
	kubevirtFiles := []string{
		"kylin-loopy/deploy/kubevirt/kubevirt-operator.yaml",
		"kylin-loopy/deploy/kubevirt/kubevirt-cr.yaml",
	}
	for _, file := range kubevirtFiles {
		if err := kubectlCreate(file); err != nil {
			return err
		}
	}
	logger.Info("Kubevirt deploy files applied successfully")

	cdiFiles := []string{
		"kylin-loopy/deploy/cdi/cdi-operator.yaml",
		"kylin-loopy/deploy/cdi/cdi-cr.yaml",
		"kylin-loopy/deploy/cdi/cdi-uploadproxy-service.yaml",
	}
	for _, file := range cdiFiles {
		if err := kubectlCreate(file); err != nil {
			return err
		}
	}
	logger.Info("CDI deploy files applied successfully")
	if err := runCommand("sh", "-c", "cd", "kylin-loopy", "&&", "./update.sh"); err != nil {
		return err
	}
	logger.Info("loopy Update script executed successfully")
	if err := kubectlApply("kylin-loopy/deploy/virt-vnc.yaml"); err != nil {
		return err
	}
	logger.Info("Virt-vnc deploy file applied successfully")
	if err := runCommand("cp", "kylin-loopy/virtctl-v1.2.0-linux-amd64", "/usr/bin/virtctl"); err != nil {
		return err
	}
	logger.Info("Virtctl copied to /usr/bin/")
	if err := runCommand("chmod", "+x", "/usr/bin/virtctl"); err != nil {
		return err
	}
	logger.Info("Virtctl copied to /usr/bin/ and made executable")

	logger.Info("Loopy Setup completed.")
	return nil
}
