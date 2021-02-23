/**
 * @Author: Huyantian
 * @Date: 2021/2/23 下午7:14
 */

package todo

type QueryParam struct {
	Status Status
}

type QueryFunc func(param *QueryParam)

func QueryStatus(status Status) QueryFunc {
	return func(param *QueryParam) {
		param.Status = status
	}
}

type UpdateParam struct {
	M map[string]interface{}
}

type UpdateFunc func(param *UpdateParam)

func UpdateStatus(status Status) UpdateFunc {
	return func(param *UpdateParam) {
		param.M["Status"] = status
	}
}
