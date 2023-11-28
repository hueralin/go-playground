package search

// defaultMatcher 实现了默认匹配器
// 空结构在创建实例时不会分配任何内存
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
func (m defaultMatcher) Search(feed *Feed, searchItem string) ([]*Result, error) {
	return nil, nil
}
