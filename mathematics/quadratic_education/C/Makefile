QE = quadratic_education/quadratic_education.c
MAIN = main.c
GCC = gcc
SOURCES = $(QE) $(MAIN)

OUTPUT_DIR = "builds"
OUTPUT_FILE = "main"
PROGRAM = $(OUTPUT_DIR)/$(OUTPUT_FILE) 

build: $(SOURCES)
	@rm -r $(OUTPUT_DIR)
	@mkdir $(OUTPUT_DIR)
	@$(GCC) $^ -o $(PROGRAM)
	@$(PROGRAM)

