package helper

import ( 
	"testing"
	"fmt"
	"runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldNinja(t *testing.T) {
	result := HelloWorld("Ninja")

	if result != "Hello Ninja" {
		// error
		// t.Fail()	
		t.Error("Result must be `Hello Ninja`")
	}

	fmt.Println("TestHelloWorldNinja done")
}

func TestHelloWorldGolang(t *testing.T) {
	result := HelloWorld("Golang")

	if result != "Hello Golang" {
		// error
		// t.FailNow()
		t.Error("Result must be `Hello Golang`")
		// panic("Result is not 'Hello Golang'")
	}
	fmt.Println("TestHelloWorldGolang done")
}

// Assert test
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
	fmt.Println("TestHelloWorldAssert with Assert Done")
}

// Require test
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
	fmt.Println("TestHelloWorldRequire with Require Done")
}

// Skip test
func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
	fmt.Println("TestHelloWorldRequire with Require Done")
}

// Main test
func TestMain(m *testing.M) {
	// before test
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	// after test
	fmt.Println("AFTER UNIT TEST")
}

// Subtest
func TestSubTest(t *testing.T) {
	t.Run("Coder", func(t *testing.T) {
		result := HelloWorld("Coder")
		require.Equal(t, "Hello Coder", result, "Result must be 'Hello Coder'")
	})
	t.Run("Programmer", func(t *testing.T) {
		result := HelloWorld("Programmer")
		require.Equal(t, "Hello Programmer", result, "Result must be 'Hello Programmer'")		
	})
	// t.Run("Developer", func(t *testing.T) {
	// 	result := HelloWorld("Developer")
	// 	require.Equal(t, "Hello Developer", result, "Result must be 'Hello Programmer'")
	// })
}

// Table test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Oggy",
			request: "Oggy",
			expected: "Hello Oggy",
		},
		{
			name: "Jack",
			request: "Jack",
			expected: "Hello Jack",
		},
		{
			name: "Bob",
			request: "Bob",
			expected: "Hello Bob",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// BenchmarkHelloWorld
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ninja")
	}
}
// BenchmarkHelloWorldCoder
func BenchmarkHelloWorldCoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ninja Coder")
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("r3ndy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("r3ndy")
		}		
	})
	b.Run("wnnaCry", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("wnnaCry")
		}		
	})
}

// Table Benchmark
func BenchmarkTableHelloWorld(b *testing.B) {
	tests := []struct {
		name string
		request string
	} {
		{
			name: "Oggy",
			request: "Oggy",
		},
		{
			name: "Jack",
			request: "Jack",
		},
		{
			name: "Bob",
			request: "Bob",
		},
		{
			name: "Rendy",
			request: "Rendy",
		},
		{
			name: "Rendy dinar",
			request: "Rendy dinar",
		},
		{
			name: "Umbu Theofilus Dendimara",
			request: "Umbu Theofilus Dendimara",
		},
	}

	for _, benchmark := range tests {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}		
		})
	}
}
