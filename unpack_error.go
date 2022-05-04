package handy

func UnpackError[T any](_ T, err error) error {
	return err
}
