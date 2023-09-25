package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(searchKey string) (string, error) {
	definition, ok := d[searchKey]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, val string) error {
	// when you pass a map to a function/method you are copying the pointer part not the underlying data structure
	// never declare a nil map like: var m map[string]string, it behaves like normal when reading but attempts to write
	// always declare as var dictionary = map[string]string{} or using make(map[string]string)
	// else it causes runtime panic
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, val string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = val
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
