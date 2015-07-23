using System;
using System.Net.Http;
using System.Net.NetworkInformation;
using System.Net.Sockets;


namespace Healthcheck
{
    internal class Program
    {
        private static void Main(string[] args)
        {
            // The health check succeeds if the process is listening on any non-local interface
            {
                try
                {
                    var client = new HttpClient();
                    var port = Environment.GetEnvironmentVariable("PORT");
                    if (port == null)
                        throw new Exception("PORT is not defined");

                    foreach (NetworkInterface netInterface in NetworkInterface.GetAllNetworkInterfaces())
                    {
                        IPInterfaceProperties ipProps = netInterface.GetIPProperties();
                        foreach (UnicastIPAddressInformation addr in ipProps.UnicastAddresses)
                        {
                            if (addr.Address.AddressFamily != AddressFamily.InterNetwork) continue;
                            if (addr.Address.ToString().StartsWith("127.")) continue;
                            var task = client.GetAsync(String.Format("http://{0}:{1}", addr.Address.ToString(), port));
                            if (task.Wait(1000))
                            {
                                if (task.Result.IsSuccessStatusCode)
                                {
                                    Console.WriteLine("healthcheck passed");
                                    Environment.Exit(0);
                                }
                            }
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