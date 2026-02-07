
def brezenhem_line(x1: int, y1: int, x2: int, y2: int):
    x = list() #новый список x
    y = list() #новый список y
    dX = abs(x2 - x1) #длина отрезка по оси X
    dY = abs(y2 - y1) #длина отрезка по оси Y
    e = dY / dX - 0.5 #
    i = x1 # x
    j = y1 # y
    m = (x2 - x1) / (y2 - y1)  # угловой коэффициент

    iX = 1 if x1 < x2 else -1
    iY = 1 if y1 < y2 else -1

    while m <= dX:
        x.append(i)
        y.append(j)

        while e > 0:
            j = j + iY
            e = e - 1
        i = i + iX
        e = e + dY / dX
        m = m + 1
    x.append(x2)
    y.append(y2)
    print(x)
    print(y)

if __name__ == "__main__":
    brezenhem_line(1, 1, 2,-5)