using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.NetworkInformation;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

namespace Healthcheck
{
    class Program
    {
        static void Main(string[] args)
        {
            // The health check succeeds if the process is listening on any non-local interface
            foreach (NetworkInterface netInterface in NetworkInterface.GetAllNetworkInterfaces())
            {
                IPInterfaceProperties ipProps = netInterface.GetIPProperties();
                foreach (UnicastIPAddressInformation addr in ipProps.UnicastAddresses)
                {
                    if (addr.Address.AddressFamily != AddressFamily.InterNetwork) continue;
                    if (addr.Address.ToString().StartsWith("127.")) continue;

                    try
                    {
                        var tcpCLient = new TcpClient(addr.Address.ToString(), Int32.Parse(Environment.GetEnvironmentVariable("PORT")));

                        System.Console.WriteLine("healthcheck passed");
                        System.Environment.Exit(0);
                    }
                    catch (Exception)
                    {
                    }
                }
            }

            System.Console.WriteLine("healthcheck failed");
            System.Environment.Exit(1);
        }
    }
}
