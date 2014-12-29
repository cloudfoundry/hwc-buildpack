using Microsoft.Web.Administration;
using System;
using System.Collections.Generic;
using System.Data;
using System.Diagnostics;
using System.IO;
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
            var containerRootPath = System.IO.Directory.GetCurrentDirectory();
            var containerID = new DirectoryInfo(containerRootPath).Name;
            //HARDCODE LOCATION OF ZIPFILE
            var zipFileLocation = containerRootPath + "\\app\\webdeploy\\Nora.zip";

            //CRAZY MSDEPLOY COMMAND LINE.  THIS SHOULD BE MADE BETTERER.
            var deployCommand = "\"C:\\Program Files\\IIS\\Microsoft Web Deploy V3\\msdeploy.exe\" -verb:sync -source:package="+zipFileLocation+" -dest:auto -setParam:name='IIS Web Application Name',value=\"" + containerID + "\" -presync:runCommand='\"C:\\Windows\\System32\\inetsrv\\appcmd set site " + containerID + " /bindings:http/*:8080:'";

            ProcessStartInfo startInfo;
            
            startInfo = new ProcessStartInfo(deployCommand);
            
            startInfo.WorkingDirectory = containerRootPath;
            var process = new Process {StartInfo = startInfo};
            process.Start();
            process.WaitForExit();

            ServerManager serverManager = ServerManager.OpenRemote("localhost");
            Site site = serverManager.Sites[containerID];
            site.Start();
        }
    }
}
