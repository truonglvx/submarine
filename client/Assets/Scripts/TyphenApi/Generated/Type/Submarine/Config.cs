// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class Config : TyphenApi.TypeBase<Config>
    {
        protected static readonly SerializationInfo<Config, string> version = new SerializationInfo<Config, string>("version", false, (x) => x.Version, (x, v) => x.Version = v);
        public string Version { get; set; }
        protected static readonly SerializationInfo<Config, string> apiServerBaseUri = new SerializationInfo<Config, string>("api_server_base_uri", false, (x) => x.ApiServerBaseUri, (x, v) => x.ApiServerBaseUri = v);
        public string ApiServerBaseUri { get; set; }
    }
}
