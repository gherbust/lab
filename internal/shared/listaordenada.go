package shared

import "sort"

func OrdenarNombres(nombres *[]string) {
	sort.Sort(sort.StringSlice(*nombres))
}
