package gotest

import "testing"

/*
压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
	func BenchmarkXXX(b *testing.B) { ... }
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
文件名也必须以_test.go结尾
*/

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Division(3, 4)
		if err != nil {
			return
		}
	}
}

/*func BenchMark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	// 做一些初始化的工作，例如读取文件数据，数据库连接之类的
	// 这样这些时间不影响我们测试函数本身的性能

	b.StopTimer() // 重新开始时间
	for i := 0; i < b.N; i++ {
		_, err := Division(3, 4)
		if err != nil {
			return
		}
	}
}*/

// 执行命令 go test webbench_test.go -test.bench=".*"