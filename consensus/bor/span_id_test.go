package bor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpanIDAt(t *testing.T) {
	assert.Equal(t, uint64(0), SpanIDAt(0))
	assert.Equal(t, uint64(0), SpanIDAt(1))
	assert.Equal(t, uint64(0), SpanIDAt(2))
	assert.Equal(t, uint64(0), SpanIDAt(zerothSpanEnd))
	assert.Equal(t, uint64(1), SpanIDAt(zerothSpanEnd+1))
	assert.Equal(t, uint64(1), SpanIDAt(zerothSpanEnd+2))
	assert.Equal(t, uint64(1), SpanIDAt(6655))
	assert.Equal(t, uint64(2), SpanIDAt(6656))
	assert.Equal(t, uint64(2), SpanIDAt(6657))
	assert.Equal(t, uint64(2), SpanIDAt(13055))
	assert.Equal(t, uint64(3), SpanIDAt(13056))
	assert.Equal(t, uint64(6839), SpanIDAt(43763456))
}

func TestSpanEndBlockNum(t *testing.T) {
	assert.Equal(t, uint64(zerothSpanEnd), SpanEndBlockNum(0))
	assert.Equal(t, uint64(6655), SpanEndBlockNum(1))
	assert.Equal(t, uint64(13055), SpanEndBlockNum(2))
	assert.Equal(t, uint64(43769855), SpanEndBlockNum(6839))
}