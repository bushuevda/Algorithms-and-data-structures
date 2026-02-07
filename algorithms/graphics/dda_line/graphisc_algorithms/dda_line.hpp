#ifndef DDA_LINE_HPP
#define DDA_LINE_HPP 1
#include <iostream>
#include <vector>

namespace graphics_algorithms{

    std::pair<std::vector<float>, std::vector<float>> dda_line(float x1, float y1, float x2, float y2);

};
#endif