#pragma once
#include <map>
#include <string>

using namespace std;

struct Tank
{
  double capacity;
  map<std::string,double> formula;


  Tank(double size, std::string content)
  {
    capacity = size;
    if (content != "/") return;
    formula[content] =100;
  }
};
