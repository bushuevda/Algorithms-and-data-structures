using System.Collections.Generic;
using System.Numerics;

namespace mathematics;

public sealed class Matrix <T>  where T: INumber<T>
{
    public int ColumnsCount{get; set;}
    public int RowsCount{get; set;}

    private List<List<T>> matrix_;


    public Matrix(int rowsCount, int columnsCount){
        RowsCount = rowsCount;
        ColumnsCount = columnsCount;
        List<List<T>> newMatrix = new List<List<T>>();
        for (int i = 0; i < rowsCount; i++){
            newMatrix.Add(Enumerable.Repeat(default(T), columnsCount).ToList()!);
        }
        matrix_ = newMatrix;
    }

    public Matrix(List<List<T>> matrix){
        RowsCount = matrix.Count;
        ColumnsCount = matrix[0].Count;
        matrix_ = matrix;
    }

    public Matrix<T> Dot(Matrix<T> m){
        if(ColumnsCount != m.RowsCount)
            throw new Exception("Matrix size columns not equil size rows!");
        List<List<T>> dot_matrix = new List<List<T>>();

        for (int i = 0; i < RowsCount; i++){
            List<T> n_row = new List<T>();
            for (int j = 0; j < m.ColumnsCount; j++){
                T? sum = default!;
                for (int k = 0; k < m.RowsCount; k++){
                    sum += matrix_[i][k] * m.matrix_[k][j];
                }
                n_row.Add(sum);
            }
            dot_matrix.Add(n_row);
        }

        Matrix<T> res = new Matrix<T>(dot_matrix);
        return res;
    }


    public static Matrix<T> operator * (Matrix<T> m1, Matrix<T> m2){
        Matrix<T> res = m1.Dot(m2);
        return res;
    }


    public static Matrix<T> operator * (Matrix<T> m, dynamic n){
        Matrix<T> res = new Matrix<T>(m.RowsCount, m.ColumnsCount);
        for(int i = 0; i < m.RowsCount; i++){
            for(int j = 0; j < m.ColumnsCount; j++){
                res.matrix_[i][j] = m.matrix_[i][j] * n;
            }
        }
        return res;
    }


    public static Matrix<T> operator - (Matrix<T> m1, Matrix<T> m2){

        if(m1.ColumnsCount != m2.ColumnsCount || m1.RowsCount != m2.RowsCount)
            throw new Exception("Matrix size columns not equil size rows!");

        Matrix<T> res = new Matrix<T>(m1.RowsCount, m1.ColumnsCount);
        for (int i = 0; i < m1.RowsCount; i++){
            for (int j = 0; j < m1.ColumnsCount; j++){
                res.matrix_[i][j] = m1.matrix_[i][j] - m2.matrix_[i][j];
            }
        }

        return res;
    }


    public static Matrix<T> operator + (Matrix<T> m1, Matrix<T> m2){

        if(m1.ColumnsCount != m2.ColumnsCount || m1.RowsCount != m2.RowsCount)
            throw new Exception("Matrix size columns not equil size rows!");

        Matrix<T> res = new Matrix<T>(m1.RowsCount, m1.ColumnsCount);
        for (int i = 0; i < m1.RowsCount; i++){
            for (int j = 0; j < m1.ColumnsCount; j++){
                res.matrix_[i][j] = m1.matrix_[i][j] + m2.matrix_[i][j];
            }
        }

        return res;
    }

    public void ShowMatrix(){
        for(int i = 0; i < RowsCount; i++){
            for(int j = 0; j < ColumnsCount; j++){
                Console.Write($"{matrix_[i][j]} \t");
            }
                Console.Write("\n");
        }
    }


    public Matrix<T> Transporate(){
        List<List<T>> transporateMatrix = new List<List<T>>();

        for(int i = 0; i < ColumnsCount; i++){
            List<T> t_column = new List<T>();
            for(int j = 0; j < RowsCount; j++){
                t_column.Add(matrix_[j][i]);
            }
            transporateMatrix.Add(t_column);
        }

        Matrix<T> res = new Matrix<T>(transporateMatrix);
        return res;
    }


}


