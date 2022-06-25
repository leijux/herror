package herror

func ResultErr[T any](result T, err error) ErrResult[T] {
	return ErrResult[T]{
		err:    err,
		result: result,
	}
}

//Try ResultErr的别名
func Try[T any](result T, err error) ErrResult[T] {
	return ResultErr(result, err)
}

func ResultsErr[T1, T2 any](result1 T1, result2 T2, err error) ErrResults[T1, T2] {
	return ErrResults[T1, T2]{
		err:     err,
		result1: result1,
		result2: result2,
	}
}

func HandleErr(err error) Err {
	return Err{
		err: err,
	}
}

func DeferErr(f func() error) DeferFuncErr {
	return DeferFuncErr{
		f,
	}
}
