package collection

import "reflect"

func GetTags(src any, tagName string, ignoreTags ...string) []string {
	var tags []string
	ty := reflect.TypeOf(src)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	if ty.Kind() != reflect.Struct {
		return tags
	}

	ignoreMap := make(map[string]struct{}, len(ignoreTags))
	for _, tag := range ignoreTags {
		ignoreMap[tag] = struct{}{}
	}

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "" {
			if _, ignored := ignoreMap[tag]; !ignored {
				tags = append(tags, tag)
			}
		}
	}

	return tags
}
