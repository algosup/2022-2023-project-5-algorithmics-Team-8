using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

namespace Algorithm
{
    public class Program
    {
        public static void Main(string[] args)
        {
            // Read the file
            List<string> listFormula = Parser.ReadCSV("./data/formula.csv");
            List<string> listTank = Parser.ReadCSV("./data/tank_configuration.csv");
            Console.WriteLine("Formula:");
            foreach (string line in listFormula)
            {
                Console.WriteLine(line);
            }
            Console.WriteLine("Tank:");
            foreach (string line in listTank)
            {
                Console.WriteLine(line);
                Parser.WriteFile("./data/output.txt", listTank.ToArray());
            }
        }
    }
}