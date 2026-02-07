#include "graphisc_algorithms/dda_line.hpp"
#include <iostream>
#include <vector>

int main(){
    std::pair<std::vector<float>, std::vector<float>> result = graphics_algorithms::dda_line(13, 0, -4, 0);
    for(auto x : result.first){
        std::cout<< x << " ";
    }
    std::cout<<std::endl;
    for(auto y : result.second){
        std::cout<< y << " ";
    }
}