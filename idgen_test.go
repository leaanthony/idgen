package idgen

import (
	"github.com/matryer/is"
	"testing"
)

func TestGenerator_NewID(t *testing.T) {

	is2 := is.New(t)

	generator := New()

	var i uint
	for i = 0; i < 10; i++ {
		id, err := generator.NewID()
		is2.NoErr(err)
		is2.Equal(id, i)
	}

}

func TestGenerator_ReleaseID(t *testing.T) {

	is2 := is.New(t)

	// Fill the generator
	generator := NewWithMaximum(10)
	var i uint = 0
	for i = 0; i < 10; i++ {
		_, err := generator.NewID()
		is2.NoErr(err)
	}

	// Next should be an error
	val, err := generator.NewID()
	is2.Equal(val, uint(0))
	is2.True(err != nil)

	// Release
	generator.ReleaseID(5)

	// Get new ID
	id, err := generator.NewID()
	is2.Equal(id, uint(5))
	is2.NoErr(err)
}

//
//func TestNew(t *testing.T) {
//	tests := []struct {
//		name string
//		want *Generator
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := New(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("New() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
