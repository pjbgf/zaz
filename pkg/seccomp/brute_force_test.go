package seccomp

import (
	"testing"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pjbgf/go-test/should"
)

type runnerStub struct {
	profile     *specs.LinuxSeccomp
	callsToFail []string
}

func (r *runnerStub) Run(profile *specs.LinuxSeccomp) {
	r.profile = profile
}

// forces failures every time a system call on r.callsToFail is not
// in the profile being currently executed.
func (r *runnerStub) HasExecuted() bool {
	for _, a := range r.callsToFail {
		contains := false
		for _, n := range r.profile.Syscalls[0].Names {
			if a == n {
				contains = true
			}
		}

		if !contains {
			return false
		}
	}

	return true
}

func TestBruteForce_GetSystemCalls(t *testing.T) {
	assertThat := func(assumption string, injected []string,
		expected *specs.LinuxSyscall, expectedErr error) {
		should := should.New(t)
		stub := &runnerStub{callsToFail: injected}
		s := NewBruteForceSource(stub)

		actual, err := s.GetSystemCalls()

		should.BeEqual(expected, actual, assumption)
		should.BeEqual(expectedErr, err, assumption)
	}

	assertThat("should return all syscalls that the container can't run without",
		[]string{"read", "write", "close"},
		&specs.LinuxSyscall{
			Action: specs.ActAllow,
			Names:  []string{"read", "write", "close"},
		},
		nil)
}

func TestBruteForce_IndexesOf(t *testing.T) {
	assertThat := func(assumption string, source []string, item string, expected []int) {
		should := should.New(t)
		stub := &runnerStub{}
		s := NewBruteForceSource(stub)

		actual := s.indexesOf(source, item)

		should.BeEqual(expected, actual, assumption)
	}

	assertThat("should return empty []int when item not found",
		[]string{"item1", "item2", "item3"},
		"item4",
		[]int{})
	assertThat("should return single index when item found once",
		[]string{"item1", "item2", "item3"},
		"item2",
		[]int{1})
	assertThat("should return two indexes when item found twice",
		[]string{"item1", "item2", "item3", "item2"},
		"item2",
		[]int{1, 3})
}

func TestBruteForce_ExcludeItemFromSlice(t *testing.T) {
	assertThat := func(assumption string, source []string, itemToExclude string, expected []string) {
		should := should.New(t)
		stub := &runnerStub{}
		s := NewBruteForceSource(stub)

		actual := s.excludeItemFromSlice(source, itemToExclude)

		should.BeEqual(expected, actual, assumption)
	}

	assertThat("should not change source slice when item not found",
		[]string{"item1", "item2", "item3"},
		"item4",
		[]string{"item1", "item2", "item3"})

	assertThat("should exclude item from slice when item is found once at the end",
		[]string{"item1", "item2", "item3"},
		"item3",
		[]string{"item1", "item2"})

	assertThat("should exclude item from slice when item is found once in the middle",
		[]string{"item1", "item2", "item3"},
		"item2",
		[]string{"item1", "item3"})

	assertThat("should exclude item from slice when item is found once in the start",
		[]string{"item1", "item2", "item3"},
		"item1",
		[]string{"item2", "item3"})

	assertThat("should exclude item from slice when item is found multiple times",
		[]string{"item1", "item2", "item3", "item2", "item4", "item2"},
		"item2",
		[]string{"item1", "item3", "item4"})
}
