package benchmark

import (
    "bytes"
    "testing"

    "github.com/stretchr/testify/assert"
)

// func BenchmarkTemplate(b *testing.B) {
// 	// 与性能无关的代码
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		// 测试代码
// 	}

// 	b.StartTimer()
// 	// 与性能无关的代码
// }

func TestConcatStringByAdd(t *testing.T) {
    assert := assert.New(t)
    elements := []string{"1", "2", "3", "4", "5"}
    rvt := ""

    for _, elem := range elements {
        rvt += elem
    }

    assert.Equal("12345", rvt)
}

func TestConcatStringAddByBuffered(t *testing.T) {
    assert := assert.New(t)
    elements := []string{"1", "2", "3", "4", "5"}
    var buffer bytes.Buffer

    for _, elem := range elements {
        buffer.WriteString(elem)
    }

    assert.Equal("12345", buffer.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
    elements := []string{"1", "2", "3", "4", "5"}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        rvt := ""
        for _, elem := range elements {
            rvt += elem
        }
    }

    b.StartTimer()
}

func BenchmarkConcatStringByBufferedd(b *testing.B) {
    elements := []string{"1", "2", "3", "4", "5"}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        var buffer bytes.Buffer
        for _, elem := range elements {
            buffer.WriteString(elem)
        }
    }

    b.StartTimer()
}
