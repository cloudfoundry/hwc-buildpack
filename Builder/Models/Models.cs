using Newtonsoft.Json;
using System.Collections.Generic;

namespace Builder.Models
{
    public class Service
    {
        [JsonProperty("name")]
        public string Name { get; set; }
        [JsonProperty("label")]
        public string Label { get; set; }
        [JsonProperty("tags")]
        public List<string> Tags { get; set; }
        [JsonProperty("credentials")]
        public IDictionary<string, string> Credentials { get; set; }
    }


    public class Services
    {
        public Services()
        {
            UserProvided = new List<Service>();
        }

        [JsonProperty("user-provided")]
        public List<Service> UserProvided { get; set; }
    }
}
