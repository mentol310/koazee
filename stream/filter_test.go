package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/internal/filter"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Filter(t *testing.T) {

	s := stream.New([]int{3, 5, 6, 1, -1}).Filter(func(element int) bool {
		return element >= 3
	})
	array := s.Out().Val()
	assert.Equal(t, []int{3, 5, 6}, array)
}

func TestStream_Filter_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The filter operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Filter(10).Out().Err())
	/**
		assert.Equal(
			t,
			errors.EmptyStream(filter.OpCode, "A nil Stream can not be filtered"),
			koazee.Stream().Filter(func() {}).Out().Err())
	**/
	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The provided function must retrieve 1 or 2 argument"),
		koazee.StreamOf([]int{2, 3, 2}).Filter(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).Filter(func(val string) bool { return true }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The provided function must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).Filter(func(val int) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The type of the Output in the provided function must be bool"),
		koazee.StreamOf([]int{2, 3, 2}).Filter(func(val int) string { return "a" }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(filter.OpCode, "The type of the argument 2 in the provided function must be string"),
		stream.New(map[string]int{"a": 2, "b": 3}).Filter(func(val int, key bool) bool { return true }).Out().Err())
}

func TestStream_Filter_WithIndex(t *testing.T) {
	expect := map[string]int{
		"a": 1,
		"c": 3,
	}
	actual := stream.New(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}).
		Filter(func(val int, key string) bool {
			return val == 1 || key == "c"
		}).Out().Val()
	assert.Equal(t, expect, actual)
}
