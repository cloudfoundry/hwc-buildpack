using NSpec;
using System;
using System.Diagnostics;
using System.IO;
using System.Net;
using System.Net.Http;

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
                    const int port = 3300;
                    Process process = null;

                    before = () =>
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
                    };

                    after = () =>
                    {
                        process.Kill();
                        process.WaitForExit();
                    };
                };
            };
        }
    }
}
