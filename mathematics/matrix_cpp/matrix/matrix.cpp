#include "matrix.hpp"
#include <iostream>
#include <cassert>

Matrix::Matrix(const int rows, const int columns){
	columns_ = columns;
	rows_ = rows;
	matrix_ = create_matrix(rows, columns);
}

Matrix::Matrix(const std::vector<std::vector<int>> matrix){
	columns_ = matrix[0].size();
	rows_ = matrix.size();
	matrix_ = matrix;
}


std::vector<std::vector<int>> Matrix::create_matrix(int rows, int columns){
	std::vector<std::vector<int>> n_matrix;
	n_matrix.resize(rows);
	for(int i = 0; i < rows; i++){
		n_matrix[i].resize(columns);
		std::fill(n_matrix[i].begin(), n_matrix[i].end(), 0);
	}
	return n_matrix;
}

int Matrix::get_columns() const noexcept{
	return columns_;
}

int Matrix::get_rows() const noexcept{
	return rows_;
}

void Matrix::show_size() const{
	std::cout<<"Rows: " << matrix_.size()<<". Columns: "<<matrix_[0].size()<<std::endl;
}


void Matrix::show_matrix() const{
	for(int i = 0; i < matrix_.size(); i++){
		for(int j = 0; j < matrix_[i].size(); j++){
			std::cout<<matrix_[i][j]<<'\t';
		}
		std::cout<<std::endl;
	}
}


int& Matrix::operator()(const int row, const int column){
	assert(column >= 0 && column <= columns_);
	assert(row >= 0 && row <= rows_);
	return matrix_[row][column];
}


Matrix& Matrix::dot(const Matrix& m){
	assert(columns_ == m.get_rows() 
			&& "matrix1 size columns must be equil matrix2 size rows!");
	std::vector<std::vector<int>> dot_matrix;
	
	for(int i = 0; i < rows_; i++){
		std::vector<int> n_row;
		for(int j = 0; j < m.get_columns(); j++){
			int sum = 0;
			for(int k = 0; k < m.get_rows(); k++){
				sum += matrix_.at(i).at(k) * m.matrix_.at(k).at(j);
			}
			n_row.push_back(sum);
		}
		dot_matrix.push_back(n_row);
	}
	matrix_ = dot_matrix;
	return *this;
}


Matrix& Matrix::transporate(){
	std::vector<std::vector<int>> transporate_matrix;
	for(int i = 0; i < columns_; i++){
		std::vector<int> t_column;
		for(int j = 0; j < rows_; j++){
			t_column.push_back(matrix_[j][i]);
		}
		transporate_matrix.push_back(t_column);
	}
	matrix_ = transporate_matrix;
	rows_ = matrix_.size();
	columns_ = matrix_[0].size();
	return *this;
}



Matrix& Matrix::operator*(Matrix& m){
	assert(columns_ == m.get_rows() 
			&& "matrix1 size columns must be equil matrix2 size rows!");
	std::vector<std::vector<int>> dot_matrix;
	
	for(int i = 0; i < rows_; i++){
		std::vector<int> n_row;
		for(int j = 0; j < m.get_columns(); j++){
			int sum = 0;
			for(int k = 0; k < m.get_rows(); k++){
				sum += matrix_.at(i).at(k) * m.matrix_.at(k).at(j);
			}
			n_row.push_back(sum);
		}
		dot_matrix.push_back(n_row);
	}
	matrix_ = dot_matrix;
	rows_ = matrix_.size();
	columns_ = matrix_[0].size();
	return *this;
}


Matrix& Matrix::operator+(Matrix& m){
	assert(columns_ == m.get_columns() 
			&& rows_ == m.get_rows() 
			&& "matrix1 size columns must be equil matrix2 size rows!");
	
	for(int i = 0; i < rows_; i++){
		for(int j = 0; j < columns_; j++){
			matrix_[i][j] += m.matrix_[i][j];
		}
	}
	
	return *this;
}


Matrix& Matrix::operator-(Matrix& m){
	assert(columns_ == m.get_columns() 
			&& rows_ == m.get_rows() 
			&& "matrix1 size columns must be equil matrix2 size rows!");
	
	for(int i = 0; i < rows_; i++){
		for(int j = 0; j < columns_; j++){
			matrix_[i][j] -= m.matrix_[i][j];
		}
	}
	
	return *this;
}


Matrix& Matrix::operator*(int n){

	for(int i = 0; i < rows_; i++){
		for(int j = 0; j < columns_; j++){
			matrix_[i][j] *= n;
		}
	}
	
	return *this;
}
