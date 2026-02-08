import math

class QuadraticEquation:
    a_: int | float
    b_: int | float
    c_: int | float
    x1_: int | float
    x2_: int | float
    root_counts_: int
    
    def __init__(self, a, b, c, x1 = None, x2 = None, root_counts = None):
        self.a_, self.b_, self.c_ = a, b, c
        if(x1):
            self.x1_ = x1
        if(x2):
            self.x2_ = x2
        if(root_counts):
            self.root_counts_ = root_counts
    

class QudarticEquationUtil:
    def calculate(a, b, c) -> QuadraticEquation:
        qe = QuadraticEquation(a, b, c)
        D = qe.a_ * qe.b_ - 4 * qe.a_ * qe.c_
        if D > 0:
            qe.x1_ = (-qe.b_ - math.sqrt(D)) / (2 * qe.a_)
            qe.x2_ = (-qe.b_ + math.sqrt(D)) / (2 * qe.a_)
            qe.root_counts_ = 2
        elif D == 0:
            qe.x1_ = -qe.b_ / (2 * qe.a_)
            qe.x2_ = math.nan
            qe.root_counts_ = 1
        else:
            qe.x1_ = math.nan
            qe.x2_ = math.nan
            qe.root_counts_ = 0
        return qe


if __name__ == "__main__":
    print(QudarticEquationUtil.calculate(-1,6,7).x1_)