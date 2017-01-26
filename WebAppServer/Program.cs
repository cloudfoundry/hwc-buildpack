using System;
using System.Net;
using System.Security.Principal;
using System.Threading;

namespace WebAppServer
{
    public static class Program
    {
        private static readonly Logger log = new Logger();
        private static readonly ManualResetEvent exitLatch = new ManualResetEvent(false);

        public static int RunWebServer(IWebServer webServer, Logger log, ManualResetEvent exitLatch)
        {
            try
            {
                log.Info("Starting web server instance...");
                webServer.Start();
                Console.WriteLine("Server Started.... press CTRL + C to stop");
                Warmup(webServer);
                exitLatch.WaitOne();
                Console.WriteLine("Server shutting down, please wait...");
                webServer.Stop();

                return 0;
            }
            catch (WebCorePortException)
            {
                log.Error("Please allow the user to access the port. eg. 'netsh http add urlacl url=http://*:9999/ user={0}'", WindowsIdentity.GetCurrent().Name);
                return 1;
            }
        }

        private static void Warmup(IWebServer server)
        {
            var req = WebRequest.Create("http://localhost:" + server.Port);
            req.Timeout = 1000;
            try
            {
                req.GetResponse();
            }
            catch (WebException ex)
            {
                // ignore exceptions. An exception can be thrown if a 404 was returned
                // which is a possibility if the application doesn't have `~/' endpoint
            }
        }

        private static int Main(string[] args)
        {
            Console.CancelKeyPress += (s, e) =>
            {
                e.Cancel = true;
                exitLatch.Set();
            };

            var exitCode = 0;
            try
            {
                var options = new Options();
                options.Parse(args);

                log.Info("Port:{0}", options.Port);
                log.Info("Webroot:{0}", options.WebRoot);

                var configGenerator = new ConfigGenerator(options.WebRoot);
                var webConfig = WebConfig.Create(Environment.ExpandEnvironmentVariables(Constants.FrameworkPaths.FourDotZeroWebConfig),
                    AppDomain.CurrentDomain.BaseDirectory);
                var settings = configGenerator.Create(
                    options.Port,
                    webConfig,
                    Constants.RuntimeVersion.VersionFourDotZero,
                    Constants.PipelineMode.Integrated, 
                    null, 
                    null);

                using (var webServer = new WebServer(settings))
                {
                    exitCode = RunWebServer(webServer, log, exitLatch);
                }
            }
            catch (Exception ex)
            {
                log.ErrorException("Error on startup.", ex);
                exitCode = 1;
            }

            return exitCode;
        }
    }
}
