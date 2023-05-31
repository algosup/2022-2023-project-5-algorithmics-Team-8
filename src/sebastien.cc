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

  FILE *file = fopen(filepath.c_str(), "r");

  // std::map<std::string, Tank *> *tanks = new std::map<std::string, Tank *>();
  // std::map<std::string, double> *formula = new std::map<std::string, double>();
  // double *total = new double();

  if (file == NULL)
  {
    std::cout << "Error opening file" << std::endl;
    exit(1);
  }
  else
  {
    char line[256];
    while (fgets(line, sizeof(line), file))
    {
      // print line
      std::cout << line << std::endl;
      // readline(line,*tanks,*formula,*total);
    }
  }

  Result result = {};
  return result;
}

// Comment => line starts with '!'
// Total => single value on line
// Formula => 2 values on line
// Tank => 3 values on line
void readline(std::string line, std::map<std::string, Tank *> &tanks, std::map<std::string, double> &formula, double &total)
{
  std::string separator = ";";
  size_t pos = 0;
  std::string token;
  std::vector<std::string> tokens;

  // if line does not contain separator check if it is a comment
  if (line[0] == '!' || line == "\n")
  {

    std::cout << line << std::endl;
    std::cout << "Comment or empty line: " << std::endl;
    return;
  }

  // while  not at eof
  printf("pos: %d\n", pos);
  token = line.substr(0, line.length());
  printf("token: %s\n", token.c_str());
  std::vector<std::string> values;
  switch (token[0])
  {
  case 48 ... 57:
  {
    total = std::stod(token);
    break;
  }
  case 65 ... 122:
  {
    formula[token.substr(0, token.find(separator))] = std::stod(token.substr(token.find(separator) + 1, token.length()));
    // formula[tokens[0]] = std::stod(tokens[1].substr(0, tokens[1].length() - 1));
    break;
  }
  case '#':
  {
    // std::cout << "Comment: " << token << std::endl;
    // std::cout << std::stod(token.substr(token.find(separator) + 1, token.find_last_of(separator) + 1)) << std::endl;
    // std::cout << token.substr(token.find_last_of(separator) + 1, token.length()) << std::endl;

    Tank *tank = new Tank(std::stod(token.substr(token.find(separator) + 1, token.length())), token.substr(token.find_last_of(separator) + 1, token.length()));
    tanks[token.substr(1, token.find(separator))] = tank;
    break;
  }
  default:
    throw;
    break;
  }
  line.erase(0, pos + separator.length());
}
