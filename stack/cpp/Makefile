files_to_delete = $(wildcard test.exe main.exe)

run:
	./main

build: clean test
	g++ main.cpp stack/stack.cpp -o main

clean:
	rm -f $(files_to_delete)
	
test: 
	g++ -std=c++11 tests/test_stack.cpp stack/stack.cpp -o test && ./test && rm test.exe 