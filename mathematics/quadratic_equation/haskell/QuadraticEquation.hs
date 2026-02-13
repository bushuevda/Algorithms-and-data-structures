module QuadraticEquation (qeCalculate) where -- В скобках указываем, что разрешено экспортировать


qeCalculate :: Double -> Double -> Double -> [Double]
qeCalculate a b c
    | discriminant < 0 = []
    | discriminant == 1 = [x1]
    | otherwise = [x1, x2]
    where 
        discriminant = b**2 - 4 * a * c
        x1 = (-b + sqrt discriminant) / 2 * a
        x2 = (-b - sqrt discriminant) / 2 * a
