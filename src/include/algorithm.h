#pragma once 
#include <vector>
#include <map>
#include <fstream>
#include <iostream>
#include <string>
#include "tank.h"


struct Formula
{
  std::string name;
  double quantity;
  std::vector<std::string> inputs;
  std::vector<double> inputs_quantity;
  double output_quantity;
  double output;
  bool is_solved;
};

std::vector<Formula> formulas;
std::vector<Tank> tanks;
