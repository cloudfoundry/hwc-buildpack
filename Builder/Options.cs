using CommandLine;
using CommandLine.Text;

namespace Builder
{
    public class Options
    {
        [Option('a', "buildDir", Required = true, HelpText = "")]
        public string BuildDir { get; set; }

        [Option('b', "buildArtifactsCacheDir", Required = false, HelpText = "")]
        public string BuildArtifactsCacheDir { get; set; }

        [Option('o', "buildpackOrder", Required = false, HelpText = "")]
        public string BuildpackOrder { get; set; }

        [Option('e', "buildpacksDir", Required = false, HelpText = "")]
        public string BuildpacksDir { get; set; }

        [Option('c', "outputBuildArtifactsCache", Required = false, HelpText = "")]
        public string OutputBuildArtifactsCache { get; set; }

        [Option('d', "outputDroplet", Required = true, HelpText = "")]
        public string OutputDroplet { get; set; }

        [Option('m', "outputMetadata", Required = true, HelpText = "")]
        public string OutputMetadata { get; set; }

        [Option('s', "skipCertVerify", Required = false, HelpText = "")]
        public string SkipCertVerify { get; set; }

        [Option('k', "skipDetect", Required = false, HelpText = "")]
        public string skipDetect { get; set; }

        [HelpOption]
        public string GetUsage()
        {
            return HelpText.AutoBuild(this,
              (HelpText current) => HelpText.DefaultParsingErrorsHandler(this, current));
        }
    }
}
