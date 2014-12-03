using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MakeDirectoryForTheEnterprise
{
    class Program
    {
        static void Main(string[] args)
        {
            Directory.CreateDirectory(args[0]);
        }
    }
}
