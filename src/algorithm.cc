#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <algorithm>
#include <iterator>
#include <map>
#include <fstream>
#include <iomanip>
#include <cmath>
#include "include/tank.h"
#include "include/algorithm.h"
#include "include/sebastien.h"


void Find_Tank(double amount)
{
  for (int i = 0; i < tanks.size(); i++)
  {
    if (tanks[i].capacity == amount)
    {
      tanks[i].formula["/"] = 100;
      return;
    }
  }
}

std::vector<Formula> Split_Formula(Formula F)
{
  std::vector<Formula> formulas;
  for (int i = 0; i < F.inputs.size(); i++)
  {
    Formula f;
    f.name = F.inputs[i];
    f.quantity = F.inputs_quantity[i];
    f.output_quantity = F.output_quantity;
    f.output = F.output;
    f.is_solved = false;
    formulas.push_back(f);
  }
  return formulas;
}

void Solve(Formula F, double a)
{
  Find_Tank(a);
  std::vector<Formula> formulas = Split_Formula(F);
  for (int i = 0; i < formulas.size(); i++)
  {
    Solve(formulas[i], formulas[i].quantity);
  }
}

//get args
int main(int argc, char *argv[])
{
  std::string path = std::string(argv[1]);
  std::cout << path << std::endl;

  
  FILE *file = fopen(path.c_str(), "r");

  std::map<std::string, Tank *> *tanks_map = new std::map<std::string, Tank *>();
  std::map<std::string, double> *formula = new std::map<std::string, double>();
  double *total = new double();

  if (file == NULL)
  {
    std::cout << "Error opening file" << std::endl;
    exit(1);
  }
  else
  {
    std::cout << "RUN" << std::endl;
    char line[256];
    while (fgets(line, sizeof(line), file))
    {
        // print line
        // std::cout << line << std::endl;
        readline(line,*tanks_map,*formula,*total);
      
    }
    std::cout << "END RUN" << std::endl;
  }

  // tanks = std::vector<Tank>(tanks_map->size());
  for (auto it = tanks_map->begin(); it != tanks_map->end(); ++it)
  {
    tanks.push_back(*it->second);
  }

  // formulas = std::vector<Formula>(formula->size());
  for (auto it = formula->begin(); it != formula->end(); ++it)
  {
    Formula f;
    f.name = it->first;
    f.output = it->second;
    formulas.push_back(f);
  }



  int n = formulas.size();
  for (int i = 0; i < n; i++)
  {
    Formula f;
    f.name = formulas[i].name;
    f.quantity = formulas[i].quantity;
    // int m = tanks.size(); this is false but since wtf then let's assume it is
    int m = tanks.size();
    for (int j = 0; j < m; j++)
    {
      std::string input = tanks[j].formula.begin()->first;
      double input_quantity = tanks[j].formula.begin()->second;
      // std::cin >> input;
      //! WTF std::cin >> input_quantity;
      f.inputs.push_back(input);
      f.inputs_quantity.push_back(input_quantity);
    }
    f.output_quantity = formulas[i].output_quantity;
    f.output = formulas[i].output;
    f.is_solved = false;
    formulas.push_back(f);
  }
  int t = tanks.size();
  for (int i = 0; i < t; i++)
  {
    Tank tank(0,"/");
    tank.capacity = tanks[i].capacity;
    tank.formula["/"] = 100;
    tanks.push_back(tank);
  }
  double a = *total; // ????????
  Solve(formulas[0], a);
  for (int i = 0; i < tanks.size(); i++)
  {
    // std::cout << tanks[i].capacity << " " << tanks[i].formula.begin()->first << " " << tanks[i].formula.begin()->second << std::endl; formatted to get lisible strings and numbers formated to 2 decimals
    std::cout << std::fixed << std::setprecision(2) << tanks[i].capacity << " " << tanks[i].formula.begin()->first << " " << tanks[i].formula.begin()->second << std::endl;
  }

  return 0;
}

