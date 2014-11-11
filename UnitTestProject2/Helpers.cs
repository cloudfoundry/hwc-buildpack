using System.Diagnostics;
using System.IO;

namespace Tailor.Tests
{
    public static class Helpers 
    {
        public static Process StartTailor(string appDir, string outputDroplet, string outputMetadata)
        {

            var process = new Process();
            process.StartInfo.FileName = @"..\..\..\Tailor\bin\Debug\Tailor.exe";
            process.StartInfo.Arguments = "--appDir=\"" + appDir + "\" --outputDroplet=\"" + outputDroplet + "\" --outputMetadata=\"" + outputMetadata + "\"";
            process.StartInfo.UseShellExecute = false;
            process.StartInfo.RedirectStandardOutput = true;
            process.Start();

            return process;
        }
    }
}
