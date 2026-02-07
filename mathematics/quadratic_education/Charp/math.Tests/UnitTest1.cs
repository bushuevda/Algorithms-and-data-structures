
using System.Runtime.CompilerServices;
using Mathematics;
using Xunit;
namespace math.Tests;
public class UnitTest1
{

    bool CompareDouble(double n1, double n2, double eps = 1.0E-4){
        return Math.Abs(n1 - n2) < eps;
    }

    bool UtilTestCalcABC(double a, double b, double c, double x1 = double.NaN, double x2 = double.NaN){
       List<double> res =  QuadraticEquation.Calculate(a, b, c);
        if(res.Count == 0 && double.IsNaN(x1) && double.IsNaN(x2)){
            return true;
        } else if(res.Count == 1 && CompareDouble(res[0], x1) && double.IsNaN(x2)){
            return true;
        } else if(res.Count == 2 && CompareDouble(res[0], x1) && CompareDouble(res[1], x2)){
            return true;
        } else {
            return false;
        }
    }


    [Fact]
    public void TestCalculateTwoRoots()
    {
        List<List<double>> testData = [
            //a = -1, b = 6, c = 7, x1 = 7, x2 = -1
            [-1, 6, 7, 7, -1],
            //a = 1, b = 5, c = -6, x1 = -6, x2 = 1
            [1, 5, -6, -6, 1],
            //a = 1, b = 1, c = -6, x1 = -3, x2 = 2
            [1, 1, -6, -3, 2],
            //a = 4, b = 1, c = -1, x1 = -0.64039, x2 = 0.39039
            [4, 1, -1, -0.64039, 0.39039],
            //a = 4, b = 11, c = 3, x1 = -2.4430, x2 = -0.30700
            [4, 11, 3, -2.4430, -0.30700]
        ];

        List<bool> results = [];
        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }



    [Fact]
    public void TestCalculateOneRoot()
    {
        List<List<double>> testData = [
            //a = 9, b = -6, c = 1, x1 =  0.33333, x2 = NaN
            [9, -6, 1, 0.33333, double.NaN],
            //a = 4, b = 4, c = 1, x1 = -0.5, x2 = NaN
            [4, 4, 1, -0.5, double.NaN],
            //a = 1, b = -6, c = 9, x1 = 3, x2 = NaN
            [1, -6, 9, 3, double.NaN],
            //a = 1, b = 12, c = 36, x1 = -6, x2 = NaN
            [1, 12, 36, -6, double.NaN],
            //a = 1, b = 4, c = 4, x1 = -2, x2 = NaN
            [1, 4, 4, -2, double.NaN]
        ];

        List<bool> results = [];
        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }



    [Fact]
    public void TestCalculateWithOutRoots()
    {
        List<List<double>> testData = [
            //a = 4, b = 2, c = 3, x1 = NaN, x2 = NaN
            [4, 2, 3, double.NaN, double.NaN],
            //a = 7, b = 6, c = 3, x1 = NaN, x2 = NaN
            [7, 6, 3, double.NaN, double.NaN],
            //a = 7, b = 6, c = 7, x1 = NaN, x2 = NaN
            [7, 6, 7, double.NaN, double.NaN],
            //a = -11, b = 1, c = -3, x1 = NaN, x2 = NaN
            [-11, 1, -3, double.NaN, double.NaN],
            //a = 21, b = -11, c = 5, x1 = NaN, x2 = NaN
            [21, -11, 5, double.NaN, double.NaN]
        ];

        List<bool> results = [];

        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }



    [Fact]
    public void TestA()
    {
        List<List<double>> testData = [
            //a = -1, b = 0, c = 0, x1 = 0, x2 = NaN
            [-1, 0, 0, 0, double.NaN],
            //a = 5, b = 0, c = 0, x1 = 0, x2 = NaN
            [5, 0, 0, 0, double.NaN],
            //a = 12, b = 0, c = 0, x1 = 0, x2 = NaN
            [12, 0, 0, 0, double.NaN],
            //a = 4, b = 0, c = 0, x1 = 0, x2 = NaN
            [4, 0, 0, 0, double.NaN],
            //a = 7, b = 0, c = 0, x1 = 0, x2 = NaN
            [7, 0, 0, 0, double.NaN]
        ];

        List<bool> results = [];
        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }


    [Fact]
    public void TestAB()
    {
        List<List<double>> testData = [
            //a = 12, b = 4, c = 0, x1 = -0.33333, x2 = 0
            [12, 4, 0, -0.33333, 0],
            //a = 5, b = -4, c = 0, x1 = 0, x2 = 0.8
            [5, -4, 0, 0, 0.8],
            //a = 5, b = 5, c = 0, x1 = -1, x2 = 0
            [5, 5, 0, -1, 0],
            //a = 13, b = 15, c = 0, x1 = -1.1538, x2 = 0
            [13, 15, 0, -1.1538, 0],
            //a = -13, b = 15, c = 0, x1 = 1.1538, x2 = 0
            [-13, 15, 0, 1.1538, 0]
        ];

        List<bool> results = [];
        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }



    [Fact]
    public void TestAC()
    {
        List<List<double>> testData = [
            //a = -1, b = 0, c = 7, x1 = 2.6458, x2 = -2.6458
            [-1, 0, 7, 2.6458, -2.6458],
            //a = 7, b = 0, c = 7, x1 = NaN, x2 = NaN
            [7, 0, 7, double.NaN, double.NaN],
            //a = 8, b = 0, c = 7, x1 = NaN, x2 = NaN
            [8, 0, 7, double.NaN, double.NaN],
            //a = 1, b = 0, c = -17, x1 = -4.1231, x2 = 4.1231
            [1, 0, -17, -4.1231, 4.1231],
            //a = 3, b = 0, c = -5, x1 = -1.2910, x2 = 1.2910
            [3, 0, -5, -1.2910, 1.2910]
        ];

        List<bool> results = [];
        foreach(var arr in testData){
            results.Add(UtilTestCalcABC(arr[0], arr[1], arr[2], arr[3], arr[4]));
        }
        List<bool> excepted = Enumerable.Repeat(true, testData.Count).ToList();

        Assert.Equal(excepted, results);
    }



}
