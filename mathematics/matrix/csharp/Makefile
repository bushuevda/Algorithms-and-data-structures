CSPROJ_MATH = "mathematics/mathematics.csproj"
CSPROJ_PROG = "prog/prog.csproj"
CSPOJ_TESTS = "tree.TestsBst/tree.TestsBst.csproj"

#Добавление зависимости dll в проект dotnet
add_ref:
	dotnet add ${CSPROJ_PROG} reference ${CSPROJ_MATH}

#Удаление зависимости из проекта dotnet
del_ref:
	dotnet remove ${CSPROJ_PROG} reference ${CSPROJ_BST}

#Создание библиотеки dotnet
NAME_NEW_LIB = "mathematics"
new_lib:
	dotnet new classlib -n ${NAME_NEW_LIB}

#Создание консольного проекта dotnet
NAME_NEW_CONSOLE = "prog"
new_console:
	dotnet new console -n ${NAME_NEW_CONSOLE}


#Создание проекта тестов dotnet xunit
NAME_PROJECT_TESTS = "tree.TestsBst"
new_tests:
	dotnet new xunit -o ${NAME_PROJECT_TESTS}

#Запуск тестов
run_tests:
	dotnet test ${CSPOJ_TESTS}

#Запуск проекта
run:
	dotnet run --project ${CSPROJ_PROG}