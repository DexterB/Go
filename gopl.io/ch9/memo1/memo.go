//+
// Package memo provides a concurrency unsafe memoization of a function of
// type Func.
package memo


// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface {}
	err    error
}

// A Memo caches the results of calling a function.
type Memo struct {
	f       Func
	cache   map[string] result
}


func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: This function is not concurrency safe.
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

//!=




