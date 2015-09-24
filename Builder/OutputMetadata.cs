using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Newtonsoft.Json;

namespace Builder
{
    public class OutputMetadata
    {
        [JsonProperty("buildpack_key")]
        public string BuildpackKey
        {
            get { return ""; }
        }

        [JsonProperty("detected_buildpack")]
        public string DetectedBuildpack
        {
            get { return "windows"; }
        }

        [JsonProperty("execution_metadata")]
        public ExecutionMetadata ExecutionMetadata { get; set; }
    }

    public class ExecutionMetadata
    {
        [JsonProperty("process_types")]
        public ProcessTypes ProcessTypes { get; set; }
    }

    public class ProcessTypes
    {
        public string StartCommand { get; set; }
        public string[] StartCommandArgs { get; set; }
        [JsonProperty("web")]
        public string Web
        {
            get { return (StartCommand + " " + String.Join(" ", StartCommandArgs)).Trim(); }
        }
    }
}