package oct2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCalendarThree(t *testing.T) {
	c := NewMyCalendarThree()
	assert.Equal(t, 1, c.Book(10, 20))
	assert.Equal(t, 1, c.Book(50, 60))
	assert.Equal(t, 2, c.Book(10, 40))
	assert.Equal(t, 3, c.Book(5, 15))
	assert.Equal(t, 3, c.Book(5, 10))
	assert.Equal(t, 3, c.Book(25, 55))
	assert.Equal(t, 4, c.Book(10, 15))
}
