package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
		name: "Hanafi",
		request : "Hanafi",
		expected : "Hello Hanafi",
		},{
		name: "Adhi",
		request : "Adhi",
		expected : "Hello Adhi",
		},
	}

	for _,test := range tests{
		t.Run(test.name,func(t *testing.T) {
			result:= HelloWord(test.request)
			require.Equal(t,test.expected,result)
		})
	}
}


func TestSubTest(t *testing.T) {
	t.Run("Adhi",func(t *testing.T) {})
	result:= HelloWord("Hanafi")
	assert.Equal(t,"Hello Hanafi",result,"Result Must be 'Hello Hanafi'")
	t.Run("hanafi",func(t *testing.T) {})
	results:= HelloWord("Hanafi")
	assert.Equal(t,"Hello Hanafi",results,"Result Must be 'Hello Hanafi'")
}

func TestMain(m *testing.M){
	fmt.Println("Testing Before Unit Test")

	m.Run()

	fmt.Println("Testing After Unit TEst")

}

func TestHelloWord(t *testing.T) {
	result := HelloWord("Adhi")
	if result != "Hello Adhi" {
		t.Error("Result Must be 'Hello Adhi'")
	}
	fmt.Println("Hello Adhi Done")
}

func TestHelloW(t *testing.T) {
	result := HelloWord("Adhi")
	if result != "Hello Adhi" {
		t.Fatal("Result Must be 'Hello Adhi'")
	}
	fmt.Println("Hello Adhi Done")
}

// testing with testify

func TestHelloWordAssert(t *testing.T)  {
	result:= HelloWord("Hanafi")
	assert.Equal(t,"Hello Hanafi",result,"Result Must be 'Hello Hanafi'")
	fmt.Println("Testing Selesai")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWord("Hanafi")
	require.Equal(t,"Hello Hanafi", result, "Result must be 'Hello Hanafi'")
}

//testing skip

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWord("Hanafi")
	assert.Equal(t, "Hello Hanafi",result, "Result must Be 'Hello Hanafi'")
}
