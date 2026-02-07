#include "matrix/matrix.hpp"
#include <iostream>
#include <memory>

int main(){
	std::unique_ptr<Matrix> m1{
	new Matrix({
		{1,2,3},
		{4,5,6},
		{7,8,9}
		})
	};
	m1->transporate();
	m1->show_matrix();
	return 0;
}
