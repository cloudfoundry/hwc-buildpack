using SharpCompress.Common;
using SharpCompress.Writer;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Builder
{
    public class TarGZFile
    {
        public static void CreateFromDirectory(string sourceDirectoryName, string destinationArchiveFileName)
        {
            var tarPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tar");
            try
            {
                using (Stream stream = File.OpenWrite(tarPath))
                {
                    using (var writer = WriterFactory.Open(stream, ArchiveType.Tar, new CompressionInfo { Type = CompressionType.None }))
                    {
                        writer.WriteAll(sourceDirectoryName, "*", SearchOption.AllDirectories);
                    }
                }
                using (Stream stream = File.OpenWrite(destinationArchiveFileName))
                {
                    using (var writer = WriterFactory.Open(stream, ArchiveType.GZip, new CompressionInfo { Type = CompressionType.GZip }))
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
