package search

import (
	"log"
	"sync"
)

// 包级变量
// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchItem string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchItem, results)
			// 递减 WaitGroup 的计数
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		// 该方法会导致 goroutine 阻塞，直到 WaitGroup 内部的计数到达 0
		waitGroup.Wait()

		// 以关闭通道的方式，通知 Display 函数可以退出程序了
		close(results)
	}()

	// 显示返回的结果
	// 并且在最后一个结果显示完后返回
	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
