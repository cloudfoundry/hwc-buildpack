using CommandLine;
using CommandLine.Text;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tailor
{
    public class Options
    {
        [Option('a', "appDir", DefaultValue = "app", HelpText = "Directory to place app in")]
        public string AppDir { get; set; }

        [Option('d', "outputDroplet", DefaultValue = "app", HelpText = "Directory to the output droplet")]
        public string OutputDroplet { get; set; }

        [Option('m', "outputMetadata", DefaultValue = "app", HelpText = "Directory to the output metadata json file")]
        public string OutputMetadata { get; set; }

        [HelpOption]
        public string GetUsage()
        {
            return HelpText.AutoBuild(this,
              (HelpText current) => HelpText.DefaultParsingErrorsHandler(this, current));
        }
    }
}
