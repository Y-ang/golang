https://www.pseudoyu.com/zh/2021/05/29/algorithm_data_structure_go/


go语言定制指南 https://chai2010.cn/go-ast-book/ch8/index.html

go三个点...的用法 https://segmentfault.com/a/1190000020638199

判断元素是否在集合中 https://juejin.cn/post/7181299102490050616
map查询
先将slice转为 map，通过查询 map 来快速查看元素是否在 slice 中。

```
// ConvertStrSlice2Map 将字符串 slice 转为 map[string]struct{}。
func ConvertStrSlice2Map(sl []string) map[string]struct{} {
set := make(map[string]struct{}, len(sl))
for _, v := range sl {
set[v] = struct{}{}
}
return set
}

// InMap判断字符串是否在map对象中。
func InMap(m map[string]struct{}, s string) bool {
_, ok := m[s]
return ok
}
// 注意：使用空结构体 `struct{}` 作为 value 的类型，因为 `struct{}` 不占用任何内存空间。
// 虽然将 slice 转为 map 的时间复杂度为 O(n)，但是只转换一次可以忽略。
// 查询元素是否在 map 中的时间复杂度为 O(1)。
```

