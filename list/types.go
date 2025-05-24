package list

// List 接口
// 该接口只定义清楚各个方法的行为和表现
type List[T any] interface {
	// Get 返回对应的下标元素
	Get(index int) (T, error)
	// Append 在末尾追加元素
	Append(src ...T) error
	// Add 在指定索引位置处插入元素
	Add(val T, index int) error
	// Set 将指定索引位置处的值设置为 val
	Set(val T, index int) error
	// Delete 删除指定索引位置处的值
	Delete(index int) (T, error)
	// Cap 计算容量
	Cap() int
	// Len 计算长度
	Len() int
	// AsSlice 将 List 转化为一个切片
	AsSlice() []T
}
