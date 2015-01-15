using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WebAppServer
{
    public class Logger
    {
        public virtual void Info(string messageFormat, params object[] args)
        {
            Console.Out.WriteLine(DateTimeOffset.UtcNow.ToString("u") + "|INFO|" + messageFormat, args);
        }

        public virtual void Error(string messageFormat, params object[] args)
        {
            Console.Error.WriteLine(DateTimeOffset.UtcNow.ToString("u") + "|ERROR|" + messageFormat, args);
        }

        public virtual void ErrorException(string message, Exception exception)
        {
            Error(message);
            Error(exception.ToString());
        }
    }
}
