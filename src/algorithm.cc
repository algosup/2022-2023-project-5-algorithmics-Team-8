#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <algorithm>
#include <iterator>
#include <map>
#include <iomanip>
#include <cmath>
#include "include/algorithm.h"

using namespace std;

void Find_Tank(float amount)
{
    for (int i = 0; i < tanks.size(); i++)
    {
        if (tanks[i].quantity_left == amount)
        {
            tanks[i].is_solved = true;
            return;
        }
    }
}

vector<Formula> Split_Formula(Formula F)
{
    vector<Formula> formulas;
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

void Solve(Formula F, float a)
{
    Find_Tank(a);
    vector<Formula> formulas = Split_Formula(F);
    for (int i = 0; i < formulas.size(); i++)
    {
        Solve(formulas[i], formulas[i].quantity);
    }
}

int main()
{
    int n;
    cin >> n;
    for (int i = 0; i < n; i++)
    {
        Formula f;
        cin >> f.name;
        cin >> f.quantity;
        int m;
        cin >> m;
        for (int j = 0; j < m; j++)
        {
            string input;
            float input_quantity;
            cin >> input;
            cin >> input_quantity;
            f.inputs.push_back(input);
            f.inputs_quantity.push_back(input_quantity);
        }
        cin >> f.output_quantity;
        cin >> f.output;
        f.is_solved = false;
        formulas.push_back(f);
    }
    int t;
    cin >> t;
    for (int i = 0; i < t; i++)
    {
        Tank tank;
        cin >> tank.name;
        cin >> tank.quantity;
        tank.quantity_left = tank.quantity;
        tank.is_solved = false;
        tanks.push_back(tank);
    }
    float a;
    cin >> a;
    Solve(formulas[0], a);
    for (int i = 0; i < tanks.size(); i++)
    {
        cout << tanks[i].name << " " << tanks[i].quantity - tanks[i].quantity_left << endl;
    }
    return 0;
}
