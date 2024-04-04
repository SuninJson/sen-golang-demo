package demo_m01_rate_limter

import (
	"container/heap"
	"time"
)

// CheckBySlidingWindow
// 确定一个起始时间点，一般就是系统启动的时间，并记录该点为时间滑动窗口的起始点。然后创建一个空的列表作为时间滑动窗口内请求进入的时间戳记录
// 当请求到来时，使用当前时间戳比较它是否在时间窗口的范围内
// 若在，则查看当前时间窗口内记录的所有请求的数量：
// （1）若超过，则拒绝请求
// （2）若没有超过，则将请求的时间戳加入到滑动窗口内的时间戳记录中
// 若不在，则查看时间戳记录，将时间戳最久远的记录删除，然后将时间窗口的起始点更新为第二久远的记录时间，然后再次检查，至到能检查到请求的时间戳
//滑动窗口是一种否决形式的限流，对于超出阈值的流量只能拒绝，很难做到削峰填谷，并且它需要额外的存储空间和使用了优先级队列，所有性能较低，单机百万请求下延迟很高
func CheckBySlidingWindow(startTime int, SlidingWindowTimeBucket int, threshold int) bool {
	var now int = int(time.Now().Unix())
	if (startTime+SlidingWindowTimeBucket) > now && now > startTime {
		if heap.Size() > threshold {
			return false
		}
		heap.Push(now)
		return true
	}

	t := heap.Pop()
	startTime = heap.Top()
	return CheckBySlidingWindow()
}
