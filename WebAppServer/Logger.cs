using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WebAppServer
{
    class Logger
    {
        public void Info(string messageFormat, params object[] args)
        {
            Console.Out.WriteLine(DateTimeOffset.UtcNow.ToString("u") + "|INFO|" + messageFormat, args);
        }

        public void ErrorException(string message, Exception exception)
        {
            Console.Error.WriteLine(DateTimeOffset.UtcNow.ToString("u") + "|ERROR|" + message);
            Console.Error.WriteLine(DateTimeOffset.UtcNow.ToString("u") + "|ERROR|" + exception.ToString());
        }
    }
}
