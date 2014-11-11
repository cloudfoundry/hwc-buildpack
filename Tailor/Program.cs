using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using SharpCompress.Common;
using SharpCompress.Writer;
using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tailor
{
    public class Program
    {
        public static void Run(Options options)
        {
            // Values are available here
            Console.WriteLine("OutputMetadata: {0}", options.OutputMetadata);

            // Create temp path
            var tmpPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(tmpPath);

            // Create app.zip
            ZipFile.CreateFromDirectory(options.AppDir, Path.Combine(tmpPath, "app.zip"));

            // Create droplet (tgz)
            var tarPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tar");
            using (Stream stream = File.OpenWrite(tarPath))
            {
                using (var writer = WriterFactory.Open(stream, ArchiveType.Tar, new CompressionInfo { Type = CompressionType.None }))
                {
                    writer.WriteAll(tmpPath, "*", SearchOption.AllDirectories);
                }
            }
            using (Stream stream = File.OpenWrite(options.OutputDroplet))
            {
                using (var writer = WriterFactory.Open(stream, ArchiveType.GZip, new CompressionInfo { Type = CompressionType.GZip }))
                {
                    writer.Write("Tar.tar", tarPath);
                }
            }

            // Delete tmp path
            File.Delete(tarPath);
            Directory.Delete(tmpPath, true);

            // Result.JSON
            JObject execution_metadata = new JObject();
            execution_metadata["start_command"] = "the start command";
            JObject detected_start_command = new JObject();
            detected_start_command["web"] = "the start command";

            JObject obj = new JObject();
            obj["execution_metadata"] = execution_metadata.ToString(Formatting.None);
            obj["detected_start_command"] = detected_start_command;
            System.IO.File.WriteAllText(options.OutputMetadata, obj.ToString());
        }

        static void Main(string[] args)
        {
            var options = new Options();
            if (!CommandLine.Parser.Default.ParseArguments(args, options))
            {
                Environment.Exit(1);
            }

            Run(options);
        }
    }
}
