package database

import (
	"errors"
	"reflect"
	"runtime"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/model"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

func typeNameFromT[T any]() string {
	return reflect.TypeFor[T]().Name()
}

func ToStruct[T any](iter *spanner.RowIterator) ([]*T, error) {
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
	entityName := typeNameFromT[T]()

	for {
		row, err := iter.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return nil, model.NewError(callerName, entityName, err)
		}

		var t T
		if err := row.ToStruct(&t); err != nil {
			return nil, model.NewErrorWithCode(codes.Internal, callerName, entityName, err)
		}

		res = append(res, &t)
	}

	return res, nil
}

func NewNotFoundError[T any]() error {
	return apperrors.NewNotFoundError(typeNameFromT[T]())
}
