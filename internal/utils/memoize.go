package utils

type MemoizedKeyConstraint interface {
	~string | int | float32 | float64 | bool
}
type MemoizedCB[Key MemoizedKeyConstraint, ReturnType any, Args any] func(key Key, args Args) ReturnType

func Memoize[Key MemoizedKeyConstraint, ReturnType any, Args any](
	fn MemoizedCB[Key, ReturnType, Args],
) MemoizedCB[Key, ReturnType, Args] {
	cache := make(map[Key]ReturnType)

	return func(key Key, args Args) ReturnType {
		if val, found := cache[key]; found {
			return val
		}

		result := fn(key, args)
		cache[key] = result

		return result
	}
}
