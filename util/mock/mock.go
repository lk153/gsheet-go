package mock

import "github.com/stretchr/testify/mock"

func GenMocksParams(times int) (params []interface{}) {
	for i := 0; i < times; i++ {
		params = append(params, mock.Anything)
	}

	return
}
