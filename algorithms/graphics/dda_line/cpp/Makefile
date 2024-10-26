files_to_delete = $(wildcard test.exe example.exe)

build: clean test
	g++ --std=c++11 example.cpp graphisc_algorithms/dda_line.cpp -o example

run:
	./example

clean:
	rm -f $(files_to_delete)
	
test: clean
	g++ --std=c++11 tests/test_dda_line.cpp graphisc_algorithms/dda_line.cpp -o test && ./test && rm test.exe

example: clean	
	g++ --std=c++11 example.cpp graphisc_algorithms/dda_line.cpp -o example && ./example
