using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GenerateFile
{
    class Program
    {
        static void Main(string[] args)
        {
            File.WriteAllText("dummyFile.txt", "My very excellent mother just sat under nine");
        }
    }
}
