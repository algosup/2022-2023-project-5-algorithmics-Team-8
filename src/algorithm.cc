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


std::map<std::string, Tank *> *used_tanks = new std::map<std::string, Tank *>();

std::string Find_Tank(double amount)
{
  for (int i = 0; i < tanks.size(); i++)
  {
    //std::cout << tanks[i].content << " " << tanks[i].content.find("/") <<  " " <<  (tanks[i].content.find("/") != std::string::npos) << std::endl;
    if (tanks[i].capacity == amount && tanks[i].content.find("/") != std::string::npos)
    {
      used_tanks->insert(tanks[i].content, tanks[i].ID);
      return tanks[i].ID;
    }
  }
  return "No tank found";
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

  for (auto it = formula->begin(); it != formula->end(); ++it){
    it->second = it->second*(*total)/100;
  }

  // formulas = std::vector<Formula>(formula->size());
  for (auto it = formula->begin(); it != formula->end(); ++it)
  {
    Formula f;
    f.name = it->first;
    f.output = it->second;
    formulas.push_back(f);
    std::cout << it->second << Find_Tank(it->second) << std::endl;
  }
  

  return 0;
}

