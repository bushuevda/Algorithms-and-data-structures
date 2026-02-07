QE = dda_line.c
MAIN = main.c
GCC = gcc
SOURCES = $(QE) $(MAIN)

OUTPUT_DIR = "builds"
OUTPUT_FILE = "main"
PROGRAM = $(OUTPUT_DIR)/$(OUTPUT_FILE) 

build: $(SOURCES)
	@rm -rf $(OUTPUT_DIR)
	@mkdir -p $(OUTPUT_DIR)
	@$(GCC) $^ -o $(PROGRAM)
	@$(PROGRAM)

