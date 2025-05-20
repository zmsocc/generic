package errs

import "fmt"

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("generic: 下标不在范围内，实际下标区间为 [0, %d], 操作的下标为 %d", length, index)
}

func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("generic: 类型转换失败，预期类型: %s, 实际值: %#v", want, got)
}
