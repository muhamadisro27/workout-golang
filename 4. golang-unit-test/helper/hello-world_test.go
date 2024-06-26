package helper

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Roozy")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Roozy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Roozy")
		}
	})
	b.Run("Alo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Alo")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct{
			name string
		  request string
	}{
		{
			name : "HelloWorld(Roozy)",
			request : "Roozy",
		},
		{
			name : "HelloWorld(Alo)",
			request : "Alo",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestHelloWorld(t *testing.T) {
  result := HelloWorld("Roozy")

	assert.Equal(t, "Hello Roozy", result, "Result must be 'Hello Roozy'")

	fmt.Println("End Test")
}

func TestSkip(t *testing.T) {

	if runtime.GOOS == "darwin" {
		t.Skip("unit test tidak dijalankan pada OS darwin")
	}
	
	result := HelloWorld("Roozy")

	require.Equal(t, "Hello Roozy", result, "Result must be 'Hello Roozy'")

}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	fmt.Println("Initiate to databases")
	
	m.Run()


	fmt.Println("AFTER UNIT TEST")
	
}

func TestSubTest(t *testing.T) {
	t.Run("Roozy", func(t *testing.T) {
		result:= HelloWorld("Roozy")

		require.Equal(t, "Hello Roozy", result)
	})
	
	t.Run("qt", func(t *testing.T) {
		result:= HelloWorld("qt")

		require.Equal(t, "Hello qt", result)
	})
}

type TestingS struct {
	name, request, expected string
}

func TestHelloWorldTable(t *testing.T) {
	tests := []TestingS{
		{
			name : "HelloWorld(Roozy)",
			request: "Roozy",
			expected: "Hello Roozy",
		},
		{
			name : "HelloWorld(Isro)",
			request: "Isro",
			expected: "Hello Isro",
		},
		{
			name : "HelloWorld(Teguh)",
			request: "Teguh",
			expected: "Hello Teguh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
	
			assert.Equal(t, test.expected, result)
		})
	}
}