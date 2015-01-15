using System;
using System.Diagnostics;
using System.IO;
using System.Threading;

namespace WebAppServer
{
    internal static class Program
    {
        private static readonly Logger log = new Logger();
        private static readonly ManualResetEvent exitLatch = new ManualResetEvent(false);
        private static readonly FileSystemWatcher fileSystemWatcher;

        static Program()
        {
            string workingDirectory =  Directory.GetCurrentDirectory();
            fileSystemWatcher = new FileSystemWatcher(workingDirectory);
            fileSystemWatcher.Created += fileSystemWatcher_Created;
            fileSystemWatcher.EnableRaisingEvents = true;
        }

        private static void fileSystemWatcher_Created(object sender, FileSystemEventArgs e)
        {
            if (e.ChangeType == WatcherChangeTypes.Created)
            {
                string lcName = e.Name.ToLowerInvariant().Trim();
                if (lcName == "iishost_stop")
                {
                    exitLatch.Set();
                }
            }
        }

        private static void Main(string[] args)
        {
            Debugger.Launch();

            Console.CancelKeyPress += (s, e) =>
            {
                e.Cancel = true;
                exitLatch.Set();
            };

            try
            {
                var port = uint.Parse(args[0]);
                var webRoot = args[1];

                log.Info("Port:{0}", port);
                log.Info("Webroot:{0}", webRoot);

                var configGenerator = new ConfigGenerator(webRoot);
                var settings = configGenerator.Create(
                    port,
                    Constants.FrameworkPaths.FourDotZeroWebConfig,
                    Constants.RuntimeVersion.VersionFourDotZero,
                    Constants.PipelineMode.Integrated, 
                    null, 
                    null);

                log.Info("starting web server instance...");
                using (var webServer = new WebServer(settings))
                {
                    webServer.Start();
                    Console.WriteLine("Server Started.... press CTRL + C to stop");

                    exitLatch.WaitOne();
                    Console.WriteLine("Server shutting down, please wait...");
                    webServer.Stop();
                }

                if (File.Exists("iishost_stop"))
                {
                    File.Delete("iishost_stop");
                }
            }
            catch (Exception ex)
            {
                log.ErrorException("Error on startup.", ex);
                Environment.Exit(1);
            }
        }
    }
}
