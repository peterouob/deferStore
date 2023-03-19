package helper

func RemoveDuplicateInt64(slide []string) []string {
	res := make([]string, 0, len(slide))
	temp := map[string]struct{}{}
	for _, item := range slide {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}
