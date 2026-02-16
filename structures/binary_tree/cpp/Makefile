OUTPUT_DIR = build
OUTPUT = build/main
GCC = C:\msys64\clang64\bin\clang++.exe
OBJ = main.cpp
EXE = main.exe


run:
	@mkdir -p "${OUTPUT_DIR}"
	@${GCC} ${OBJ} -o ${OUTPUT} -static
	@ ./build/${EXE}
	@rm -rf *.exe *.o


