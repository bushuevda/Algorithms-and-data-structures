#include <iostream>
#include <vector>
#include "dda_line.hpp"
#include <cmath>
#include <algorithm>

std::pair<std::vector<float>, std::vector<float>> graphics_algorithms::dda_line(float x1, float y1, float x2, float y2){
    int i = 0;

    std::vector<float> x;
    std::vector<float> y;

    float x_start = roundf(x1);
    float y_start = roundf(y1);
    float x_end = roundf(x2);
    float y_end = roundf(y2);

    int length = std::max(abs(x_end - x_start), abs(y_end - y_start));

    float dX = (x2 - x1) / length;
    float dY = (y2 - y1) / length;
    
    x.push_back(x1);
    y.push_back(y1);
    i++;

    while(i < length){
        x.push_back(x[i - 1] + dX);
        y.push_back(y[i - 1] + dY);
        i++;
    }

    x.push_back(x2);
    y.push_back(y2);

    std::pair<std::vector<float>, std::vector<float>> result;
    result.first = x;
    result.second = y;

    return result;
}
