using System;
using Newtonsoft.Json;

namespace Builder
{
    public class OutputMetadata
    {
        [JsonProperty("lifecycle_type")]
        public string LifecycleType
        {
            get { return "buildpack"; }
        }
        [JsonProperty("lifecycle_metadata")]
        public LifecycleMetadata LifecycleMetadata
        {
            get { return new LifecycleMetadata(); }
        }
        [JsonProperty("process_types")]
        public ProcessTypes ProcessTypes {
            get
            {
                return new ProcessTypes()
                {
                    Web = (ExecutionMetadata.StartCommand + " " + String.Join(" ", ExecutionMetadata.StartCommandArgs)).Trim(),
                };
            }
        }

        [JsonProperty("execution_metadata")]
        public string ExecutionMetadataJson
        {
            get { return JsonConvert.SerializeObject(ExecutionMetadata); }
        }

        [JsonIgnore]
        public ExecutionMetadata ExecutionMetadata;
    }

    public class ExecutionMetadata
    {
        [JsonProperty("start_command")]
        public string StartCommand { get; set; }
        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs { get; set; }
    }

    public class LifecycleMetadata
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
    }

    public class ProcessTypes
    {
        [JsonProperty("web")]
        public string Web { get; set; }
    }
}