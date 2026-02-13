import QuadraticEquation (qeCalculate)

main :: IO ()
main = do
    let result = qeCalculate (-6) 2 3
    print result
