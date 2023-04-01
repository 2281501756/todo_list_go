package cache

import (
	"fmt"
	"strconv"
)

const TaskRankKey = "TaskRank"

func GetTaskViewKey(id uint) string {
	return fmt.Sprintf("task:view:%s", strconv.Itoa(int(id)))
}
