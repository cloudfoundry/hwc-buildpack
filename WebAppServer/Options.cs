using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using System.Threading.Tasks;

namespace WebAppServer
{
    public class Options
    {
        public uint Port { get; set; }
        public string WebRoot { get; set; }

        public void Parse(string[] args)
        {
            Console.Out.WriteLine("PORT == {0}", Environment.GetEnvironmentVariable("PORT"));

            Port = uint.Parse(Environment.GetEnvironmentVariable("PORT"));
            WebRoot = Path.GetFullPath(".");
        }
    }
}
