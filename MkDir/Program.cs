using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MkDir
{
    class Program
    {
        static void Main(string[] args)
        {
            Directory.CreateDirectory(args[0]);
        }
    }
}
