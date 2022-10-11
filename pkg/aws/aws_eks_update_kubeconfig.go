package aws

import (
	e "github.com/cloudposse/terraform-provider-utils/internal/exec"
	cfg "github.com/cloudposse/terraform-provider-utils/pkg/config"
	u "github.com/cloudposse/terraform-provider-utils/pkg/utils"
)

// ExecuteAwsEksUpdateKubeconfig executes 'aws eks update-kubeconfig'
// https://docs.aws.amazon.com/cli/latest/reference/eks/update-kubeconfig.html
func ExecuteAwsEksUpdateKubeconfig(kubeconfigContext cfg.AwsEksUpdateKubeconfigContext) error {
	err := e.ExecuteAwsEksUpdateKubeconfig(kubeconfigContext)

	if err != nil {
		u.PrintErrorToStdError(err)
		return err
	}

	return nil
}