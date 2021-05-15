package spacelift

import (
	"fmt"
	s "github.com/cloudposse/terraform-provider-utils/internal/stack"
)

// CreateSpaceliftStacks takes a list of paths to YAML config files, processes and deep-merges all imports,
// and returns a map of Spacelift stack configs
func CreateSpaceliftStacks(filePaths []string, processStackDeps bool, processComponentDeps bool) (map[string]interface{}, error) {
	var _, mapResult, err = s.ProcessYAMLConfigFiles(filePaths, processStackDeps, processComponentDeps)
	if err != nil {
		return nil, err
	}
	return TransformStackConfigToSpaceliftStacks(mapResult)
}

// TransformStackConfigToSpaceliftStacks takes a a map of stack configs and transforms it to a map of Spacelift stacks
func TransformStackConfigToSpaceliftStacks(stacks map[string]interface{}) (map[string]interface{}, error) {
	res := map[string]interface{}{}

	for stackName, stackConfig := range stacks {
		config := stackConfig.(map[interface{}]interface{})
		imports := []string{}

		if i, ok := config["imports"]; ok {
			imports = i.([]string)
		}

		if i, ok := config["components"]; ok {
			componentsSection := i.(map[string]interface{})

			if terraformComponents, ok := componentsSection["terraform"]; ok {
				terraformComponentsMap := terraformComponents.(map[string]interface{})

				for component, v := range terraformComponentsMap {
					spaceliftStackName := fmt.Sprintf("%s-%s", stackName, component)
					componentMap := v.(map[string]interface{})

					componentVars := map[interface{}]interface{}{}
					if i, ok2 := componentMap["vars"]; ok2 {
						componentVars = i.(map[interface{}]interface{})
					}

					componentSettings := map[interface{}]interface{}{}
					if i, ok2 := componentMap["settings"]; ok2 {
						componentSettings = i.(map[interface{}]interface{})
					}

					componentEnv := map[interface{}]interface{}{}
					if i, ok2 := componentMap["env"]; ok2 {
						componentEnv = i.(map[interface{}]interface{})
					}

					spaceliftConfig := map[string]interface{}{}
					spaceliftConfig["imports"] = imports
					spaceliftConfig["vars"] = componentVars
					spaceliftConfig["settings"] = componentSettings
					spaceliftConfig["env"] = componentEnv
					res[spaceliftStackName] = spaceliftConfig
				}
			}

		}
	}

	return res, nil
}
