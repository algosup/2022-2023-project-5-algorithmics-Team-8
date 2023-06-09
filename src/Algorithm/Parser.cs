using System;
using System.Collections.Generic;
using System.IO;

///<summary>
///Parser class, used to parse CSV files
///</summary>
public class Parser
{
    ///<summary>
    ///ReadCSV reads a CSV file and returns a list of strings separated by commas
    ///<param name="path">The path to the CSV file</param>
    ///<returns>A list of strings separated by commas</returns>
    ///</summary>

    public static List<string> ReadCSV(string path)
    {
        List<string> list = new List<string>();
        try
        {
            using (var reader = new StreamReader(path))
            {
                while (!reader.EndOfStream)
                {
                    var line = reader.ReadLine();
                    list.Add(line);
                    string[] lineArray = line.Split(',');
                }
            }
            return list;
        }
        catch (FileNotFoundException)
        {
            Console.WriteLine("File not found");
            return list;
        }
    }
    public static void WriteFile(string path, string[] lines)
    {
        try
        {
            using (var writer = new StreamWriter(path))
            {
                foreach (string line in lines)
                {
                    writer.WriteLine(line);
                }
            }
        }
        catch (FileNotFoundException)
        {
            Console.WriteLine("File not found");
        }
    }
}
