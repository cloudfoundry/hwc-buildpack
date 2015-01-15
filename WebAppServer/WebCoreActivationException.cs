using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace WebAppServer
{
    using System.ComponentModel;

    public class WebCoreActivationException : Exception
    {
        // ref: http://blogs.iis.net/ksingla/archive/2007/12/20/ins-amp-outs-of-hostable-web-core.aspx
        public static readonly Dictionary<string, string> HResults = new Dictionary<string, string>
        {
            {"80070032", "HWC Activation failed - the request is not supported, check application event log for further details."},
            {"80070038", "HWC Activation failed - appHostConfig has multiple app pools defined in it"},
            {"800700B7", "HWC Activation failed - invalid binding: port/IP in use already or app pool name in use"},
            {"80070015", "HWC Activation failed - no websites defined in appHostConfig"},
            {"80070044", "HWC Activation failed - app pool defined but site pointing to nonexistend app pool"},
            {"8007000D", "HWC Activation failed - All configuration system checks are still valid. So invalid xml in configuration or duplicate site ids in appHostConfig."},
            {"80070057", "HWC Activation failed - bad parameter"},
            {"80070490", "HWC Activation failed - appHostConfig path is not valid"},
            {"80070420", "HWC Activation failed - service already running"}
        };

        public WebCoreActivationException(int result) : this(result.ToString("X"), result) { }

        private WebCoreActivationException(string hResult, int result) : base(
            HResults.ContainsKey(hResult) 
            ? new Win32Exception(result).Message + ": " + HResults[hResult] 
            : new Win32Exception(result).Message)
        {
            HResult = result;
        }
    }

    public class WebCorePortException : WebCoreActivationException
    {
        public WebCorePortException() : base(unchecked((int)WebServer.HostableWebCore.ERROR_ACCESS_DENIED))
        {
        }
    }
}
