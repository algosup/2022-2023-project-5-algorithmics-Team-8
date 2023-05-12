#include "include/sebastien.h"
#include "include/tank.h"
#include <vector>
#include <string>



/**
 * reads the content of a file into a Result struct
 * @param filepath Relative path of the file to 
 */

Result configparser(std::string filepath)
{
  std::string line;
  std::ifstream file;
  file.open(filepath);
  
  std::map<std::string,Tank*> *tanks = new std::map<std::string,Tank*>();
  std::map<std::string,double> *formula = new std::map<std::string,double>();
  double *total = new double();
  
  while(getline(file,line))
  {
       readline(line,*tanks,*formula,*total);
  }
  
  Result result = {*tanks,*formula};
  return result;
}


// Comment => line starts with '!'
// Total => single value on line
// Formula => 2 values on line
// Tank => 3 values on line
void readline(std::string line, std::map<std::string,Tank*>& tanks, std::map<std::string,double>& formula,double& total)
{
  std::string separator = ";";
  size_t pos = 0;
  std::string token;
  std::vector<std::string> tokens;
  while ((pos = line.find(separator)) != std::string::npos)
  {
    token = line.substr(0,pos);
    std::vector<std::string> values;
    if(token == "!") return;
  }
  switch (token.length())
  {
      case 1:
      {
        total = std::stod(tokens[0]);
        break;
      }
      case 2:
      {
        formula[tokens[0]] =std::stod(tokens[1].substr(0,tokens[1].length()-1));
        break;
      }
      case 3:
      {
        Tank *tank = new Tank(std::stod(tokens[1]),tokens[2] == "/" ? "" : tokens[2]);
        tanks[tokens[0]] = tank;
        break;
      }
      default:
        throw;
        break;
  }
    
  
}
