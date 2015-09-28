using SharpCompress.Common;
using SharpCompress.Writer;
using System.IO;

namespace Builder
{
    public class TarGZFile
    {
        public static void CreateFromDirectory(string fullSourcePath, string destinationArchiveFileName)
        {
            var tarPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tar");
            var parentPath = Directory.GetParent(fullSourcePath).FullName;
            try
            {
                using (Stream stream = File.OpenWrite(tarPath))
                {
                    using (var writer = WriterFactory.Open(stream, ArchiveType.Tar, new CompressionInfo {Type = CompressionType.None}))
                    {
                        foreach (string source in Directory.EnumerateFiles(fullSourcePath, "*", SearchOption.AllDirectories))
                            writer.Write(source.Substring(parentPath.Length), source);
                    }
                }
                using (Stream stream = File.OpenWrite(destinationArchiveFileName))
                {
                    using (var writer = WriterFactory.Open(stream, ArchiveType.GZip, new CompressionInfo {Type = CompressionType.GZip}))
                    {
                        writer.Write("Tar.tar", tarPath);
                    }
                }
            }
            finally
            {
                if (File.Exists(tarPath)) File.Delete(tarPath);
            }
        }
    }
}
