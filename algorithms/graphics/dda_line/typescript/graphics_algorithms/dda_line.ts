function dda_line(x1: number, y1: number, x2: number, y2: number): Array<Array<number>>{
    let i: number = 0;
    let x: Array<number> = Array();
    let y: Array<number> = Array();

    let x_start: number = Math.round(x1);
    let y_start: number = Math.round(y1);
    let x_end: number = Math.round(x2);
    let y_end: number = Math.round(y2);

    let length: number = Math.max(Math.abs(x_end - x_start), Math.abs(y_end - y_start));

    let dX: number = (x2 - x1) / length;
    let dY: number = (y2 - y1) / length;

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

    let result: Array<Array<number>> = Array();

    result.push(x, y);

    return result;
}


export {dda_line};