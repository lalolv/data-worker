package workers

import (
	"fmt"
	"strconv"

	"github.com/lalolv/data-worker/metadata"
	"github.com/lalolv/data-worker/utils"
)

// IDNo 身份证号生成
func IDNo() string {
	// AreaCode 随机一个+4位随机数字(不够左填充0)
	areaCode := metadata.AreaCode[utils.RandIntRange(0, len(metadata.AreaCode))] +
		fmt.Sprintf("%0*d", 4, utils.RandIntRange(1, 9999))

	birthday := Date().Format("20060102")
	randomCode := fmt.Sprintf("%0*d", 3, utils.RandIntRange(0, 999))
	prefix := areaCode + birthday + randomCode

	return prefix + verifyCode(prefix)
}

// 获取 VerifyCode
func verifyCode(cardID string) string {
	tmp := 0
	for i, v := range metadata.Wi {
		t, _ := strconv.Atoi(string(cardID[i]))
		tmp += t * v
	}

	return metadata.ValCodeArr[tmp%11]
}
