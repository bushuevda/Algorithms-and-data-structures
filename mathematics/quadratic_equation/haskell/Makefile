OUTPUT_DIR = build
MAIN_FILE = Main.hs
EXE = main.exe

run:
	@mkdir -p "${OUTPUT_DIR}"
	@ghc -i. --make ${MAIN_FILE} -outputdir "${OUTPUT_DIR}" -o "${OUTPUT_DIR}/${EXE}"
	@ ./build/${EXE}
	@rm -rf *.exe *.hi *.o

init: 
	cabal init --minimal --simple
