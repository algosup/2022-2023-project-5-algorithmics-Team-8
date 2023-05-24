#pragma once 
#include <vector>
#include <map>
#include <fstream>
#include <iostream>
#include <string>

struct Formula
{
    std::string name;
    float quantity;
    std::vector<std::string> inputs;
    std::vector<float> inputs_quantity;
    float output_quantity;
    float output;
    bool is_solved;
};

struct Tank
{
    std::string name;
    float quantity;
    float quantity_left;
    bool is_solved;
};

struct Result
{
  std::map<std::string,Tank*> tanks;
  std::map<std::string,double> formula;
};


std::vector<Formula> formulas;
std::vector<Tank> tanks;
