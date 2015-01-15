using System;
using System.IO;
using System.Runtime.InteropServices;

namespace WebAppServer
{
    // Ref: http://msdn.microsoft.com/en-us/library/ms689327%28v=vs.90%29.aspx
    // May require: // netsh http add urlacl url=http://*:PORT/ user=DOMAIN\user
    internal class WebServer : IDisposable
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
        internal static class HostableWebCore
        {
            const string HWCPath = @"%windir%\system32\inetsrv\hwebcore.dll";

            private static bool _isActivated;

            private delegate int FnWebCoreShutdown(bool immediate);
            private delegate int FnWebCoreActivate(
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
                int result = WebCoreActivate(appHostConfig, rootWebConfig, instanceName);
                if (result != 0)
                {
                    throw new WebCoreActivationException(result);
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
