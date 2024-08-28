package util

func Merge[T any](items ...[]T) []T {
	for _, further := range items[1:] {
		items[0] = append(items[0], further...)
	}

	return items[0]
}
