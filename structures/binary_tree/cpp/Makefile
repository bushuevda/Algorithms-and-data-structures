OUTPUT_DIR = build
OUTPUT = build/main
CXX = C:\msys64\clang64\bin\clang++.exe
OBJ = main.cpp
EXE = main.exe


run:
	@mkdir -p "${OUTPUT_DIR}"
	@${CXX} ${OBJ} -o ${OUTPUT} -static
	@ ./build/${EXE}
	@rm -rf *.exe *.o


