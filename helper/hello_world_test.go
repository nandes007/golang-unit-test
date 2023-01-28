package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Nandes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nandes")
		}
	})
	b.Run("Simanjuntak", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Simanjuntak")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nandes")
	}
}

func BenchmarkHelloWorldSteven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Steven")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Nandes",
			request:  "Nandes",
			expected: "Hello Nandes",
		},
		{
			name:     "Simanjuntak",
			request:  "Simanjuntak",
			expected: "Hello Simanjuntak",
		},
		{
			name:     "Steven",
			request:  "Steven",
			expected: "Hello Steven",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Nandes", func(t *testing.T) {
		result := HelloWorld("Nandes")
		require.Equal(t, "Hello Nandes", result, "Result must be 'Hello Nandes' ")
	})
	t.Run("Simanjuntak", func(t *testing.T) {
		result := HelloWorld("Simanjuntak")
		require.Equal(t, "Hello Simanjuntak", result, "Result must be 'Hello Simanjuntak' ")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before unit test")

	m.Run()

	// after
	fmt.Println("After unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Nandes")
	require.Equal(t, "Hello Nandes", result, "Result must be 'Hello Nandes' ")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nandes")
	require.Equal(t, "Hello Nandes", result, "Result must be 'Hello Nandes' ")
	fmt.Println("Di eksekusi with require")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nandes")
	assert.Equal(t, "Hello Nandes", result, "Result must be 'Hello Nandes' ")
	fmt.Println("Di eksekusi with assert")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")

	if result != "Hello World" {
		// error
		t.Error("Result must be 'Hello World' ")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldNandes(t *testing.T) {
	result := HelloWorld("Nandes")

	if result != "Hello Nandes" {
		// error
		t.Fatal("Result must be 'Hello Nandes' ")
	}

	fmt.Println("TestHelloWorldNandes Done")
}
