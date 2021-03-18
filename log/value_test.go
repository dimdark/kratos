package log

import "testing"

func TestValue(t *testing.T) {
	logger := With(DefaultLogger, "caller", Caller(4))
	logger.Print("message", "helloworld")
	testValues2(logger)
}

func testValues2(logger Logger) {
	logger.Print("message", "testValues2")
	testValues3(logger)
}

func testValues3(logger Logger) {
	logger.Print("message", "testValues3")
	testValues4(logger)
}

func testValues4(logger Logger) {
	logger.Print("message", "testValues4")
	testValues5(logger)
}

func testValues5(logger Logger) {
	logger.Print("message", "testValues5")
}

