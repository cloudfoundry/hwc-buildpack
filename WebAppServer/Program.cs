using System;
using System.Threading;

namespace WebAppServer
{
    internal static class Program
    {
        private static readonly Logger log = new Logger();
        private static readonly ManualResetEvent exitLatch = new ManualResetEvent(false);

        private static void Main(string[] args)
        {
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
            }
            catch (Exception ex)
            {
                log.ErrorException("Error on startup.", ex);
                Environment.Exit(1);
            }
        }
    }
}
