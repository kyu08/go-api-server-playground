package repository

import (
	"reflect"
	"runtime"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

func toStruct[T any](iter *spanner.RowIterator) ([]*T, error) {
	res := make([]*T, 0, iter.RowCount)

	callerName := func() string {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			return "unknown"
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return "unknown"
		}

		return fn.Name()
	}()
	entityName := reflect.TypeFor[T]().Name()

	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError(callerName, entityName, err)
		}

		var t T
		if err := row.ToStruct(&t); err != nil {
			return nil, newErrorWithCode(codes.Internal, callerName, entityName, err)
		}

		res = append(res, &t)
	}

	return res, nil
}
