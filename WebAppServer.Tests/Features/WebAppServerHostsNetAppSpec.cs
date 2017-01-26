using NSpec;
using System;
using System.Diagnostics;
using System.IO;
using System.Net;
using System.Net.Http;
using System.Security.Principal;
using System.Linq;

namespace WebAppServer.Tests
{
    class WebAppServerHostsNetAppSpec : nspec
    {
        const string processName = "WebAppServer.exe";
        const int port = 43311;
        private void describe_()
        {
            Process process = null;

            before = () =>
            {
                KillOrphanWebAppServer();
            };

            after = () =>
            {
                if (!process.HasExited)
                {
                    process.Kill();
                    process.WaitForExit();
                }

                if (process.ExitCode == 1)
                {
                    throw new Exception(
                        String.Format(
                            "WebAppServer failed to start. Check that the current user has access to the port (eg. run 'netsh http add urlacl url=http://*:{0}/ user={1}' as an Administrator).",
                            port,
                            WindowsIdentity.GetCurrent().Name));
                }
            };

            describe["Given that I have a ASP.NET MVC application"] = () =>
            {
                describe["When I pass it to WebAppServer"] = () =>
                {

                    act = () =>
                    {
                        process = StartApp("Nora");
                    };

                    it["runs it on the specified port"] = () =>
                    {
                        var client = new HttpClient();
                        var response = client.GetAsync("http://localhost:" + port).GetAwaiter().GetResult();
                        response.StatusCode.should_be(HttpStatusCode.OK);
                        response.Content.ReadAsStringAsync().Result.should_be("\"hello i am nora\"");
                    };

                    it["does not add unexpected custom headers"] = () =>
                    {
                        var client = new HttpClient();
                        var response = client.GetAsync("http://localhost:" + port).GetAwaiter().GetResult();
                        response.StatusCode.should_be(HttpStatusCode.OK);

                        var CustomHeaders = response.Headers.Where((x) => x.Key.StartsWith("X-")).Select((x) => x.Key).ToList();
                        CustomHeaders.should_be(new string[] { "X-AspNet-Version" });
                    };

                };
            };

            describe["Given that I have a ASP classic application"] = () =>
            {
                it["runs it on the specified port"] = () =>
                {
                    var client = new HttpClient();
                    var response = client.GetAsync("http://localhost:" + port).GetAwaiter().GetResult();
                    response.StatusCode.should_be(HttpStatusCode.OK);
                    response.Content.ReadAsStringAsync().Result.should_be("Hello World!");
                };

                act = () =>
                {
                    process = StartApp("ASP-classic");
                };
            };
        }

        private Process StartApp(string fixtureName)
        {
            Process process;
            var workingDir =
                Path.GetFullPath(
                    Path.Combine(System.Reflection.Assembly.GetExecutingAssembly().CodeBase, "..", "..", "..")
                        .Replace("file:///", ""));
            var webRoot = Path.Combine(workingDir, "Fixtures", fixtureName);

            process = new Process
            {
                StartInfo =
                {
                    FileName = Path.Combine(workingDir, "..", "WebAppServer", "bin", processName),
                    Arguments = ".",
                    WorkingDirectory = webRoot,
                    RedirectStandardInput = true,
                    RedirectStandardError = true,
                    RedirectStandardOutput = true,
                    UseShellExecute = false
                }
            };
            process.StartInfo.EnvironmentVariables["PORT"] = port.ToString();

            process.Start();

            // wait for the web app server to start to avoid race conditions
            while (!process.HasExited)
            {
                var readLine = process.StandardOutput.ReadLine();
                readLine.should_not_be_null();
                if (readLine.Contains("Server Started"))
                {
                    break;
                }
            }
            return process;
        }

        private void KillOrphanWebAppServer()
        {
            const string force = "/f";
            const string alsoKillChildProcesses = "/t";
            const string useProcessName = "/im";
            var killOrphanProcess = new Process
            {
                StartInfo =
                {
                    FileName = "taskkill.exe",
                    Arguments =
                        String.Format("{0} {1} {2} {3}", force, alsoKillChildProcesses, useProcessName,
                            processName),
                    UseShellExecute = false,
                    RedirectStandardError = true
                }
            };

            killOrphanProcess.Start();
            var stderr = killOrphanProcess.StandardError.ReadToEnd();
            if (stderr.Contains("Access is denied"))
            {
                throw new Exception(String.Format("Could not remove orphan {0}." +
                                                  "This is likely because WinowsAppLifecycle was previously run as Administrator," +
                                                  " and is currently running under a non privleged user. To fix:" +
                                                  " run \"taskkill.exe /f /t /im {0}\" in admin command prompt", processName));

            }
            killOrphanProcess.WaitForExit();
        }
    }
}

