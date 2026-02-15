OUTPUT_DIR = build
OUTPUT = build/main
GCC = g++
OBJ = *.cpp
EXE = main.exe


run:
	@mkdir -p "${OUTPUT_DIR}"
	@${GCC} ${OBJ} -o ${OUTPUT}
	@ ./build/${EXE}
	@rm -rf *.exe *.o


