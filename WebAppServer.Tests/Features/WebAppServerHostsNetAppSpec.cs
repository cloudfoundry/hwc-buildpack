using NSpec;
using System;
using System.Diagnostics;
using System.IO;
using System.Net;
using System.Net.Http;
using System.Security.Principal;
using System.Threading;

namespace WebAppServer.Tests
{
    class WebAppServerHostsNetAppSpec : nspec
    {
        private void describe_()
        {
            describe["Given that I have a ASP.NET MVC application"] = () =>
            {
                describe["When I pass it to WebAppServer"] = () =>
                {
                    int port = 3300;
                    Process process = null;

                    before = () => port = 3300;
                    
                    act = () =>
                    {
                        var workingDir = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "WebAppServer.Tests");
                        var webRoot = Path.Combine(workingDir, "Fixtures", "Nora");

                        process = new Process
                        {
                            StartInfo =
                            {
                                FileName = Path.Combine(workingDir, "bin", "debug", "WebAppServer.exe"),
                                Arguments = String.Format("{0} \"{1}\"", port, webRoot),
                                WorkingDirectory = workingDir,
                                RedirectStandardInput = true,
                                RedirectStandardError = true,
                                UseShellExecute = false
                            }
                        };

                        process.Start();                    
                    };

                    it["runs it on the specified port"] = () =>
                    {
                        var client = new HttpClient();
                        var response = client.GetAsync("http://localhost:" + port).GetAwaiter().GetResult();
                        response.StatusCode.should_be(HttpStatusCode.OK);
                        response.Content.ReadAsStringAsync().Result.should_be("\"hello i am nora\"");
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
                };
            };
        }
    }
}
