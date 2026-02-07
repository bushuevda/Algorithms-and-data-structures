package graphics_algorithms

import kotlin.math.round
import kotlin.math.abs
import kotlin.math.max

fun dda_line(x1: Float, y1: Float, x2: Float, y2: Float): ArrayList<ArrayList<Float>>{
    var i: Int = 0
    var x: ArrayList<Float> = ArrayList()
    var y: ArrayList<Float> = ArrayList()

    var x_start: Float = round(x1)
    var y_start: Float = round(y1)
    var x_end: Float = round(x2)
    var y_end: Float = round(y2)

    var length: Float = max(abs(x_end - x_start), abs(y_end - y_start))

    var dX = (x2 - x1) / length
    var dY = (y2 - y1) / length

    x.add(x1)
    y.add(y1)
    i++

    while(i < length){
        x.add(x.get(i - 1) + dX)
        y.add(y.get(i - 1) + dY)
        i++
    }

    x.add(x2)
    y.add(y2)

    var result: ArrayList<ArrayList<Float>> = ArrayList()
    result.add(x)
    result.add(y)

    return result
}   