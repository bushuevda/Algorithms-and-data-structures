
def dda_line(x1: float, y1: float, x2: float, y2: float) -> tuple[list[int], list[int]]:
    i = 0
    x = list()
    y = list()
    x_start = round(x1)
    y_start = round(y1)
    x_end = round(x2)
    y_end = round(y2)
    L = max(abs(x_end - x_start), abs(y_end - y_start))
    dX = (x2 - x1) / L
    dY = (y2 - y1) / L
    x.append(float(x1))
    y.append(float(y1))
    i += 1
    while i < L:
        x.append(float(x[i - 1] + dX))
        y.append(float(y[i - 1] + dY))
        i += 1
    x.append(float(x2))
    y.append(float(y2))
    return x, y


if __name__ == "__main__":
    print(dda_line(4, 3, 2, 1))
 
