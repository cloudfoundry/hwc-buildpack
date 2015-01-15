using System;
using System.IO;
using System.Runtime.InteropServices;

namespace WebAppServer
{
    public interface IWebServer : IDisposable
    {
        void Start();
        void Stop();
    }

    // Ref: http://msdn.microsoft.com/en-us/library/ms689327%28v=vs.90%29.aspx
    // May require: // netsh http add urlacl url=http://*:PORT/ user=DOMAIN\user
    public class WebServer : IWebServer
    {
        private readonly ConfigSettings configSettings;

        public WebServer(ConfigSettings configSettings)
        {
            this.configSettings = configSettings;
        }

        public void Dispose()
        {
            Stop();
        }

        public void Start()
        {
            if (!HostableWebCore.IsActivated)
            {
                HostableWebCore.Activate(configSettings.AppConfigPath, configSettings.RootWebConfigPath, Guid.NewGuid().ToString());
            }
        }

        public void Stop()
        {
            if (HostableWebCore.IsActivated)
            {
                HostableWebCore.Shutdown(false);
            }
        }

        #region Hostable WebCore
        public static class HostableWebCore
        {
            public const uint ERROR_ACCESS_DENIED = 0x80070005;

            const string HWCPath = @"%windir%\system32\inetsrv\hwebcore.dll";

            private static bool _isActivated;

            private delegate uint FnWebCoreShutdown(bool immediate);
            private delegate uint FnWebCoreActivate(
                [In, MarshalAs(UnmanagedType.LPWStr)]string appHostConfig,
                [In, MarshalAs(UnmanagedType.LPWStr)]string rootWebConfig,
                [In, MarshalAs(UnmanagedType.LPWStr)]string instanceName);

            private static readonly FnWebCoreActivate WebCoreActivate;
            private static readonly FnWebCoreShutdown WebCoreShutdown;

            static HostableWebCore()
            {
                var hostableWebCorePath = Environment.ExpandEnvironmentVariables(HWCPath);
                if (!File.Exists(hostableWebCorePath))
                {
                    throw new FileNotFoundException("Unable to locate hostable web core library, ensure IIS 7+ is installed", hostableWebCorePath);
                }

                // Load the library and get the function pointers for the WebCore entry points
                IntPtr hwc = HWCInterop.LoadLibrary(hostableWebCorePath);

                IntPtr procaddr = HWCInterop.GetProcAddress(hwc, "WebCoreActivate");
                WebCoreActivate = (FnWebCoreActivate) Marshal.GetDelegateForFunctionPointer(procaddr, typeof(FnWebCoreActivate));

                procaddr = HWCInterop.GetProcAddress(hwc, "WebCoreShutdown");
                WebCoreShutdown = (FnWebCoreShutdown) Marshal.GetDelegateForFunctionPointer(procaddr, typeof(FnWebCoreShutdown));
            }

            /// <summary>
            /// Specifies if Hostable WebCore has been activated
            /// </summary>
            public static bool IsActivated
            {
                get { return _isActivated; }
            }

            /// <summary>
            /// Activate the HWC
            /// </summary>
            /// <param name="appHostConfig">Path to ApplicationHost.config to use</param>
            /// <param name="rootWebConfig">Path to the Root Web.config to use</param>
            /// <param name="instanceName">Name for this instance</param>
            public static void Activate(string appHostConfig, string rootWebConfig, string instanceName)
            {
                uint result = WebCoreActivate(appHostConfig, rootWebConfig, instanceName);
                if (result != 0)
                {
                    if (result == ERROR_ACCESS_DENIED)
                        throw new WebCorePortException();

                    throw new WebCoreActivationException((int)result);
                }

                _isActivated = true;
            }

            /// <summary>
            /// Shutdown HWC
            /// </summary>
            public static void Shutdown(bool immediate)
            {
                if (_isActivated)
                {
                    WebCoreShutdown(immediate);
                    _isActivated = false;
                }
            }

            private static class HWCInterop
            {
                [DllImport("kernel32.dll", SetLastError = true)]
                internal static extern IntPtr LoadLibrary(String dllname);

                [DllImport("kernel32.dll", SetLastError = true)]
                internal static extern IntPtr GetProcAddress(IntPtr hModule, String procname);
            }
        }
        #endregion
    }
}
