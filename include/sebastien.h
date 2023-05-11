#pragma once
#include <vector>
#include <map>
#include <fstream>
#include <iostream>
#include <string>
#include "tank.h"
#include "step.h"

struct Result
{
  std::map<std::string,Tank> tanks;
  std::map<std::string,double> formula;
};


void readline(std::string line, std::map<std::string,Tank>& tanks, std::map<std::string,double>& formula,double& total);

Result parseconfig(std::string filepath); 
