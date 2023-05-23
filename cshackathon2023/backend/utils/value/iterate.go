package value

import "backend/types/response"

func Iterate[A any, B any](a []A, mapper func(a A) (B, *response.ErrorInstance)) ([]B, *response.ErrorInstance) {
	result := make([]B, 0)
	for _, el := range a {
		mapped, err := mapper(el)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped)
	}
	return result, nil
}
