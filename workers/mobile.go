package workers

import (
	"fmt"

	"github.com/lalolv/data-worker/metadata"
	"github.com/lalolv/data-worker/utils"
)

// Mobile 随机生成手机号
func Mobile() string {
	prefixIndex := utils.RandIntRange(0, len(metadata.MobilePrefix))
	otherData := fmt.Sprintf("%0*d", 8, utils.RandIntRange(0, 100000000))

	return metadata.MobilePrefix[prefixIndex] + otherData
}
