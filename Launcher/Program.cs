using System;
using System.Diagnostics;
using System.IO;
using System.Runtime.InteropServices;
using System.Threading;
using Newtonsoft.Json;

namespace Launcher
{
    public class ExecutionMetadata
    {
        [JsonProperty("start_command")]
        public string StartCommand { get; set; }

        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs { get; set; }
    }

    internal class Program
    {
        private static int Main(string[] args)
        {
            if (Environment.GetEnvironmentVariable("ARGJSON") != null && Environment.GetEnvironmentVariable("ARGJSON").Length >= 2)
                args = JsonConvert.DeserializeObject<string[]>(Environment.GetEnvironmentVariable("ARGJSON"));

            if (args.Length < 2)
            {
                Console.Error.WriteLine("Launcher was run with insufficient arguments. Usage: launcher.exe <app directory> <start command>");
                return 1;
            }

            var containerRoot = Directory.GetCurrentDirectory();
            var workingDirectory = Path.Combine(containerRoot, args[0]);
            var executablePathAndArgs = args[1];

            if (String.IsNullOrWhiteSpace(executablePathAndArgs))
            {
                Console.Error.WriteLine("Could not determine a start command. Use the -c flag to 'cf push' to specify a custom start command.");
                return 1;
            }
            if (executablePathAndArgs[0] == '/')
            {
                executablePathAndArgs = containerRoot + executablePathAndArgs;
            }
            Console.Out.WriteLine("Running {0}", executablePathAndArgs);

            Directory.SetCurrentDirectory(workingDirectory);
            return StartProcess(executablePathAndArgs);
        }

        private static int StartProcess(string executablePathAndArgs)
        {
            STARTUPINFO startupInformation = new STARTUPINFO();
            PROCESS_INFORMATION processInformation;
            var result = CreateProcess(null, executablePathAndArgs, IntPtr.Zero, IntPtr.Zero, false, 0, IntPtr.Zero, null, ref startupInformation, out processInformation);
            if (!result)
            {
                return Marshal.GetLastWin32Error();
            }
            WaitForSingleObject(processInformation.hProcess, INFINITE);
            UInt32 exitCode = 0;
            GetExitCodeProcess(processInformation.hProcess, ref exitCode);
            return (int) exitCode;
        }

        [DllImport("kernel32.dll", SetLastError = true)]
        static extern bool CreateProcess(string lpApplicationName, string lpCommandLine, IntPtr lpProcessAttributes,
                                        IntPtr lpThreadAttributes, bool bInheritHandles, uint dwCreationFlags,
                                        IntPtr lpEnvironment, string lpCurrentDirectory, ref STARTUPINFO lpStartupInfo,
                                        out PROCESS_INFORMATION lpProcessInformation);

        [DllImport("kernel32.dll", SetLastError = true)]
        static extern UInt32 WaitForSingleObject(IntPtr hHandle, UInt32 dwMilliseconds);

        [DllImport("kernel32.dll", SetLastError = true)]
        public static extern bool GetExitCodeProcess(IntPtr process, ref UInt32 exitCode);

        const UInt32 INFINITE = 0xFFFFFFFF;
    }

    public struct PROCESS_INFORMATION
    {
        public IntPtr hProcess;
        public IntPtr hThread;
        public uint dwProcessId;
        public uint dwThreadId;
    }

    public struct STARTUPINFO
    {
        public uint cb;
        public string lpReserved;
        public string lpDesktop;
        public string lpTitle;
        public uint dwX;
        public uint dwY;
        public uint dwXSize;
        public uint dwYSize;
        public uint dwXCountChars;
        public uint dwYCountChars;
        public uint dwFillAttribute;
        public uint dwFlags;
        public short wShowWindow;
        public short cbReserved2;
        public IntPtr lpReserved2;
        public IntPtr hStdInput;
        public IntPtr hStdOutput;
        public IntPtr hStdError;
    }

    public struct SECURITY_ATTRIBUTES
    {
        public int length;
        public IntPtr lpSecurityDescriptor;
        public bool bInheritHandle;
    }
}
