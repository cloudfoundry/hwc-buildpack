using System;
using System.Net.Http;


namespace Healthcheck
{
    class Program
    {
        static void Main(string[] args)
        {
            // The health check succeeds if the process is listening on any non-local interface
            {
                try
                {
                    var client = new HttpClient();
                    var port = Environment.GetEnvironmentVariable("PORT");
                    if (port == null)
                        throw new Exception("PORT is not defined");

                    var task = client.GetAsync("http://127.0.0.1:" + port);
                    if (task.Wait(1000))
                    {
                        if (task.Result.IsSuccessStatusCode)
                        {
                            Console.WriteLine("healthcheck passed");
                            Environment.Exit(0);
                        }
                    }
                }
                catch (Exception e)
                {
                    Console.WriteLine(e.ToString());
                }

                Console.WriteLine("healthcheck failed");

                Environment.Exit(1);
            }
        }
    }
}
