using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Builder
{
    public class TempDirectory : IDisposable
    {
        string _tmpPath;
        public TempDirectory()
        {
            _tmpPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(_tmpPath);
        }

        public string PathString()
        {
            return _tmpPath;
        }

        public string Combine(string path)
        {
            return Path.Combine(_tmpPath, path);
        }

        #region IDisposable Members
        public void Dispose()
        {
            Directory.Delete(_tmpPath, true);
        }
        #endregion
    }
}
