using mathematics;


Matrix<int> m1 = new Matrix<int>([
    [1,2,3],
    [4,5,6],
    [7,8,9],
]);

Matrix<int> m2 = new Matrix<int>([
    [9,8,7],
    [6,5,4],
    [3,2,1],
]);

var res = m1 - m2;
m1.ShowMatrix();
res.ShowMatrix();