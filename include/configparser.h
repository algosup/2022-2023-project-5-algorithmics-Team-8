#pragma once
#include <string>
#include <vector>
#include "tank.h"
using namespace std;

class configparser
{
  vector<Tank> tanks;
  configparser(string path);
};
