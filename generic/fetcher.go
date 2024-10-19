package generic

import (
	"context"
	"maps"
)

type FetcherMapFnOutput[T any, K comparable] struct {
	Items     map[K]T
	PageCount int64
}

type fetcherMapFn[T any, K comparable] func(ctx context.Context, page int64) (output FetcherMapFnOutput[T, K], err error)

func FetcherMap[T any, K comparable](ctx context.Context, fn fetcherMapFn[T, K]) (map[K]T, error) {
	allItems := make(map[K]T)
	page := int64(1)

	for {
		dataOutput, err := fn(ctx, page)
		if err != nil {
			return nil, err
		}

		maps.Copy(allItems, dataOutput.Items)

		if page >= dataOutput.PageCount {
			break
		}
		page++
	}

	return allItems, nil
}

type FetcherFnOutput[T any] struct {
	Items     []T
	PageCount int64
}
type fetcherFn[T any] func(ctx context.Context, page int64) (output FetcherFnOutput[T], err error)

func Fetcher[T any](ctx context.Context, fn fetcherFn[T]) ([]T, error) {
	allItems := make([]T, 0)
	page := int64(1)

	for {
		dataOutput, err := fn(ctx, page)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, dataOutput.Items...)

		if page >= dataOutput.PageCount {
			break
		}
		page++
	}

	return allItems, nil
}
