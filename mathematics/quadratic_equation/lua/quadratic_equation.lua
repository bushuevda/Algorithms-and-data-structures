local QE = {}

function QE.calc(a, b, c)
    local roots = {}
    local D = b ^ 2 - 4 * a * c
    
    if D == 0 then
        local x1 = (-b - math.sqrt(D)) / 2 * a
        table.insert(roots, x1)
    elseif D > 0 then
        local x1 = (-b - math.sqrt(D)) / 2 * a
        local x2 = (-b + math.sqrt(D)) / 2 * a
        table.insert(roots, x1)
        table.insert(roots, x2)
    end
    return roots
end


return QE