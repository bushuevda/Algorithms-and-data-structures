function dda_line(x1, y1, x2, y2){
    let i = 0;
    let x = Array();
    let y = Array();

    let x_start = Math.round(x1);
    let y_start = Math.round(y1);
    let x_end = Math.round(x2);
    let y_end = Math.round(y2);

    let length = Math.max(Math.abs(x_end - x_start), Math.abs(y_end - y_start));

    let dX = (x2 - x1) / length;
    let dY = (y2 - y1) / length;

    x.push(x1);
    y.push(y1);
    i++;
    
    while(i < length){
        x.push(x[i - 1] + dX);
        y.push(y[i - 1] + dY);
        i++;
    }

    x.push(x2);
    y.push(y2);

    let result = Array();
    result.push(x,y);

    return result;
}

export  {dda_line};