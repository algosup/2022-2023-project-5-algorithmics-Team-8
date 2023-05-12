#pragma once
#include <map>
#include <string>


struct Tank
{
  double capacity;
  std::map<std::string,double> formula;


  Tank(double size, std::string content)
  {
    capacity = size;
    if (content != "/") return;
    formula[content] =100;
  }
};
