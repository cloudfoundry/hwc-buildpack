using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Newtonsoft.Json;

namespace Builder
{
    public class OutputMetadata
    {
        public ExecutionMetadata ExecutionMetadata { get; set; }

        [JsonProperty("execution_metadata")]
        public string execution_metadata
        {
            get { return JsonConvert.SerializeObject(ExecutionMetadata); }
        }

        [JsonProperty("detected_start_command")]
        public DetectedStartCommand DetectedStartCommand {
            get
            {
                return new DetectedStartCommand()
                {
                    Web =
                        (ExecutionMetadata.StartCommand + " " + String.Join(" ", ExecutionMetadata.StartCommandArgs))
                            .Trim(),
                };
            }
        }
    }

    public class DetectedStartCommand
    {
        [JsonProperty("web")]
        public string Web { get; set; }
    }

    public class ExecutionMetadata
    {
        public ExecutionMetadata()
        {
            StartCommand = "";
            StartCommandArgs = new string[] { };
        }

        [JsonProperty("start_command")]
        public string StartCommand
        {
            get;
            set;
        }

        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs
        {
            get;
            set;
        }
    }
}