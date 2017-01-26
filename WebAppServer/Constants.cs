namespace WebAppServer
{
    public class Constants
    {
        public class ConfigXPath
        {
            public static string ApplicationHost = "/configuration/system.applicationHost";
            public static string AppPools = ApplicationHost + "/applicationPools";
            public static string Sites = ApplicationHost + "/sites";
            public static string SiteDefaults = Sites + "/siteDefaults";

            public static string WebServer = "/configuration/system.webServer";
        }

        public class FrameworkPaths
        {
            public static string TwoDotZero = @"%windir%\Microsoft.NET\Framework\v2.0.50727";
            public static string FourDotZero = @"%windir%\Microsoft.NET\Framework\v4.0.30319";

            public static string TwoDotZeroWebConfig = TwoDotZero + @"\Config\web.config";
            public static string FourDotZeroWebConfig = FourDotZero + @"\Config\web.config";
        }

        public class RuntimeVersion
        {
            public static string VersionFourDotZero = "v4.0";
            public static string VersionTwoDotZero = "v2.0";
        }

        public class PipelineMode
        {
            public static string Integrated = "Integrated";
            public static string Classic = "Classic";
        }
    }
}
