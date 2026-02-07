#ifndef MATRIX_HPP
#define MATRIX_HPP 1
#include <vector>

class Matrix final{
private:
	int columns_;
	int rows_;
	std::vector<std::vector<int>> matrix_;
	std::vector<std::vector<int>> create_matrix(int rows, int columns);
	
public:	
	int get_columns() const noexcept;
	int get_rows() const noexcept;
	
	Matrix& dot(const Matrix& m);
	Matrix& transporate();

	void show_size() const; 
	void show_matrix() const; 
	
	int& operator()(const int row, const int column);
	
	Matrix& operator*(Matrix& m);
	Matrix& operator*(const int n);
	Matrix& operator+(Matrix& m);
	Matrix& operator-(Matrix& m);

	Matrix(const int rows, const int columns);
	Matrix(const std::vector<std::vector<int>> matrix);
	Matrix(Matrix&&) = default;
	Matrix& operator=(Matrix&&) = default;
	Matrix(const Matrix&) = default;
	Matrix& operator=(const Matrix&) = default;
	~Matrix()= default;

};


#endif
