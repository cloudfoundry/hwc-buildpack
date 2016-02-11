using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.NetworkInformation;
using System.Net.Sockets;
using System.Threading;
using System.Web.Script.Serialization;


namespace Healthcheck
{
    internal class Program
    {
        private static void Main(string[] args)
        {
            var client = new HttpClient();
            var instancePorts = Environment.GetEnvironmentVariable("CF_INSTANCE_PORTS");
            if (instancePorts == null)
                throw new Exception("CF_INSTANCE_PORTS is not defined");

            var internalPort = args[1];
            var externalPort = getExternalPort(instancePorts, internalPort);
            if (externalPort == "")
            {
                Console.WriteLine("healthcheck failed, port mapping not found for " + internalPort + " in " +
                                  instancePorts);
                Environment.Exit(1);
            }

            try
            {
                var task =
                    client.GetAsync(String.Format("http://{0}:{1}", Environment.GetEnvironmentVariable("CF_INSTANCE_IP"), externalPort));
                if (task.Wait(1000))
                {
                    if (task.Result.IsSuccessStatusCode)
                    {
                        Console.WriteLine("healthcheck passed");
                        Environment.Exit(0);
                    }
                    else
                    {
                        Console.Error.WriteLine("Got error response: " +
                                          task.Result.Content.ReadAsStringAsync().Result);
                    }
                }
                else
                {
                    Console.WriteLine("waiting for process to start up");
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }

            Console.WriteLine("healthcheck failed");

            Environment.Exit(1);
        }

        private static string getExternalPort(string jsonInstancePorts, string internalPort)
        {
            var serializer = new JavaScriptSerializer();
            var instancePorts = serializer.Deserialize<List<Dictionary<string, string>>>(jsonInstancePorts);
            var match = instancePorts.FirstOrDefault(x => x["internal"] == internalPort);
            if (match == null)
            {
                return "";
            }
            return match["external"];
        }
    }
}