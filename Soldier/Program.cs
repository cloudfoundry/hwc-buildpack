using System;
using System.Collections.Generic;
using System.Data;
using System.Diagnostics;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;

namespace Soldier
{
    class Program
    {
        static void Main(string[] args)
        {
            var thingToRun = args[0];
            var workingDir = args[1];
            var command = thingToRun.Split(' ');
            ProcessStartInfo startInfo;
            if (command.Length > 1)
            {
                startInfo = new ProcessStartInfo(command[0], command[1]);
            }
            else
            {
                startInfo = new ProcessStartInfo(command[0]);
            }
            startInfo.WorkingDirectory = workingDir;
            var process = new Process {StartInfo = startInfo};
            process.Start();
            process.WaitForExit();
        }
    }
}
