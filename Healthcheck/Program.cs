using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Healthcheck
{
    class Program
    {
        static void Main(string[] args)
        {
            System.Console.WriteLine("If I was to check a port, it would be: {0}", Environment.GetEnvironmentVariable("PORT"));
            System.Console.WriteLine("Hi I am dummy Healthcheck, please replace.");
        }
    }
}
