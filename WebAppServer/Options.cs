using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
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
            Port = UInt32.Parse(args[0]);
            WebRoot = Path.GetFullPath(args[1]);
        }
    }
}
