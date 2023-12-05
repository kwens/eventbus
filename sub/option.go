/**
 * @Author: kwens
 *
 * @Date: 2023-06-06 08:55:08
 * @Description:
 */
package sub

type SubOption func(*subOption)

type subOption struct {
	subName string
}

func defaultSubOpt() *subOption {
	return &subOption{}
}

func WithSubName(name string) SubOption {
	return func(so *subOption) {
		so.subName = name
	}
}
