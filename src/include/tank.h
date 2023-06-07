#pragma once
#include <map>
#include <string>


struct Tank
{
  std::string ID;
  double capacity;
  std::string content;


  Tank(std::string Id, double size, std::string wine)
  {
    ID = Id;
    capacity = size;
    content = wine;
  }
};
