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
            int port = new Random().Next(10000, 50000);
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
                TcpListener tcpListener = null;
                before = () => tcpListener = startServer(IPAddress.Any, port);
                after = () => tcpListener.Stop();

                it["exits 0 and logs it succeeded"] = () =>
                {
                    process.StandardOutput.ReadToEnd().should_be("healthcheck passed\r\n");
                    process.ExitCode.should_be(0);
                };
            };

            describe["when the address is listening only on localhost"] = () =>
            {
                TcpListener tcpListener = null;
                before = () => tcpListener = startServer(IPAddress.Parse("127.0.0.1"), port);
                after = () => tcpListener.Stop();

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
                    process.StandardOutput.ReadToEnd().should_be("healthcheck failed\r\n");
                    process.ExitCode.should_be(1);
                };
            };
        }

        private TcpListener startServer(IPAddress host, int port)
        {
            var tcpListener = new TcpListener(host, port);
            var listenThread = new Thread(new ThreadStart(() =>
            {
                try
                {
                    tcpListener.Start();
                    tcpListener.AcceptTcpClient();
                }
                catch (Exception e) { }
            }));
            listenThread.Start();
            return tcpListener;
        }
    }
}