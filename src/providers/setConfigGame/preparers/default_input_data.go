package setConfigGamepreparers

import (
	setConfigGameinterfaces "github.com/autoika/api-config/src/providers/setConfigGame/interfaces"
	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
)

func DefaultInputData(originalInputData setConfigGameinterfaces.InputData) setConfigGameinterfaces.InputData {
	defaults := setConfigGameinterfaces.InputData{}

	return packagegeneralutils.MergeDefaults(originalInputData, defaults)
}
