using System.Text;
using NSpec;
using System;
using System.Diagnostics;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Threading;

namespace Healthcheck.Tests.Specs
{
    class HealthcheckSpecs : nspec
    {
        public void describe_()
        {
            int port = -1;
            before = () => port = new Random().Next(10000, 50000);
            Process process = null;

            act = () =>
            {
                var workingDir = Path.GetFullPath(Path.Combine(System.Reflection.Assembly.GetExecutingAssembly().CodeBase, "..", "..", "..", "..", "Healthcheck", "bin").Replace("file:///", ""));
                process = new Process
                {
                    StartInfo =
                    {
                        FileName = Path.Combine(workingDir, "Healthcheck.exe"),
                        WorkingDirectory = workingDir,
                        RedirectStandardOutput = true,
                        UseShellExecute = false
                    }
                };
                process.StartInfo.EnvironmentVariables["PORT"] = port.ToString();
                process.Start();
                process.WaitForExit();
            };

            describe["when the address is listening"] = () =>
            {
                HttpListener httpListener = null;
                before = () => httpListener = startServer("*", port);
                after = () => httpListener.Stop();

                it["exits 0 and logs it succeeded"] = () =>
                {
                    process.StandardOutput.ReadToEnd().should_be("healthcheck passed\r\n");
                    process.ExitCode.should_be(0);
                };
            };

            describe["when the address is listening only on localhost"] = () =>
            {
                HttpListener httpListener = null;
                before = () => httpListener = startServer("127.0.0.1", port);
                after = () => httpListener.Stop();

                it["exits 1 and logs it failed"] = () =>
                {
                    process.StandardOutput.ReadToEnd().should_be("healthcheck failed\r\n");
                    process.ExitCode.should_be(1);
                };
            };

            describe["when the address is not listening"] = () =>
            {
                it["exits 1 and logs it failed"] = () =>
                {
                    process.StandardOutput.ReadToEnd().should_contain("healthcheck failed\r\n");
                    process.ExitCode.should_be(1);
                };
            };
        }

        private HttpListener startServer(string host, int port)
        {
            var listener = new HttpListener();
            listener.Prefixes.Add(String.Format("http://{0}:{1}/", host, port));
            listener.Start();
            var listenThread = new Thread(new ThreadStart(() =>
            {
                try
                {
                    for (;;)
                    {
                        var httpContext = listener.GetContext();
                        httpContext.Response.StatusCode = 200;
                        var resp = UTF8Encoding.UTF8.GetBytes("Hello!");
                        httpContext.Response.OutputStream.Write(resp, 0, resp.Length);
                        httpContext.Response.OutputStream.Close();
                    }
                }
                catch (Exception e)
                {
                    // ignore the exception and exit
                }
            }));
            listenThread.Start();
            return listener;
        }
    }
}