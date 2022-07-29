package split

import (
	"fmt"
	"reflect"
	"testing"
)

// go test 测设所有用例
// -v 显示详情
// -run <正则> 测设所有满足正则的用例
// go test -cover 显示覆盖率
// go test -coverprofile=c.out 生成覆盖率文件
// go tool cover -html c.out 在 html 显示 c.out 覆盖率

// 单元测试
func TestSplit(t *testing.T) {
	type ts struct {
		s   string
		seq string
		exp []string
	}
	tsm := map[string]ts{
		"test1": ts{"a,b,c", ",", []string{"a", "b", "c"}},
		"test2": ts{"a,b,c", "", []string{"a,b,c"}},
		"test3": ts{"a,b,a", "a", []string{",b,"}},
		"test4": ts{"a,b,a", ",b,", []string{"a", "a"}},
	}
	for name, ti := range tsm {
		t.Run(name, func(t *testing.T) {
			res := Split(ti.s, ti.seq)
			if !reflect.DeepEqual(ti.exp, res) {
				t.Errorf("期望得到%v,实际得到%v", ti.exp, res)
			}
		})
	}
}

// go test -bench <正则> 基准测试
// -cpu <n> 使用 n 个 cpu 测试
// -benchmem 显示内存使用及内存分配情况

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a,b,c", ",")
	}
}

// 示例函数
func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// OutPut: 3
}
