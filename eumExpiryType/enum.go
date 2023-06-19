package eumExpiryType

// Enum 过期策略
type Enum int

const (
	// AbsoluteExpiration 绝对时间，到期自动移除
	AbsoluteExpiration Enum = iota
	// SlidingExpiration 任意一次访问都会重置过期时间
	SlidingExpiration
)
