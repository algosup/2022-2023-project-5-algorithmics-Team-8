#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <algorithm>
#include <iterator>
#include <map>
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

int main()
{
  int n;
  std::cin >> n;
  for (int i = 0; i < n; i++)
  {
    Formula f;
    std::cin >> f.name;
    std::cin >> f.quantity;
    int m;
    std::cin >> m;
    for (int j = 0; j < m; j++)
    {
      std::string input;
      double input_quantity;
      std::cin >> input;
      std::cin >> input_quantity;
      f.inputs.push_back(input);
      f.inputs_quantity.push_back(input_quantity);
    }
    std::cin >> f.output_quantity;
    std::cin >> f.output;
    f.is_solved = false;
    formulas.push_back(f);
  }
  int t;
  std::cin >> t;
  for (int i = 0; i < t; i++)
  {
    Tank tank(0,"/");
    std::cin >> tank.capacity;
    tank.formula["/"] = 100;
    tanks.push_back(tank);
  }
  double a;
  std::cin >> a;
  Solve(formulas[0], a);
  for (int i = 0; i < tanks.size(); i++)
  {
    std::cout << tanks[i].capacity << " " << tanks[i].formula["/"] << std::endl;
  }
  return 0;
}
