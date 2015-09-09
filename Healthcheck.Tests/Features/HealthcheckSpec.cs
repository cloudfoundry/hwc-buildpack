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
            before = () => port = GetFreeTcpPort();
            Process process = null;
            string processOutputData = null;
            string processErrorData = null;

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
                        RedirectStandardError = true,
                        UseShellExecute = false
                    }
                };

                process.StartInfo.EnvironmentVariables["PORT"] = port.ToString();

                process.Start();
                processOutputData = process.StandardOutput.ReadToEnd();
                processErrorData = process.StandardError.ReadToEnd();
                process.WaitForExit();
            };

            describe["when the server is returning non success status code"] = () =>
            {
                HttpListener httpListener = null;
                var stacktrace = "BOOOOOOM";
                before = () => httpListener = startServer("*", port, 500, stacktrace);
                after = () => httpListener.Stop();

                it["exits 1 and logs the stack trace"] = () =>
                {
                    processOutputData.should_contain("healthcheck failed\r\n");
                    processErrorData.should_contain(stacktrace);
                    process.ExitCode.should_be(1);
                };
            };

            describe["when the address is listening"] = () =>
            {
                HttpListener httpListener = null;
                before = () => httpListener = startServer("*", port);
                after = () => httpListener.Stop();

                it["exits 0 and logs it succeeded"] = () =>
                {
                    processOutputData.should_be("healthcheck passed\r\n");
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
                    processOutputData.should_contain("healthcheck failed\r\n");
                    process.ExitCode.should_be(1);
                };
            };

            describe["when the address is not listening"] = () =>
            {
                it["exits 1 and logs it failed"] = () =>
                {
                    processOutputData.should_contain("healthcheck failed\r\n");
                    process.ExitCode.should_be(1);
                };
            };
        }

        private int GetFreeTcpPort()
        {
            var tcpl = new TcpListener(IPAddress.Any, 0);
            tcpl.Start();

            var freePort = (tcpl.LocalEndpoint as IPEndPoint).Port;
            tcpl.Stop();

            return freePort;
        }

        private HttpListener startServer(string host, int port, int statusCode = 200, string content = "Hello!")
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
                        httpContext.Response.StatusCode = statusCode;
                        var resp = UTF8Encoding.UTF8.GetBytes(content);
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