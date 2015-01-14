using Microsoft.Web.Administration;
using System;
using System.Collections.Generic;
using System.Data;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices;
using System.Security;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Soldier
{
    class Program
    {
        static void Main(string[] args)
        {
            var containerRootPath = System.IO.Directory.GetCurrentDirectory();
            var containerID = new DirectoryInfo(containerRootPath).Name;
            var webDeployPath = Path.Combine(containerRootPath, "webdeploy");
            var zipFileLocation =
                Directory.GetFiles(webDeployPath, "*.zip", SearchOption.TopDirectoryOnly).SingleOrDefault();

            //CRAZY MSDEPLOY COMMAND LINE.  THIS SHOULD BE MADE BETTERER.
            var deployCommand = "\"C:\\Program Files\\IIS\\Microsoft Web Deploy V3\\msdeploy.exe\"";         
            var deployCommandArgs = "-verb:sync -source:package="+zipFileLocation+" -dest:auto -setParam:name='IIS Web Application Name',value=\"" + containerID + "\" -presync:runCommand='\"C:\\Windows\\System32\\inetsrv\\appcmd set site " + containerID + " /bindings:http/*:8080:'";

            ProcessStartInfo startInfo;
            startInfo = new ProcessStartInfo(deployCommand, deployCommandArgs);
            startInfo.WorkingDirectory = webDeployPath;
            var process = new Process {StartInfo = startInfo};
            process.Start();
            process.WaitForExit();

            var stockWebConfig = Path.Combine(containerRootPath, "tmp", "circus", "Web.config");
            var webConfigDestination = Path.Combine(containerRootPath, "Web.config");
            File.Copy(stockWebConfig, webConfigDestination);
            ServerManager serverManager = ServerManager.OpenRemote("localhost");
            Site site = serverManager.Sites[containerID];
            site.Start();
            Thread.Sleep(20000);
        }
    }
}
