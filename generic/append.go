package generic

type appendOption struct {
	uniq bool
}

type AppendOption func(oo *appendOption)

func WithUnique(t bool) AppendOption {
	return func(oo *appendOption) {
		oo.uniq = t
	}
}

// Appends extracts data from a slice of items using the provided getData function.
// The `unique` parameter is optional. If provided, it determines whether to include only unique IDs.
// If `unique` is true, only unique IDs are appended to the result. If false or omitted, all IDs are included.
func Appends[K comparable, T any](items []T, getData func(T) K, options ...AppendOption) []K {
	appends := make([]K, 0)
	opts := appendOption{}

	for _, option := range options {
		option(&opts)
	}

	if opts.uniq {
		seen := make(map[K]struct{})
		for _, item := range items {
			data := getData(item)
			if _, exists := seen[data]; !exists {
				seen[data] = struct{}{}
				appends = append(appends, data)
			}
		}
		clear(seen)
	} else {
		for _, item := range items {
			appends = append(appends, getData(item))
		}
	}
	return appends
}

func ConvertToMap[K comparable, T any](items []T, getID func(T) K) map[K]T {
	result := make(map[K]T, len(items))
	for _, item := range items {
		result[getID(item)] = item
	}
	return result
}
