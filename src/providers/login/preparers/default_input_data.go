package loginpreparers

import (
	logininterfaces "github.com/golang-etl/base-playwright/src/providers/login/interfaces"
	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
)

func DefaultInputData(originalInputData logininterfaces.InputData) logininterfaces.InputData {
	defaults := logininterfaces.InputData{}

	return packagegeneralutils.MergeDefaults(originalInputData, defaults)
}
