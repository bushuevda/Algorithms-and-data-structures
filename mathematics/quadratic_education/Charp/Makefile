CSPROJ_MATH = "math/math.csproj"
CSPROJ_PROG = "prog/prog.csproj"
CSPOJ_TESTS = "math.Tests/math.Tests.csproj"

#Добавление зависимости dll в проект dotnet
add_ref:
	dotnet add ${CSPOJ_TESTS} reference ${CSPROJ_MATH}

#Удаление зависимости из проекта dotnet
del_ref:
	dotnet remove ${CSPROJ_PROG} reference ${CSPROJ_MATH}

#Создание библиотеки dotnet
NAME_NEW_LIB = "name_lib"
new_lib:
	dotnet new classlib -n ${NAME_NEW_LIB}

#Создание консольного проекта dotnet
NAME_NEW_CONSOLE = "name_console"
new_console:
	dotnet new console -n ${NAME_NEW_CONSOLE}


#Создание проекта тестов dotnet xunit
NAME_PROJECT_TESTS = "math.Tests"
new_tests:
	dotnet new xunit -o ${NAME_PROJECT_TESTS}

#Запуск тестов
run_tests:
	dotnet test ${CSPOJ_TESTS}


